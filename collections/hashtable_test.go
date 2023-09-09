package collections

import (
	"strconv"
	"testing"
	"fmt"
)

func Test_Simple_Hashtable(t *testing.T) {
	hashMap := NewHashMap("Test")
	ADDED_ELEMS_COUNT := 30_000

	key := "Key"
	value := "Value"

	for i := 0; i < ADDED_ELEMS_COUNT; i++ {

		key := key + strconv.Itoa(i)
		value := value + strconv.Itoa(i)
		hashMap.Push(key, value)
		expectedValue, _ := hashMap.Get(key)
		
		if value != expectedValue {
			t.Errorf(fmt.Sprintf("Expected %s be equal %s", value, expectedValue))
		}
	}

	t.Cleanup(func(){
		if hashMap.count != uint32(ADDED_ELEMS_COUNT) {
			t.Errorf(fmt.Sprintf("Expected hashMap.count be equal %d", ADDED_ELEMS_COUNT))
		}
	})
}

func Test_Parallel_RWD_Hashtable(t *testing.T) {
	hashMap := NewHashMap("Test")
	ADDED_ELEMS_COUNT := 30_000
	DELETED_ELEMS_COUNT := 5_000

	//Worker 1: Add elements to hashmap
	t.Run("Worker 1: Parallel addition of elements to the map. ", func(t *testing.T) {
		t.Parallel()
		key := "1_Key"
		value := "1_Value"
		for i := 0; i < ADDED_ELEMS_COUNT; i++ {

			key := key + strconv.Itoa(i)
			value := value + strconv.Itoa(i)
			hashMap.Push(key, value)
			expectedValue, _ := hashMap.Get(key)
			
			if value != expectedValue {
				t.Errorf(fmt.Sprintf("Expected %s be equal %s", value, expectedValue))
			}
		}

	})

	//Worker 2: Add elements to hashmap
	t.Run("Worker 2: Parallel addition of elements to the map. ", func(t *testing.T) {
		t.Parallel()

		key := "2_Key"
		value := "2_Value"

		for i := 0; i < ADDED_ELEMS_COUNT; i++ {
			key := key + strconv.Itoa(i)
			value := value + strconv.Itoa(i)
			hashMap.Push(key, value)
			expectedValue, _ := hashMap.Get(key)
			
			if value != expectedValue {
				t.Errorf(fmt.Sprintf("Expected %s be equal %s", value, expectedValue))
			}
		}

	})

	//Delete elements from hashmap
	t.Cleanup(func(){
		key := "1_Key"
		value := "1_Value"
		for i := 0; i < DELETED_ELEMS_COUNT; i++ {
			key := key + strconv.Itoa(i)
			value := value + strconv.Itoa(i)
			err := hashMap.Pop(key)
			if err != nil{
				t.Error("Expected elements to be added:", err)
			}
			expectedValue, _ := hashMap.Get(key)

			if value == expectedValue {
				t.Errorf(fmt.Sprintf("Expected %s be equal %s", value, expectedValue))
			}
		}

		exptectedCountElems := uint32((ADDED_ELEMS_COUNT*2) - DELETED_ELEMS_COUNT)
		if hashMap.count != exptectedCountElems {
			t.Errorf(fmt.Sprintf("Expected hashMap.count be equal %d", exptectedCountElems))
		}
	})
}

func Benchmark_HashTable(b *testing.B) {
	hashMap := NewHashMap("Test")

	key := "Key"
	value := "Value"

	for i := 0; i < b.N; i++ {
		hashMap.Push(key+strconv.Itoa(i), value+strconv.Itoa(i))
	}
}
