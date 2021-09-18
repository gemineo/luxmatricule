// Package luxmatricule provides fonctions to deal with luxembourguish matricules
package luxmatricule

import (
	"fmt"
	"strings"

	"github.com/osamingo/checkdigit"
)

// v0.0.1

// This function returns true is s contains only digits; false otherwise
func containsOnlyDigits(s string) bool {
	isNotDigit := func(c rune) bool { return c < '0' || c > '9' }

	return strings.IndexFunc(s, isNotDigit) == -1
}

// Compute a Lux matricule checksum
func generateChecksum(s string) (string, error) {
	// compute the Luhn value of the query
	l := checkdigit.NewLuhn()
	lcd, err := l.Generate(s)
	if err != nil {
		return "", fmt.Errorf("Failed to generate Luhn check digit")
	}

	v := checkdigit.NewVerhoeff()
	vcd, err := v.Generate(s)
	if err != nil {
		return "", fmt.Errorf("Failed to generate Verhoeff check digit")
	}

	checksum := fmt.Sprintf("%d%d", lcd, vcd)

	return checksum, nil
}

// IsValid is used to check if a Lux matricule is valid
func IsValid(matricule string) (bool, error) {
	r := true

	if len([]rune(matricule)) != 13 {
		// Matricule is of wrong size
		r = false
	} else if !containsOnlyDigits(matricule) {
		// Matricule contains non-digit character
		r = false
	} else {
		// compute the checksum of the matricule's root
		checksum, err := generateChecksum(matricule[0:11])
		if err != nil {
			return false, err
		}
		// Compare checksums (computed vs given)
		if checksum != matricule[11:] {
			r = false
		}
	}

	return r, nil
}
