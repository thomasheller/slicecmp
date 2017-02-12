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
foo    `)
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
foo bar `)
}

func TestPrettyPrintLeftWider(t *testing.T) {
	a := []string{"foo", "fooooo", "foo"}
	b := []string{"foo", "foo", "foo"}
	testPrettyPrint(t, "foo", a, "foo", b, '-', 1, `foo    foo
----------
foo    foo
fooooo foo
foo    foo`)
}

func TestPrettyPrintRightWider(t *testing.T) {
	a := []string{"foo", "foo", "foo"}
	b := []string{"foo", "fooooo", "foo"}
	testPrettyPrint(t, "foo", a, "foo", b, '-', 1, `foo foo   
----------
foo foo   
foo fooooo
foo foo   `)
}

func testPrettyPrint(t *testing.T, headingA string, sliceA []string, headingB string, sliceB []string, separator rune, spacing int, expected string) {
	actual := PrettyPrint(headingA, sliceA, headingB, sliceB, separator, spacing)
	if actual != expected {
		t.Errorf("Test failed.\nInput slice a: %#v\nInput slice b: %#v\nExpected:\n\n%s\n\nActual:\n\n%s\n\n", sliceA, sliceB, expected, actual)
	}
}

func TestPrettyPrintMultiIdentical(t *testing.T) {
	sl := []string{"foo", "foo", "foo"}
	testPrettyPrintMulti(t, "foo", sl, "foo", sl, "foo", sl, '-', 1, `foo foo foo
-----------
foo foo foo
foo foo foo
foo foo foo`)
}

func TestPrettyPrintMultiLeft(t *testing.T) {
	a := []string{"foo", "foo", "foo"}
	b := []string{"bar", "bar"}
	c := []string{"baz", "baz"}
	testPrettyPrintMulti(t, "foo", a, "bar", b, "baz", c, '-', 1, `foo bar baz
-----------
foo bar baz
foo bar baz
foo        `)
}

func TestPrettyPrintMultiMiddle(t *testing.T) {
	a := []string{"foo", "foo"}
	b := []string{"bar", "bar", "bar"}
	c := []string{"baz", "baz"}
	testPrettyPrintMulti(t, "foo", a, "bar", b, "baz", c, '-', 1, `foo bar baz
-----------
foo bar baz
foo bar baz
    bar    `)
}

func TestPrettyPrintMultiRight(t *testing.T) {
	a := []string{"foo", "foo"}
	b := []string{"bar", "bar"}
	c := []string{"baz", "baz", "baz"}
	testPrettyPrintMulti(t, "foo", a, "bar", b, "baz", c, '-', 1, `foo bar baz
-----------
foo bar baz
foo bar baz
        baz`)
}

func TestPrettyPrintMultiWider(t *testing.T) {
	sl := []string{"foo", "foo", "foo"}
	testPrettyPrintMulti(t, "foo", sl, "foo", sl, "foo", sl, '-', 5, `foo     foo     foo
-------------------
foo     foo     foo
foo     foo     foo
foo     foo     foo`)
}

func TestPrettyPrintMultiOtherSeparator(t *testing.T) {
	sl := []string{"foo", "foo", "foo"}
	testPrettyPrintMulti(t, "foo", sl, "foo", sl, "foo", sl, '=', 1, `foo foo foo
===========
foo foo foo
foo foo foo
foo foo foo`)
}

func TestPrettyPrintMultiLeftHeading(t *testing.T) {
	a := []string{"foo", "foo", "foo"}
	b := []string{"bar", "bar", "bar"}
	c := []string{"baz", "baz", "baz"}
	testPrettyPrintMulti(t, "alpha", a, "bar", b, "baz", c, '-', 1, `alpha bar baz
-------------
foo   bar baz
foo   bar baz
foo   bar baz`)
}

func TestPrettyPrintMultiMiddleHeading(t *testing.T) {
	a := []string{"foo", "foo", "foo"}
	b := []string{"bar", "bar", "bar"}
	c := []string{"baz", "baz", "baz"}
	testPrettyPrintMulti(t, "foo", a, "beta", b, "baz", c, '-', 1, `foo beta baz
------------
foo bar  baz
foo bar  baz
foo bar  baz`)
}

func TestPrettyPrintMultiRightHeading(t *testing.T) {
	a := []string{"foo", "foo", "foo"}
	b := []string{"bar", "bar", "bar"}
	c := []string{"baz", "baz", "baz"}
	testPrettyPrintMulti(t, "foo", a, "bar", b, "gamma", c, '-', 1, `foo bar gamma
-------------
foo bar baz  
foo bar baz  
foo bar baz  `)
}

func TestPrettyPrintMultiLeftWider(t *testing.T) {
	a := []string{"foo", "fooooo", "foo"}
	b := []string{"bar", "bar", "bar"}
	c := []string{"baz", "baz", "baz"}
	testPrettyPrintMulti(t, "foo", a, "bar", b, "baz", c, '-', 1, `foo    bar baz
--------------
foo    bar baz
fooooo bar baz
foo    bar baz`)
}

func TestPrettyPrintMultiMiddleWider(t *testing.T) {
	a := []string{"foo", "foo", "foo"}
	b := []string{"bar", "baaaar", "bar"}
	c := []string{"baz", "baz", "baz"}
	testPrettyPrintMulti(t, "foo", a, "bar", b, "baz", c, '-', 1, `foo bar    baz
--------------
foo bar    baz
foo baaaar baz
foo bar    baz`)
}

func TestPrettyPrintMultiRightWider(t *testing.T) {
	a := []string{"foo", "foo", "foo"}
	b := []string{"bar", "bar", "bar"}
	c := []string{"baz", "baaaaz", "baz"}
	testPrettyPrintMulti(t, "foo", a, "bar", b, "baz", c, '-', 1, `foo bar baz   
--------------
foo bar baz   
foo bar baaaaz
foo bar baz   `)
}

func testPrettyPrintMulti(t *testing.T, headingA string, sliceA []string, headingB string, sliceB []string, headingC string, sliceC []string, separator rune, spacing int, expected string) {
	actual := PrettyPrintMulti(separator, spacing, []string{headingA, headingB, headingC}, sliceA, sliceB, sliceC)

	if actual != expected {
		t.Errorf("Test failed.\nInput slice a: %#v\nInput slice b: %#v\nInput slice c: %#v\nExpected:\n\n%s\n\nActual:\n\n%s\n\n", sliceA, sliceB, sliceC, expected, actual)
	}
}

func TestTable(t *testing.T) {
	a := []string{"foooooo", "foo", "foo"}
	b := []string{"bar", "bar"}
	c := []string{"", "baaaaz", "baz"}
	d := []string{"baz", "baaaaz", "baz"}
	e := []string{"baz", "baaaaz", "baz"}

	expected := `Alpha    Beta  Gamma   Delta   Epsilon
======================================
foooooo  bar           baz     baz    
foo      bar   baaaaz  baaaaz  baaaaz 
foo            baz     baz     baz    `

	actual := PrettyPrintMulti('=', 2, []string{"Alpha", "Beta", "Gamma", "Delta", "Epsilon"}, a, b, c, d, e)

	if actual != expected {
		t.Errorf("Test failed.\nInput slice a: %#v\nInput slice b: %#v\nInput slice c: %#v\nInput slice d: %#v\nInput slice e: %#v\nExpected:\n\n%s\n\nActual:\n\n%s\n\n", a, b, c, d, e, expected, actual)
	}
}
