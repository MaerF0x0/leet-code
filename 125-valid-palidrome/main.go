package main

import "fmt"

type tCase struct {
	in       string
	expected bool
}

func main() {
	printASCIIHints()
	tCases := []tCase{
		{
			in:       "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			in:       "race a car",
			expected: false,
		},
		{
			in:       "a|bbacba",
			expected: false, // c on right side
		},
		{
			in:       "a;b[b]a=c;b-a",
			expected: false, // c on right side after non-alphanum
		},
		{ // edge case, empty string is a palindrome
			in:       "| |",
			expected: true,
		},
		{ // single letter is palindrome
			in:       ";;;A ",
			expected: true,
		},
		{ // single letter is palindrome
			in:       "A ~!#@",
			expected: true,
		},
		{ // odd len Palindromes too
			in:       "abcdedcba ",
			expected: true,
		},
		{ // middle char is not alphanum
			in:       "abcd|dcba ",
			expected: true,
		},
		{ // oddlen w/ non alphanums
			in:       "abcd|edcba ",
			expected: true,
		},
		{ // oddlen w/ non alphanums
			in:       "abcde|dcba ",
			expected: true,
		},
		{ //  case insensitive palindrome
			in:       "a;;;A ",
			expected: true,
		},
		{ // accept numerics too
			in:       "a;10501;;A ", // a10501A
			expected: true,
		},
	}

	for _, tcase := range tCases {
		actual := isPalindrome(tcase.in)

		if actual != tcase.expected {
			fmt.Printf("FAILED -- actual=%t, expected=%t; in=%v\n", actual, tcase.expected, tcase.in)
		} else {
			fmt.Printf("PASSED -- actual=%t, expected=%t; in=%v\n", actual, tcase.expected, tcase.in)
		}
	}
}

func isPalindrome(s string) bool {
	ss := []rune(s)
	i, j := 0, len(ss)-1
	for i < j {
		if !isAlphaNumeric(ss[i]) {
			i++
			continue
		}
		if !isAlphaNumeric((ss[j])) {
			j--
			continue
		}

		if toLower(ss[i]) != toLower(ss[j]) {

			return false
		}
		i++
		j--
	}
	return true
}

func printASCIIHints() {
	cues := []rune{'0', '9', 'A', 'Z', 'a', 'z'}
	fmt.Println(string(cues))
	fmt.Println([]int32(cues))
}

func isAlphaNumeric(in rune) bool {
	return (ASCII_0 <= in && in <= ASCII_9) ||
		(ASCII_A <= in && in <= ASCII_Z) ||
		(ASCII_a <= in && in <= ASCII_z)
}

// toLower presumes you already only have a letter
func toLower(in rune) rune {
	if ASCII_A <= in && in <= ASCII_Z {
		return in + 32
	}
	return in
}

const (
	ASCII_0 = int32(48)
	ASCII_9 = int32(57)
	ASCII_A = int32(65)
	ASCII_Z = int32(90)
	ASCII_a = int32(97)
	ASCII_z = int32(122)
)
