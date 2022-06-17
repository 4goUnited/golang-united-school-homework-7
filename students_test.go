package coverage

import (
	"os"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW
func Test_Len(t *testing.T) {
	tData := map[string]struct {
		p People
		expLen int
	}{
		"empty_slice": {[]Person{}, 0},
		"one_people": {[]Person{{"Ivan", "Ivanov", time.Now()}}, 1},
		"two_people": {[]Person{{"Ivan", "Ivanov", time.Now()}, {"Sergey", "Alekseev", time.Now()}}, 2},
		"three_nil_values": {[]Person{Person{}, Person{}, Person{}}, 3}, }
	for name, tcase := range tData {
		got := tcase.p.Len()
		if got != tcase.expLen {
			t.Errorf("[%s] expected: %d, got %d", name, tcase.expLen, got)
		}
	}
}
