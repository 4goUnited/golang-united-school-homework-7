package coverage

import (
	"os"
	"testing"
	"time"
)

var bDate time.Time = time.Date(2000, 1, 1, 12, 0, 0, 0, time.UTC)

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

func Test_Less(t *testing.T) {
	tData := map[string]struct {
		p People
		expLess bool
	}{
		"all_equal": {[]Person{{"Ivan", "Ivanov", bDate}, {"Ivan", "Ivanov", bDate}}, false},
		"diff_last_name_1": {[]Person{{"Ivan", "Ivanov", bDate}, {"Ivan", "Alekseev", bDate}}, false},
		"diff_last_name_2": {[]Person{{"Ivan", "Alekseev", bDate}, {"Ivan", "Ivanov", bDate}}, true},
		"diff_first_name_1": {[]Person{{"Ivan", "Ivanov", bDate}, {"Alex", "Alekseev", bDate}}, false},
		"diff_first_name_2": {[]Person{{"Alex", "Ivanov", bDate}, {"Ivan", "Alekseev", bDate}}, true},
		"diff_bDate_name_1": {[]Person{{"Alex", "Ivanov", bDate}, {"Ivan", "Alekseev", bDate.AddDate(1, 0, 0)}}, false},
		"diff_bDate_name_2": {[]Person{{"Alex", "Ivanov", bDate}, {"Ivan", "Alekseev", bDate.AddDate(-1, 0, 0)}}, true}, }

	for name, tcase := range tData {
		got := tcase.p.Less(0, 1)
		if got != tcase.expLess {
			t.Errorf("[%s] expected: %t, got %t", name, tcase.expLess, got)
		}
	}
}

