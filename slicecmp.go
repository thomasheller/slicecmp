package slicecmp

import (
	"bytes"
)

// PrettyPrintMulti pretty-prints slices side by side
func PrettyPrintMulti(separator rune, spacing int, headings []string, slices ...[]string) string {
	var result bytes.Buffer

	if len(headings) != len(slices) {
		panic("Number of headings and slices must be identical.")
	}

	// calculate widths:

	widths := make([]int, len(slices))

	for colIdx, slice := range slices {
		heading := headings[colIdx]

		width := maxStringLen(slice)

		if width < len(heading) {
			width = len(heading)
		}

		// width += spacing

		widths[colIdx] = width
	}

	// calculate total width:

	totalWidth := sum(widths) + spacing*(len(headings)-1)

	// find the longest slices:

	maxRowIndex := maxSliceLen(slices)

	// write headings and their padding:

	for colIdx, heading := range headings {
		result.WriteString(heading)

		width := widths[colIdx]

		if colIdx < len(headings)-1 { // skip last spacing
			width += spacing
		}

		for i := width - len(heading); i > 0; i-- {
			result.WriteString(" ")
		}
	}

	// print separator line:

	result.WriteString("\n")

	for i := 0; i < totalWidth; i++ {
		result.WriteString(string(separator))
	}

	// print slice values and their padding:

	for rowIdx := 0; rowIdx < maxRowIndex; rowIdx++ {
		result.WriteString("\n")

		for colIdx, slice := range slices {

			column := ""
			if len(slice) > rowIdx {
				column = slice[rowIdx]
			}

			result.WriteString(column)

			width := widths[colIdx]

			if colIdx < len(slices)-1 { // skip last padding
				width += spacing
			}

			for i := len(column); i < width; i++ {
				result.WriteString(" ")
			}
		}
	}

	return result.String()
}

// PrettyPrint pretty-prints two slices side by side
func PrettyPrint(headingA string, sliceA []string, headingB string, sliceB []string, separator rune, spacing int) string {
	var result bytes.Buffer

	widthA := maxStringLen(sliceA)

	if widthA < len(headingA) {
		widthA = len(headingA)
	}

	widthA += spacing

	result.WriteString(headingA)

	for i := widthA - len(headingA); i > 0; i-- {
		result.WriteString(" ")
	}

	result.WriteString(headingB)

	widthB := maxStringLen(sliceB)

	for i := widthB - len(headingB); i > 0; i-- {
		result.WriteString(" ")
	}

	result.WriteString("\n")

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

		for i := len(column); i < widthA; i++ {
			result.WriteString(" ")
		}

		if len(sliceB) > idx {
			column = sliceB[idx]

			result.WriteString(column)
		} else {
			column = ""
		}

		for i := len(column); i < widthB; i++ {
			result.WriteString(" ")
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

// maxSliceLen returns the longest slices length in a slice of slices
func maxSliceLen(slices [][]string) int {
	var result int

	for _, s := range slices {
		if len(s) > result {
			result = len(s)
		}
	}

	return result
}

// maxStringLen returns the longest string length in a slice
func maxStringLen(s []string) int {
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

// sum calculates the sum of integers
func sum(ints []int) int {
	var result int
	for _, i := range ints {
		result += i
	}
	return result
}
