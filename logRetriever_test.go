package convoso

import (
	"testing"
	"time"
)

func TestLogRetrieverHasCorrectDateInputs(t *testing.T) {
	start := time.Now().AddDate(0, 0, 1) // one day from now
	end := time.Now()
	_, err := GetLogs(start, end)

	if err == nil {
		t.Errorf("Didn't catch the error of invalid dates")
	}
}
