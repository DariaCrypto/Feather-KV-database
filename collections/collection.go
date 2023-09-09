package collections

import (
	"errors"
)

type HashMapCollection struct {
	hashMap map[string]*HashMap
}

type SortedSetCollection struct {
	sortedSet map[string]*SortedSet
}

func NewHashMapCollection() *HashMapCollection {
	hm := make(map[string]*HashMap)
	return &HashMapCollection{hashMap:hm}
}

func NewSortedSetCollection() *SortedSetCollection {
	ss := make(map[string]*SortedSet)
	return &SortedSetCollection{sortedSet:ss}
}

func (hm *HashMapCollection) HGet(nameHM, key string) (string, error) {
	hashmap, ok := hm.hashMap[nameHM]
	if !ok {
		return "", errors.New("collections: hashmap not found")
	}

	result, err := hashmap.Get(key)
	if err != nil {
		return "", errors.New("collections: can't get record from a hashmap")
	}
	return result, nil
}

func (hm *HashMapCollection) HSet(nameHM, key, value string) error {
	hashmap, ok := hm.hashMap[nameHM]
	if !ok {
		hm.hashMap[nameHM] = NewHashMap(nameHM)
	}

	err := hashmap.Push(key, value)
	if err != nil {
		return errors.New("collections: can't add record to a hashmap")
	}
	return nil
}

func (hm *HashMapCollection) HDelete(nameHM, key, value string) error {
	hashmap, ok := hm.hashMap[nameHM]
	if !ok {
		return errors.New("collections: hashmap not found")
	}

	err := hashmap.Pop(key)
	if err != nil {
		return errors.New("collections: can't delete record from a hashmap")
	}
	return nil
}

func (ss *SortedSetCollection) ZSet(name, value string, score uint32) error {
	sortedset, ok := ss.sortedSet[name]
	if !ok {
		ss.sortedSet[name] = NewSortedSet(name)
	}

	sortedset.addNode(value, score)
	return nil
}

func (ss *SortedSetCollection) ZDelete(name, value string, score uint32) error {
	sortedset, ok := ss.sortedSet[name]
	if !ok {
		return errors.New("collections: storage not found")
	}
	sortedset.deleteNode(value, score)
	return nil
}

func (ss *SortedSetCollection) ZGet(name, value string) (uint32, string, error) {
	sortedset, ok := ss.sortedSet[name]
	if !ok {
		return 0, "", errors.New("collections: storage not found")
	}

	score, value := sortedset.getNode(value)
	return score, value, nil
}
