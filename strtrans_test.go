package strtrans

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/frankMilde/rol/testutils"
)

func Test_LinebreaksToTwoLinebreaks_haveMultipleLinebreaks_reducedToTwoLinebreaks(t *testing.T) {
	tests := testutils.ConversionTests{
		{
			In: `a
b
c


d




e





`,
			Want: `a
b
c

d

e

`,
		},
	}

	for _, test := range tests {
		got := LinebreaksToTwoLinebreaks(test.In)
		if !reflect.DeepEqual(test.Want, got) {
			_, file, line, _ := runtime.Caller(0)
			fmt.Printf("%s:%d:\n\ncall LineBreaks(%#v)\n\texp: %#v\n\n\tgot: %#v\n\n",
				filepath.Base(file), line, test.In, test.Want, got)
			t.FailNow()
		}
	}
	testutils.Cleanup()

}

func Test_LinebreaksToSpace_haveMultipleLinebreaks_reducedToSpace(t *testing.T) {
	tests := testutils.ConversionTests{
		{
			In: `a
b
c


d




e





`,
			Want: `a b c d e `,
		},
	}

	for _, test := range tests {
		got := LinebreaksToSpace(test.In)
		if !reflect.DeepEqual(test.Want, got) {
			_, file, line, _ := runtime.Caller(0)
			fmt.Printf("%s:%d:\n\ncall LineBreaksToSpace(%#v)\n\texp: %#v\n\n\tgot: %#v\n\n",
				filepath.Base(file), line, test.In, test.Want, got)
			t.FailNow()
		}
	}
	testutils.Cleanup()

}

func Test_BrHtmlTag(t *testing.T) {
	tests := testutils.ConversionTests{
		{
			In: `a<br/>
			b<br/>
			c`,
			Want: `a \\ 
			b \\ 
			c`,
		},
		{
			In: `a<br>
			b<br>
			c`,
			Want: `a \\ 
			b \\ 
			c`,
		},
	}

	for _, test := range tests {
		got := BrHtmlTagToLatexLinebreak(test.In)
		if !reflect.DeepEqual(test.Want, got) {
			_, file, line, _ := runtime.Caller(0)
			fmt.Printf("%s:%d:\n\ncall BrHtmlTag(%#v)\n\texp: %#v\n\n\tgot: %#v\n\n",
				filepath.Base(file), line, test.In, test.Want, got)
			t.FailNow()
		}
	}
	testutils.Cleanup()

}

func Test_Spaces_haveSpacesBeforeLinebreaks_spaceAreRemoved(t *testing.T) {
	tests := testutils.ConversionTests{
		{
			In:   `a   `,
			Want: `a`,
		},
	}

	for _, test := range tests {
		got := Spaces(test.In)
		if !reflect.DeepEqual(test.Want, got) {
			_, file, line, _ := runtime.Caller(0)
			fmt.Printf("%s:%d:\n\ncall Spaces(%#v)\n\texp: %#v\n\n\tgot: %#v\n\n",
				filepath.Base(file), line, test.In, test.Want, got)
			t.FailNow()
		}
	}
	testutils.Cleanup()

}

func Test_AllSubStrings_haveSpacesBeforeFootnotes_spaceAreRemoved(t *testing.T) {
	tests := testutils.ConversionTests{
		{
			In: `test. \footnote{text} Test

test. \footnote{text} Test

test. \footnote{text} Test

test. \footnote{text} Test`,
			Want: `test.\footnote{text} Test

test.\footnote{text} Test

test.\footnote{text} Test

test.\footnote{text} Test`,
		},
	}

	for _, test := range tests {
		got := AllSubStrings(test.In, `(\s+\\footnote)`, "\\footnote", 1)
		if !reflect.DeepEqual(test.Want, got) {
			_, file, line, _ := runtime.Caller(0)
			fmt.Printf("%s:%d:\n\ncall AllSubStrings(%#v)\n\texp: %#v\n\n\tgot: %#v\n\n",
				filepath.Base(file), line, test.In, test.Want, got)
			t.FailNow()
		}
	}
	testutils.Cleanup()

}
