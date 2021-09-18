package luxmatricule

import "testing"

// Thanks to article https://medium.com/rungo/unit-testing-made-easy-in-go-25077669318
// for the very interesting insights

// define input-result struct type
type TestMatriculeItem struct {
	matricule string
	result    bool
	hasError  bool
}

// test IsValid function
func TestIsValid(t *testing.T) {

	matriculeItems := []TestMatriculeItem{
		// test different lengths
		{"", false, false},
		{"A", false, false},
		{"12A3", false, false},
		{"123", false, false},

		// Test with non-digit char
		{"123456789012A", false, false},
		{"1B23456789012", false, false},
		{"C12456789012A", false, false},
		{"+12456789012A", false, false},
		{"+124567890122", false, false},
		{"+1245678901223", false, false},

		// Test invalid matricules (random numbers)
		{"1234567890123", false, false},

		// Test matricules which do no make sense (e.g month > 12) ?

		// Valid matricules
		{"2015123101725", true, false},
		{"2014082601689", true, false},
		{"1974091186700", true, false},
		{"1982080285657", true, false},
	}

	for _, item := range matriculeItems {
		result, err := IsValid(item.matricule)

		if item.hasError {
			// Expected an error
			if err == nil {
				t.Errorf("IsValid() with arg '%v' : FAILED, expected an error but got value '%v'", item.matricule, result)
			} else {
				t.Logf("IsValid() with arg '%v' : PASSED, expected an error and got error '%v'", item.matricule, err)
			}
		} else {
			// Expected a value
			if result != item.result {
				t.Errorf("IsValid() with arg '%v' : FAILED, expected '%v' but got value '%v'", item.matricule, item.result, result)
			} else {
				t.Logf("IsValid() with arg '%v' : PASSED, expected '%v' and got value '%v'", item.matricule, item.result, result)
			}
		}
	}

}
