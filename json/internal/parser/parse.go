package parser


import (
	"bufio"
	"io"
	"fmt"
	"unicode"
)


func ParseJson(rd bufio.Reader) (Object, error) {

	obj := Object{}

	valid, err := checkByte(rd, '{')
	if err != nil {
		return obj, err
	}
	if !valid {
		return obj, fmt.Errorf("Invalid start byte, must start with {")
	}



}


//return true if the next byte is what you expect.
func checkByte(rd bufio.Reader, expected byte) (bool, error) {
	actual, err := rd.ReadByte()
	if err != nil {
		return false, err
	}
	match := actual == expected
	return match, nil

}



//reads through all the whitespace till the next non-whitespace char, leaves that in buffer
func clearWhitespace(rd bufio.Reader) error {
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
