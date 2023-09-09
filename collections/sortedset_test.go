package collections

import (
	"testing"
)

func Test_Simple_SortedSet_Test(t *testing.T) {
	sortedSet := NewSortedSet("Sorted Set")

	value := "Emily"
	var score uint32 = 33

	sortedSet.addNode(value, score)
	expectedScore, expectedValue := sortedSet.getNode(value)
	if expectedScore != score || expectedValue != value {
		t.Errorf("Expected value be equal %s and score equal %d", value, score)
	}

	value = "Ben"
	score = 13
	sortedSet.addNode(value, score)
	expectedScore, expectedValue = sortedSet.getNode(value)
	if expectedScore != score || expectedValue != value {
		t.Errorf("Expected value be equal %s and score equal %d", value, score)
	}

	value = "Sten"
	score = 53
	sortedSet.addNode(value, score)
	expectedScore, expectedValue = sortedSet.getNode(value)
	if expectedScore != score || expectedValue != value {
		t.Errorf("Expected value be equal %s and score equal %d", value, score)
	}

	value = "Martin"
	score = 9
	sortedSet.addNode(value, score)
	expectedScore, expectedValue = sortedSet.getNode(value)
	if expectedScore != score || expectedValue != value {
		t.Errorf("Expected value be equal %s and score equal %d", value, score)
	}

	value = "Li"
	score = 21
	sortedSet.addNode(value, score)
	expectedScore, expectedValue = sortedSet.getNode(value)
	if expectedScore != score || expectedValue != value {
		t.Errorf("Expected value be equal %s and score equal %d", value, score)
	}

	value = "Zack"
	score = 61
	sortedSet.addNode(value, score)
	expectedScore, expectedValue = sortedSet.getNode(value)
	if expectedScore != score || expectedValue != value {
		t.Errorf("Expected value be equal %s and score equal %d", value, score)
	}

	value = "River"
	score = 8
	sortedSet.addNode(value, score)
	expectedScore, expectedValue = sortedSet.getNode(value)
	if expectedScore != score || expectedValue != value {
		t.Errorf("Expected value be equal %s and score equal %d", value, score)
	}

	value = "Kate"
	score = 11
	sortedSet.addNode(value, score)
	expectedScore, expectedValue = sortedSet.getNode(value)
	if expectedScore != score || expectedValue != value {
		t.Errorf("Expected value be equal %s and score equal %d", value, score)
	}

	maxValue := "Zack"
	var maxScore uint32 = 61
	expectedMaxValue, expectedMaxScore := sortedSet.getMaxScoreNode()
	if expectedMaxScore != maxScore || expectedMaxValue != maxValue {
		t.Errorf("Expected value be equal %s and score equal %d", value, score)
	}

	minValue := "River"
	var minScore uint32 = 8
	sortedSet.deleteNode(maxValue, maxScore)
	expectedSecondMaxValue, expectedSecondMaxScore := sortedSet.getMinScoreNode()
	if expectedSecondMaxScore != minScore || expectedSecondMaxValue != minValue {
		t.Errorf("Expected value be equal %s and score equal %d", value, score)
	}

	deletedValue := "Zack"
	var deletedScore uint32 = 61
	sortedSet.deleteNode(deletedValue, deletedScore)
	expectedScore, expectedValue = sortedSet.getNode(deletedValue)
	if expectedValue == deletedValue || expectedScore == deletedScore {
		t.Errorf("Expected value not be equal %s and score equal %d", value, score)
	}

}

func Test_Parallel_RWD_SortedSet(t *testing.T) {
	sortedSet := NewSortedSet("Sorted Set")
	var value string
	var score uint32

	t.Run("Worker 1: Parallel addition of elements to the sortedset. ", func(t *testing.T) {
		t.Parallel()
		value = "Emily"
		score = 33

		sortedSet.addNode(value, score)
		expectedScore, expectedValue := sortedSet.getNode(value)
		if expectedScore != score || expectedValue != value {
			t.Errorf("Expected value be equal %s and score equal %d", value, score)
		}

		value = "Kate"
		score = 90
		sortedSet.addNode(value, score)
		expectedScore, expectedValue = sortedSet.getNode(value)
		if expectedScore != score || expectedValue != value {
			t.Errorf("Expected value be equal %s and score equal %d", value, score)
		}

	})

	t.Run("Worker 2: Parallel addition of elements to the sortedset. ", func(t *testing.T) {
		t.Parallel()

		value = "Ben"
		score = 13
		sortedSet.addNode(value, score)
		expectedScore, expectedValue := sortedSet.getNode(value)
		if expectedScore != score || expectedValue != value {
			t.Errorf("Expected value be equal %s and score equal %d", value, score)
		}

		value = "Emily"
		score = 33
		sortedSet.addNode(value, score)
		expectedScore, expectedValue = sortedSet.getNode(value)
		if expectedScore != score || expectedValue != value {
			t.Errorf("Expected value be equal %s and score equal %d", value, score)
		}
	})

	t.Run("Worker 3: Parallel addition of elements to the sortedset. ", func(t *testing.T) {
		t.Parallel()

		value = "Sten"
		score = 53
		sortedSet.addNode(value, score)
		expectedScore, expectedValue := sortedSet.getNode(value)
		if expectedScore != score || expectedValue != value {
			t.Errorf("Expected value be equal %s and score equal %d", value, score)
		}
	})
	t.Run("Worker 4: Parallel addition of elements to the sortedset. ", func(t *testing.T) {
		t.Parallel()

		value = "Martin"
		score = 9
		sortedSet.addNode(value, score)
		expectedScore, expectedValue := sortedSet.getNode(value)
		if expectedScore != score || expectedValue != value {
			t.Errorf("Expected value be equal %s and score equal %d", value, score)
		}

		value = "River"
		score = 12
		sortedSet.addNode(value, score)
		expectedScore, expectedValue = sortedSet.getNode(value)
		if expectedScore != score || expectedValue != value {
			t.Errorf("Expected value be equal %s and score equal %d", value, score)
		}

	})
	t.Run("Worker 5: Parallel addition of elements to the sortedset. ", func(t *testing.T) {
		t.Parallel()

		value = "Li"
		score = 21
		sortedSet.addNode(value, score)
		expectedScore, expectedValue := sortedSet.getNode(value)
		if expectedScore != score || expectedValue != value {
			t.Errorf("Expected value be equal %s and score equal %d", value, score)
		}

		value = "Sten"
		score = 23
		sortedSet.addNode(value, score)
		expectedScore, expectedValue = sortedSet.getNode(value)
		if expectedScore != score || expectedValue != value {
			t.Errorf("Expected value be equal %s and score equal %d", value, score)
		}
	})
	t.Run("Worker 6: Parallel addition of elements to the sortedset. ", func(t *testing.T) {
		t.Parallel()

		value = "Zack"
		score = 61
		sortedSet.addNode(value, score)
		expectedScore, expectedValue := sortedSet.getNode(value)
		if expectedScore != score || expectedValue != value {
			t.Errorf("Expected value be equal %s and score equal %d", value, score)
		}
	})
	t.Run("Worker 7: Parallel addition of elements to the sortedset. ", func(t *testing.T) {
		t.Parallel()

		value = "River"
		score = 8
		sortedSet.addNode(value, score)
		expectedScore, expectedValue := sortedSet.getNode(value)
		if expectedScore != score || expectedValue != value {
			t.Errorf("Expected value be equal %s and score equal %d", value, score)
		}
	})
	t.Run("Worker 8: Parallel addition of elements to the sortedset. ", func(t *testing.T) {
		t.Parallel()

		value = "Kate"
		score = 11
		sortedSet.addNode(value, score)
		expectedScore, expectedValue := sortedSet.getNode(value)
		if expectedScore != score || expectedValue != value {
			t.Errorf("Expected value be equal %s and score equal %d", value, score)
		}
	})

}
