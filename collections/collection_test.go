package collections

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_Simple_Collection(t *testing.T) {
	collection := newHashMapCollection("Test")

	m, _ := collection.CreateHashMap("HM1")
	g, _ := collection.CreateHashMap("HM2")

	ADDED_ELEMS_COUNT := 30_000

	key := "Key"
	value := "Value"

	for i := 0; i < ADDED_ELEMS_COUNT; i++ {

		key := key + strconv.Itoa(i)
		value := value + strconv.Itoa(i)
		m.Push(key, value)
		expectedValue, _ := m.Get(key)

		if value != expectedValue {
			t.Errorf(fmt.Sprintf("Expected %s be equal %s", value, expectedValue))
		}
	}

	t.Cleanup(func() {
		if g.count == uint32(ADDED_ELEMS_COUNT) {
			t.Errorf(fmt.Sprintf("Expected hashMap.count be equal %d", 0))
		}

		if m.count != uint32(ADDED_ELEMS_COUNT) {
			t.Errorf(fmt.Sprintf("Expected hashMap.count be equal %d", ADDED_ELEMS_COUNT))
		}
	})
}
