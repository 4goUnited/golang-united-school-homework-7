package coverage

import (
	"fmt"
	"os"
	"strconv"
	"reflect"
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
		p      People
		expLen int
	}{
		"empty_slice":      {[]Person{}, 0},
		"one_people":       {[]Person{{"Ivan", "Ivanov", time.Now()}}, 1},
		"two_people":       {[]Person{{"Ivan", "Ivanov", time.Now()}, {"Sergey", "Alekseev", time.Now()}}, 2},
		"three_nil_values": {[]Person{Person{}, Person{}, Person{}}, 3}}

	for name, tcase := range tData {
		got := tcase.p.Len()
		if got != tcase.expLen {
			t.Errorf("[%s] expected: %d, got %d", name, tcase.expLen, got)
		}
	}
}

func Test_Less(t *testing.T) {
	tData := map[string]struct {
		p       People
		expLess bool
	}{
		"all_equal":         {[]Person{{"Ivan", "Ivanov", bDate}, {"Ivan", "Ivanov", bDate}}, false},
		"diff_last_name_1":  {[]Person{{"Ivan", "Ivanov", bDate}, {"Ivan", "Alekseev", bDate}}, false},
		"diff_last_name_2":  {[]Person{{"Ivan", "Alekseev", bDate}, {"Ivan", "Ivanov", bDate}}, true},
		"diff_first_name_1": {[]Person{{"Ivan", "Ivanov", bDate}, {"Alex", "Alekseev", bDate}}, false},
		"diff_first_name_2": {[]Person{{"Alex", "Ivanov", bDate}, {"Ivan", "Alekseev", bDate}}, true},
		"diff_bDate_name_1": {[]Person{{"Alex", "Ivanov", bDate}, {"Ivan", "Alekseev", bDate.AddDate(1, 0, 0)}}, false},
		"diff_bDate_name_2": {[]Person{{"Alex", "Ivanov", bDate}, {"Ivan", "Alekseev", bDate.AddDate(-1, 0, 0)}}, true}}

	for name, tcase := range tData {
		got := tcase.p.Less(0, 1)
		if got != tcase.expLess {
			t.Errorf("[%s] expected: %t, got %t", name, tcase.expLess, got)
		}
	}
}

func Test_Swap(t *testing.T) {
	tData := map[string]struct {
		p       People
		expSwap People
	}{
		"equal": {[]Person{{"Ivan", "Ivanov", bDate}, {"Alex", "Alekseev", bDate.AddDate(1, 0, 0)}},
			[]Person{{"Alex", "Alekseev", bDate.AddDate(1, 0, 0)}, {"Ivan", "Ivanov", bDate}}}}

	for name, tcase := range tData {
		got := tcase.p
		got.Swap(0, 1)
		if got[0] != tcase.expSwap[0] && got[1] != tcase.expSwap[1] {
			t.Errorf("[%s] expected: %v, got %v", name, tcase.expSwap, got)
		}
	}
}

func Test_NewMatrix(t *testing.T) {
	tData := map[string]struct {
		str       string
		expMatrix *Matrix
		expErr    error
	}{
		"empty_string": {"", nil, &strconv.NumError{Func: "Atoi", Num: "", Err: strconv.ErrSyntax}}, 
		"valid_matr_1x1" : {"1", &Matrix{1, 1, []int{1}}, nil},
		"valid_matr_1x3" : {"1 2 3", &Matrix{1, 3, []int{1, 2, 3}}, nil},
		"valid_matr_3x1" : {"1\n2\n3", &Matrix{3, 1, []int{1, 2, 3}}, nil},
		"valid_matr_2x2" : {"1 2\n3 4", &Matrix{2, 2, []int{1, 2, 3, 4}}, nil},
		"invalid_matr_1x1" : {"a", nil, &strconv.NumError{Func: "Atoi", Num: "a", Err: strconv.ErrSyntax}},
		"invalid_matr_2x2" : {"1 2\n3", nil, fmt.Errorf("Rows need to be the same length")}}

	for name, tcase := range tData {
		got, err := New(tcase.str)
		if !reflect.DeepEqual(got, tcase.expMatrix) {
			t.Errorf("[%s] expected: %v, got %v", name, tcase.expMatrix, got)
		}
		if err != nil && err.Error() != tcase.expErr.Error() {
			t.Errorf("%s:\n wrong type of error is wrapped into the returned error: got %s, want %s", name, err, tcase.expErr)
		}
	}
}
