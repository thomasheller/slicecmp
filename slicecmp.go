package slicecmp

import (
	"bytes"
)

// PrettyPrint pretty-prints two slices side by side
func PrettyPrint(headingA string, sliceA []string, headingB string, sliceB []string, separator rune, spacing int) string {
	var result bytes.Buffer

	widthA := maxLen(sliceA)

	if widthA < len(headingA) {
		widthA = len(headingA)
	}

	widthA += spacing

	result.WriteString(headingA)

	for i := 0; i < widthA-len(headingA); i++ {
		result.WriteString(" ")
	}

	result.WriteString(headingB)
	result.WriteString("\n")

	widthB := maxLen(sliceB)

	if widthB < len(headingB) {
		widthB = len(headingB)
	}

	for i := 0; i < widthA+widthB; i++ {
		result.WriteString(string(separator))
	}

	for idx := 0; idx < max(len(sliceA), len(sliceB)); idx++ {
		result.WriteString("\n")

		column := ""
		if len(sliceA) > idx {
			column = sliceA[idx]
		}

		result.WriteString(column)

		if len(sliceB) > idx {
			for i := len(column); i < widthA; i++ {
				result.WriteString(" ")
			}

			result.WriteString(sliceB[idx])
		}
	}

	return result.String()
}

// Equal reports if two string slices are identical
func Equal(a, b []string) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for idx := range a {
		if a[idx] != b[idx] {
			return false
		}
	}

	return true
}

// maxLen returns the longest string length in a slice
func maxLen(s []string) int {
	var result int

	for _, e := range s {
		if len(e) > result {
			result = len(e)
		}
	}

	return result
}

// max returns the larger of two integers
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
