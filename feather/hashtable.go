package inmemory

import (
	"errors"
	"sync"
	"sync/atomic"
)

const (
	BUCKET_SIZE         = 1024
	INIT_BACKET_SIZE    = 2
	BUCKET_LOAD_FACTORY = 0.86
)

type HashMap struct {
	name        string
	mu          sync.Mutex
	count       uint32 // total number of elements in buckets
	size        uint32 // size of buckets
	buckets     []Bucket
	isRehashing atomic.Bool
}

type Bucket struct {
	Locker sync.Mutex
	Data   *Node
}

type Node struct {
	key   string
	value string
	next  *Node
}

func NewHashMap(name string) *HashMap {
	buckets := make([]Bucket, BUCKET_SIZE)
	return &HashMap{name: name, size: BUCKET_SIZE, buckets: buckets}
}

func (hashMap *HashMap) getLoadFactor() float32 {
	return float32(hashMap.count) / float32(hashMap.size)
}

func (hashMap *HashMap) getIndexBucket(key string, size uint32) uint32 {
	return Hash([]byte(key)) & (size - 1)
}

// If our backets are big, hashmaps need to resize
func (hashMap *HashMap) checkAndRehash() {
	hashMap.mu.Lock()
	if hashMap.getLoadFactor() >= BUCKET_LOAD_FACTORY {
		hashMap.rehash()
	}
	hashMap.mu.Unlock()
}

// Insert to HashMap
func (hashMap *HashMap) Push(key, value string) error {
	defer hashMap.checkAndRehash()
	index := hashMap.getIndexBucket(key, hashMap.size)
	hashMap.buckets[index].Locker.Lock()
	defer hashMap.buckets[index].Locker.Unlock()

	if hashMap.buckets[index].Data == nil {
		hashMap.buckets[index].Data = &Node{key: key, value: value, next: nil}
		hashMap.count++
		return nil
	} else {
		startingNode := hashMap.buckets[index].Data

		if startingNode.key == key {
			startingNode.value = value
			return nil
		} else {
			for ; startingNode.next != nil; startingNode = startingNode.next {
				if startingNode.next.key == key {
					startingNode.next.value = value
					return nil
				}
			}
			startingNode.next = &Node{key: key, value: value, next: nil}
			hashMap.count++
		}
	}
	return errors.New("hashmap: item has not been added")
}

// Get from HashMap
func (hashMap *HashMap) Get(key string) (string, error) {
	index := hashMap.getIndexBucket(key, hashMap.size)
	hashMap.buckets[index].Locker.Lock()
	defer hashMap.buckets[index].Locker.Unlock()
	node := hashMap.buckets[index].Data
	for node != nil {
		if node.key == key {
			return node.value, nil
		}
		node = node.next
	}

	node = hashMap.buckets[index].Data
	for node != nil {
		if node.key == key {
			return node.value, nil
		}
		node = node.next
	}

	return "", errors.New("hashmap: there is no node with this key")
}

// Pop to HashMap
func (hashMap *HashMap) Pop(key string) error {
	index := hashMap.getIndexBucket(key, hashMap.size)

	hashMap.buckets[index].Locker.Lock()
	defer hashMap.buckets[index].Locker.Unlock()

	if hashMap.buckets[index].Data != nil {
		startingNode := hashMap.buckets[index].Data

		if startingNode.key == key {
			hashMap.buckets[index].Data = startingNode.next
			hashMap.count--
			return nil
		}

		for prev := startingNode; startingNode != nil; startingNode = startingNode.next {
			if startingNode.key == key {
				prev.next = startingNode.next
				hashMap.count--
				return nil
			}
			prev = startingNode
		}
	}
	return errors.New("hashmap: can't find an item to delete")
}

func (h *HashMap) lockAllBackets() {
	for idx, _ := range h.buckets {
		h.buckets[idx].Locker.Lock()
	}
}

func (h *HashMap) unlockAllBackets() {
	for idx, _ := range h.buckets {
		h.buckets[idx].Locker.Unlock()
	}
}

func (hashMap *HashMap) rehash() {
	if !hashMap.isRehashing.Load() {
		hashMap.isRehashing.Swap(true)
		hashMap.lockAllBackets()

		size := hashMap.size << 1
		newBuckets := make([]Bucket, size)

		for idx, _ := range hashMap.buckets {
			bucket := &hashMap.buckets[idx]
			if bucket.Data != nil {
				startingNode := bucket.Data
				for startingNode != nil {
					newIndex := hashMap.getIndexBucket(startingNode.key, size)
					if newBuckets[newIndex].Data == nil {
						newBuckets[newIndex].Data = startingNode
						startingNode = startingNode.next
						newBuckets[newIndex].Data.next = nil
					} else {
						newBacketNode := newBuckets[newIndex].Data
						for ; newBacketNode != nil; newBacketNode = newBacketNode.next {
							if newBacketNode.next == nil {
								newBacketNode.next = startingNode
								startingNode = startingNode.next
								newBacketNode.next.next = nil
								break
							}
						}
					}

				}
			}
		}
		hashMap.unlockAllBackets()
		hashMap.size = size
		hashMap.buckets = newBuckets
		hashMap.isRehashing.Swap(false)
	}

}

func Hash(data []byte) uint32 {
	var hash uint32 = 5381
	for _, b := range data {
		hash = (hash * 33) ^ uint32(b)
	}
	return hash
}
