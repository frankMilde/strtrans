// strtrans provides routines to transform strings
package strtrans

import (
	"log"
	"regexp"
	"strings"

	"github.com/frankMilde/strdel"
)

// MultiToSingleSpaces transforms
// (1) unicode whitespaces into ascii white space,
// (2) multiple white spaces into a single white space,
// (3) removes white space from the start and end of string s.
func MultiToSingleSpaces(s string) string {
	s = UnicodeToAsciiSpaces(s)
	regWhiteSpace := regexp.MustCompile(`\s{2,}`)
	s = regWhiteSpace.ReplaceAllString(s, " ")

	return strings.TrimSpace(s)
}

// UnicodeToAsciiSpaces transforms all unicode separator and space
// characters from unicode categroy code [Zs] into the simple ascii spaces.
// Also see http://www.fileformat.info/info/unicode/category/Zs/list.htm
func UnicodeToAsciiSpaces(s string) string {
	s = strings.Replace(s, "\u00A0", " ", -1) // \u00A0 = &nbsp = 'non-breaking space'
	s = strings.Replace(s, "\u1680", " ", -1) // ogham space mark
	s = strings.Replace(s, "\u2000", " ", -1) // en quad
	s = strings.Replace(s, "\u2001", " ", -1) // em quad
	s = strings.Replace(s, "\u2002", " ", -1) // en space
	s = strings.Replace(s, "\u2003", " ", -1) // em space
	s = strings.Replace(s, "\u2004", " ", -1) // three-per-em space
	s = strings.Replace(s, "\u2005", " ", -1) // four-per-em space
	s = strings.Replace(s, "\u2006", " ", -1) // six-per-em space
	s = strings.Replace(s, "\u2007", " ", -1) // figure spacE
	s = strings.Replace(s, "\u2008", " ", -1) // punctuation space
	s = strings.Replace(s, "\u2009", " ", -1) // thin space
	s = strings.Replace(s, "\u200A", " ", -1) // hair space
	s = strings.Replace(s, "\u202F", " ", -1) // narrow no-break space
	s = strings.Replace(s, " ", " ", -1)      // narrow no-break space
	s = strings.Replace(s, "\u205F", " ", -1) // medium mathematical space
	s = strings.Replace(s, "\u3000", " ", -1) // ideographic space

	return s
}

// LinebreaksToTwoLinebreaks will reduce all multiple newlines into 2
// newlines.
func LinebreaksToTwoLinebreaks(s string) string {
	s = strdel.TrailingSpaces(s)

	//s = strings.Replace(s, " \n", "\n", -1)
	//s = strings.Replace(s, "\t\n", "\n", -1)
	//s = strings.Replace(s, "\f\n", "\n", -1)
	//s = strings.Replace(s, "\r\n", "\n", -1)

	// trim multipple linebreaks to two linebreaks
	multiLinebreak := `(\r\n|\n){3,}`
	regMultiLinebreak := regexp.MustCompile(multiLinebreak)
	s = regMultiLinebreak.ReplaceAllString(s, "\n\n")

	//s = strings.Replace(s, "\n\n\n\n\n\n\n\n\n", "\n\n", -1)
	//s = strings.Replace(s, "\n\n\n\n\n\n\n\n", "\n\n", -1)
	//s = strings.Replace(s, "\n\n\n\n\n\n\n", "\n\n", -1)
	//s = strings.Replace(s, "\n\n\n\n\n\n", "\n\n", -1)
	//s = strings.Replace(s, "\n\n\n\n\n", "\n\n", -1)
	//s = strings.Replace(s, "\n\n\n\n", "\n\n", -1)
	//s = strings.Replace(s, "\n\n\n", "\n\n", -1)

	return s
}

// LinebreaksToSpaces will reduce all newlines into spaces.
func LinebreaksToSpace(s string) string {
	s = strdel.TrailingSpaces(s)

	multiLinebreak := `(\r\n|\n)+`
	regMultiLinebreak := regexp.MustCompile(multiLinebreak)
	s = regMultiLinebreak.ReplaceAllString(s, " ")

	return s
}

// BrHtmlTagToLatexLinebreak will transform all br tags into latex \\
// linbreaks..
func BrHtmlTagToLatexLinebreak(s string) string {
	s = strdel.TrailingSpaces(s)

	br := `<br>|<br\/>`
	regBr := regexp.MustCompile(br)
	s = regBr.ReplaceAllString(s, ` \\ `)

	return s
}

// AllSubStrings will replace all 'pattern' matches at
// 'submatchIndexTReplace'  of 'input' by 'replace'
func AllSubStrings(input string, pattern string, replace string, submatchIndexToReplace int) string {
	return SubString(input, pattern, replace, submatchIndexToReplace, -1)
}

// SubStrings will replace the nth `occurence` of 'pattern' matches at
// 'submatchIndexTReplace'  of 'input' by 'replace'
func SubString(input string, pattern string, replace string, submatchIndexToReplace int, occurence int) string {
	re := regexp.MustCompile(pattern)

	allmatches := re.FindAllStringSubmatch(input, occurence)

	for _, match := range allmatches {
		input = strings.Replace(input, match[submatchIndexToReplace], replace, -1)
	}

	return input
}

// AllButMatches transforms all parts of input string that do not belong to
// the list of matches.
// It splits input string along the slice of matches and applies the
// transform function on the splitted parts between the matches, effectivly
// ignoring the matches itself. The complete string is then put back
// together and returned.
// Should a match not be contained is input and empty string is returned.
func AllButMatches(input string, matches []string, transform func(string) string) string {

	if len(matches) == 0 {
		return transform(input)
	}

	var output string
	var after string = input

	for _, match := range matches {
		if !strings.Contains(after, match) {
			log.Println("Error: |" + match + "|\n\nnot found in\n\n|" + after + "|\n")
			return ""
		}
		split := strings.SplitN(after, match, 2)
		before := split[0]
		after = split[1]
		output = output + transform(before) + match
	}
	return output + transform(after)
}
