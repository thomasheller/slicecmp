package slicecmp

import (
	"fmt"
	"testing"
)

func TestCompare(t *testing.T) {
	testCases := []struct {
		a        []string
		b        []string
		expected bool
	}{
		{nil, nil, true},

		{nil, []string{"foo"}, false},

		{[]string{"foo"}, nil, false},

		{[]string{"foo", "bar"}, []string{"foo", "bar"}, true},

		{[]string{"foo", "foo"}, []string{"bar", "bar"}, false},

		{[]string{"foo", "foo"}, []string{"foo"}, false},

		{[]string{"foo"}, []string{"foo", "foo"}, false},

		{[]string{"foo", "foo", "foo"}, []string{"bar", "bar"}, false},

		{[]string{"foo", "foo"}, []string{"bar", "bar", "bar"}, false},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v %v", tc.a, tc.b), func(t *testing.T) {
			actual := Equal(tc.a, tc.b)

			if actual != tc.expected {
				t.Errorf("expected: %t, actual: %t\n", tc.expected, actual)
			}
		})
	}
}

func TestSprintf2Identical(t *testing.T) {
	sl := []string{"foo", "foo", "foo"}
	testSprintf2(t, "foo", sl, "foo", sl, '-', 1, `foo foo
-------
foo foo
foo foo
foo foo`)
}

func TestSprintf2Left(t *testing.T) {
	a := []string{"foo", "foo", "foo"}
	b := []string{"bar", "bar"}
	testSprintf2(t, "foo", a, "bar", b, '-', 1, `foo bar
-------
foo bar
foo bar
foo    `)
}

func TestSprintf2Right(t *testing.T) {
	a := []string{"foo", "foo"}
	b := []string{"bar", "bar", "bar"}
	testSprintf2(t, "foo", a, "bar", b, '-', 1, `foo bar
-------
foo bar
foo bar
    bar`)
}

func TestSprintf2Wider(t *testing.T) {
	sl := []string{"foo", "foo", "foo"}
	testSprintf2(t, "foo", sl, "foo", sl, '-', 5, `foo     foo
-----------
foo     foo
foo     foo
foo     foo`)
}

func TestSprintf2OtherSeparator(t *testing.T) {
	sl := []string{"foo", "foo", "foo"}
	testSprintf2(t, "foo", sl, "foo", sl, '=', 1, `foo foo
=======
foo foo
foo foo
foo foo`)
}

func TestSprintf2LeftHeading(t *testing.T) {
	a := []string{"foo", "foo", "foo"}
	b := []string{"bar", "bar", "bar"}
	testSprintf2(t, "alpha", a, "bar", b, '-', 1, `alpha bar
---------
foo   bar
foo   bar
foo   bar`)
}

func TestSprintf2RightHeading(t *testing.T) {
	a := []string{"foo", "foo", "foo"}
	b := []string{"bar", "bar", "bar"}
	testSprintf2(t, "foo", a, "beta", b, '-', 1, `foo beta
--------
foo bar 
foo bar 
foo bar `)
}

func TestSprintf2LeftWider(t *testing.T) {
	a := []string{"foo", "fooooo", "foo"}
	b := []string{"foo", "foo", "foo"}
	testSprintf2(t, "foo", a, "foo", b, '-', 1, `foo    foo
----------
foo    foo
fooooo foo
foo    foo`)
}

func TestSprintf2RightWider(t *testing.T) {
	a := []string{"foo", "foo", "foo"}
	b := []string{"foo", "fooooo", "foo"}
	testSprintf2(t, "foo", a, "foo", b, '-', 1, `foo foo   
----------
foo foo   
foo fooooo
foo foo   `)
}

func testSprintf2(t *testing.T, headingA string, sliceA []string, headingB string, sliceB []string, separator rune, spacing int, expected string) {
	actual := Sprintf2(headingA, sliceA, headingB, sliceB, separator, spacing)
	if actual != expected {
		t.Fatalf("Test failed.\nInput slice a: %#v\nInput slice b: %#v\nExpected:\n\n%s\n\nActual:\n\n%s\n\n", sliceA, sliceB, expected, actual)
	}
}

func TestSprintfIdentical(t *testing.T) {
	sl := []string{"foo", "foo", "foo"}
	testSprintf(t, "foo", sl, "foo", sl, "foo", sl, '-', 1, `foo foo foo
-----------
foo foo foo
foo foo foo
foo foo foo`)
}

func TestSprintfUnicode(t *testing.T) {
	a := []string{"föö", "föö", "foo"}
	b := []string{"foo", "foo", "fåå"}
	c := []string{"foo", "foo", "foo"}
	testSprintf(t, "foo", a, "foo", b, "foo", c, '-', 1, `foo foo foo
-----------
föö foo foo
föö foo foo
foo fåå foo`)
}

func TestSprintfLeft(t *testing.T) {
	a := []string{"foo", "foo", "foo"}
	b := []string{"bar", "bar"}
	c := []string{"baz", "baz"}
	testSprintf(t, "foo", a, "bar", b, "baz", c, '-', 1, `foo bar baz
-----------
foo bar baz
foo bar baz
foo        `)
}

func TestSprintfMiddle(t *testing.T) {
	a := []string{"foo", "foo"}
	b := []string{"bar", "bar", "bar"}
	c := []string{"baz", "baz"}
	testSprintf(t, "foo", a, "bar", b, "baz", c, '-', 1, `foo bar baz
-----------
foo bar baz
foo bar baz
    bar    `)
}

func TestSprintfRight(t *testing.T) {
	a := []string{"foo", "foo"}
	b := []string{"bar", "bar"}
	c := []string{"baz", "baz", "baz"}
	testSprintf(t, "foo", a, "bar", b, "baz", c, '-', 1, `foo bar baz
-----------
foo bar baz
foo bar baz
        baz`)
}

func TestSprintfWider(t *testing.T) {
	sl := []string{"foo", "foo", "foo"}
	testSprintf(t, "foo", sl, "foo", sl, "foo", sl, '-', 5, `foo     foo     foo
-------------------
foo     foo     foo
foo     foo     foo
foo     foo     foo`)
}

func TestSprintfOtherSeparator(t *testing.T) {
	sl := []string{"foo", "foo", "foo"}
	testSprintf(t, "foo", sl, "foo", sl, "foo", sl, '=', 1, `foo foo foo
===========
foo foo foo
foo foo foo
foo foo foo`)
}

func TestSprintfLeftHeading(t *testing.T) {
	a := []string{"foo", "foo", "foo"}
	b := []string{"bar", "bar", "bar"}
	c := []string{"baz", "baz", "baz"}
	testSprintf(t, "alpha", a, "bar", b, "baz", c, '-', 1, `alpha bar baz
-------------
foo   bar baz
foo   bar baz
foo   bar baz`)
}

func TestSprintfMiddleHeading(t *testing.T) {
	a := []string{"foo", "foo", "foo"}
	b := []string{"bar", "bar", "bar"}
	c := []string{"baz", "baz", "baz"}
	testSprintf(t, "foo", a, "beta", b, "baz", c, '-', 1, `foo beta baz
------------
foo bar  baz
foo bar  baz
foo bar  baz`)
}

func TestSprintfRightHeading(t *testing.T) {
	a := []string{"foo", "foo", "foo"}
	b := []string{"bar", "bar", "bar"}
	c := []string{"baz", "baz", "baz"}
	testSprintf(t, "foo", a, "bar", b, "gamma", c, '-', 1, `foo bar gamma
-------------
foo bar baz  
foo bar baz  
foo bar baz  `)
}

func TestSprintfLeftWider(t *testing.T) {
	a := []string{"foo", "fooooo", "foo"}
	b := []string{"bar", "bar", "bar"}
	c := []string{"baz", "baz", "baz"}
	testSprintf(t, "foo", a, "bar", b, "baz", c, '-', 1, `foo    bar baz
--------------
foo    bar baz
fooooo bar baz
foo    bar baz`)
}

func TestSprintfMiddleWider(t *testing.T) {
	a := []string{"foo", "foo", "foo"}
	b := []string{"bar", "baaaar", "bar"}
	c := []string{"baz", "baz", "baz"}
	testSprintf(t, "foo", a, "bar", b, "baz", c, '-', 1, `foo bar    baz
--------------
foo bar    baz
foo baaaar baz
foo bar    baz`)
}

func TestSprintfRightWider(t *testing.T) {
	a := []string{"foo", "foo", "foo"}
	b := []string{"bar", "bar", "bar"}
	c := []string{"baz", "baaaaz", "baz"}
	testSprintf(t, "foo", a, "bar", b, "baz", c, '-', 1, `foo bar baz   
--------------
foo bar baz   
foo bar baaaaz
foo bar baz   `)
}

func testSprintf(t *testing.T, headingA string, sliceA []string, headingB string, sliceB []string, headingC string, sliceC []string, separator rune, spacing int, expected string) {
	actual := Sprintf(separator, spacing, AlignLeft, []string{headingA, headingB, headingC}, sliceA, sliceB, sliceC)

	if actual != expected {
		t.Fatalf("Test failed.\nInput slice a: %#v\nInput slice b: %#v\nInput slice c: %#v\nExpected:\n\n%s\n\nActual:\n\n%s\n\n", sliceA, sliceB, sliceC, expected, actual)
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

	actual := Sprintf('=', 2, AlignLeft, []string{"Alpha", "Beta", "Gamma", "Delta", "Epsilon"}, a, b, c, d, e)

	if actual != expected {
		t.Fatalf("Test failed.\nInput slice a: %#v\nInput slice b: %#v\nInput slice c: %#v\nInput slice d: %#v\nInput slice e: %#v\nExpected:\n\n%s\n\nActual:\n\n%s\n\n", a, b, c, d, e, expected, actual)
	}
}

func TestSprintfFailsOnInvalidArguments(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Sprintf didn't panic as expected when called with different number of headings and slices.")
		}
	}()

	a := []string{"foo"}
	Sprintf('-', 1, AlignLeft, []string{"Alpha"}, a, a)
}

func TestAlignRight(t *testing.T) {
	a := []string{"foooooo", "foo", "foo"}
	b := []string{"bar", "bar"}
	c := []string{"", "baaaaz", "baz"}
	d := []string{"baz", "baaaaz", "baz"}
	e := []string{"baz", "baaaaz", "baz"}

	expected := `Alpha    Beta  Gamma   Delta   Epsilon
======================================
foooooo   bar             baz      baz
    foo   bar  baaaaz  baaaaz   baaaaz
    foo           baz     baz      baz`

	actual := Sprintf('=', 2, AlignRight, []string{"Alpha", "Beta", "Gamma", "Delta", "Epsilon"}, a, b, c, d, e)

	if actual != expected {
		t.Fatalf("Test failed.\nInput slice a: %#v\nInput slice b: %#v\nInput slice c: %#v\nInput slice d: %#v\nInput slice e: %#v\nExpected:\n\n%s\n\nActual:\n\n%s\n\n", a, b, c, d, e, expected, actual)
	}
}

func TestTransform(t *testing.T) {
	expected := [][]string{
		[]string{"A1", "B1", "C1"},
		[]string{"A2", "B2", "C2"},
		[]string{"A3", "B3", "C3"},
	}
	actual := Transform([][]string{
		[]string{"A1", "A2", "A3"},
		[]string{"B1", "B2", "B3"},
		[]string{"C1", "C2", "C3"},
	})
	for i := 0; i < 3; i++ {
		for k := 0; k < 3; k++ {
			if expected[i][k] != actual[i][k] {
				t.Fatalf("Row %d, column %d doesn't match: want %s, have %s", i, k, expected[i][k], actual[i][k])
			}
		}
	}
}
