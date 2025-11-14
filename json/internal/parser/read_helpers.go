package parser

import (
	"bufio"
	"unicode"
)

// return true if the next byte is what you expect.
func checkByte(rd *bufio.Reader, expected byte) (bool, error) {
	actual, err := rd.ReadByte()
	if err != nil {
		return false, err
	}
	rd.UnreadByte() // restore the byte
	match := actual == expected
	return match, nil
}

//	func checkOpeningBrace(rd bufio.Reader) (bool, error) {
//		valid, err := checkByte(rd, '{')
//		if err != nil {
//			return obj, err
//		}
//		if !valid {
//			return obj, fmt.Errorf("Invalid start byte, must start with {")
//		}
//	}
//
// reads through all the whitespace till the next non-whitespace char, leaves that in buffer
func clearWhitespace(rd *bufio.Reader) error {
	for {
		r, _, err := rd.ReadRune()
		if err != nil {
			return err
		}
		if !unicode.IsSpace(r) {
			rd.UnreadRune()
			return nil
		}
	}
}
