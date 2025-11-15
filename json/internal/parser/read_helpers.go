package parser

import (
	"bufio"
	"fmt"
	"unicode"
)

// return true if the next byte is what you expect.
func checkByte(rd *bufio.Reader, expected byte) (bool, error) {
	actual, err := rd.ReadByte()
	if err != nil {
		return false, err
	}
	rd.UnreadByte() // restore the byte
	fmt.Printf("checking for byte: %v | found byte: %v\n", string(expected), string(actual))
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

// checks the upcoming bytes to see what type of thing this is
func getValueType(rd *bufio.Reader) (ValueType, error) {
	b, err := rd.ReadByte()
	if err != nil {
		return TypeNull, err
	}
	rd.UnreadByte()

	switch b {
	case '"':
		return TypeString, nil
	case '[':
		return TypeArray, nil
	case '{':
		return TypeObject, nil
	case '-':
		return TypeNumber, nil
	case 't': // start of "true" or "false"
		return TypeBoolean, nil
	case 'f':
		return TypeBoolean, nil
	}

	// TODO: implement better checking for numbers

	return TypeNull, nil
}
