package slicecmp

import (
	"testing"
)

func TestCompareBothNil(t *testing.T) {
	testCompare(t, nil, nil, true)
}

func TestCompareLeftNil(t *testing.T) {
	testCompare(t, nil, []string{"foo"}, false)
}

func TestCompareRightNil(t *testing.T) {
	testCompare(t, []string{"foo"}, nil, false)
}

func TestCompareIdentical(t *testing.T) {
	testCompare(t, []string{"foo", "bar"}, []string{"foo", "bar"}, true)
}

func TestCompareSameLength(t *testing.T) {
	testCompare(t, []string{"foo", "foo"}, []string{"bar", "bar"}, false)
}

func TestCompareSameLeft(t *testing.T) {
	testCompare(t, []string{"foo", "foo"}, []string{"foo"}, false)
}

func TestCompareSameRight(t *testing.T) {
	testCompare(t, []string{"foo"}, []string{"foo", "foo"}, false)
}

func TestCompareDifferentLeft(t *testing.T) {
	testCompare(t, []string{"foo", "foo", "foo"}, []string{"bar", "bar"}, false)
}

func TestCompareDifferentRight(t *testing.T) {
	testCompare(t, []string{"foo", "foo"}, []string{"bar", "bar", "bar"}, false)
}

func testCompare(t *testing.T, a, b []string, expected bool) {
	actual := Equal(a, b)

	if actual != expected {
		t.Errorf("Test failed.\nInput slice a: %#v\nInput slice b: %#v\nExpected: %t\nActual: %t\n", a, b, expected, actual)
	}
}

func TestPrettyPrintIdentical(t *testing.T) {
	sl := []string{"foo", "foo", "foo"}
	testPrettyPrint(t, "foo", sl, "foo", sl, '-', 1, `foo foo
-------
foo foo
foo foo
foo foo`)
}

func TestPrettyPrintLeft(t *testing.T) {
	a := []string{"foo", "foo", "foo"}
	b := []string{"bar", "bar"}
	testPrettyPrint(t, "foo", a, "bar", b, '-', 1, `foo bar
-------
foo bar
foo bar
foo`)
}

func TestPrettyPrintRight(t *testing.T) {
	a := []string{"foo", "foo"}
	b := []string{"bar", "bar", "bar"}
	testPrettyPrint(t, "foo", a, "bar", b, '-', 1, `foo bar
-------
foo bar
foo bar
    bar`)
}

func TestPrettyPrintWider(t *testing.T) {
	sl := []string{"foo", "foo", "foo"}
	testPrettyPrint(t, "foo", sl, "foo", sl, '-', 5, `foo     foo
-----------
foo     foo
foo     foo
foo     foo`)
}

func TestPrettyPrintOtherSeparator(t *testing.T) {
	sl := []string{"foo", "foo", "foo"}
	testPrettyPrint(t, "foo", sl, "foo", sl, '=', 1, `foo foo
=======
foo foo
foo foo
foo foo`)
}

func TestPrettyPrintLeftHeading(t *testing.T) {
	a := []string{"foo", "foo", "foo"}
	b := []string{"bar", "bar", "bar"}
	testPrettyPrint(t, "alpha", a, "bar", b, '-', 1, `alpha bar
---------
foo   bar
foo   bar
foo   bar`)
}

func TestPrettyPrintRightHeading(t *testing.T) {
	a := []string{"foo", "foo", "foo"}
	b := []string{"bar", "bar", "bar"}
	testPrettyPrint(t, "foo", a, "beta", b, '-', 1, `foo beta
--------
foo bar
foo bar
foo bar`)
}

func testPrettyPrint(t *testing.T, headingA string, sliceA []string, headingB string, sliceB []string, separator rune, spacing int, expected string) {
	actual := PrettyPrint(headingA, sliceA, headingB, sliceB, separator, spacing)
	if actual != expected {
		t.Errorf("Test failed.\nInput slice a: %#v\nInput slice b: %#v\nExpected:\n\n%s\n\nActual:\n\n%s\n\n", sliceA, sliceB, expected, actual)
	}
}
