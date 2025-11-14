package parser

import (
	"bufio"
	"fmt"
	// "io"
	// "unicode"
)

func ParseObject(rd *bufio.Reader) (Object, error) {
	obj := Object{values: map[string]Value{}}

	valid, err := checkByte(rd, '{')
	if err != nil {
		return obj, err
	}
	if !valid {
		return obj, fmt.Errorf("invalid start byte, must start with {")
	}
	rd.ReadByte() // consume the {
	err = clearWhitespace(rd)
	if err != nil {
		return obj, err
	}

	isEnd, err := checkByte(rd, '}')
	if err != nil {
		return obj, err
	}
	if isEnd {
		return obj, nil
	}

	for {

		clearWhitespace(rd)
		name, err := parseString(rd)
		if err != nil {
			return obj, err
		}
		fmt.Printf("name: %s\n", name)

		isColon, err := checkByte(rd, ':')
		if err != nil {
			return obj, err
		}
		if !isColon {
			return obj, fmt.Errorf("must separate name and value with :")
		}
		rd.ReadByte()
		clearWhitespace(rd)
		value, err := parseValue(rd)
		if err != nil {
			return obj, err
		}
		obj.values[name] = value

		isEnd, err = checkByte(rd, '}')
		if err != nil {
			return obj, err
		}
		if isEnd {
			break
		}

		isComma, err := checkByte(rd, ',')
		if err != nil {
			return obj, err
		}
		if !isComma {
			return obj, fmt.Errorf("invalid format for Object, needs , to separate values")
		}
		rd.ReadByte() // consume comma

	}
	return obj, nil
}

func parseString(rd *bufio.Reader) (string, error) {
	isString, err := checkByte(rd, '"')
	if err != nil {
		return "", err
	}
	if !isString {
		return "", fmt.Errorf("must have a string in Object, starting with \" ")
	}
	rd.ReadByte() // consume first quote

	resultStr := ""
	for {
		b, err := rd.ReadByte()
		if err != nil {
			return "", err
		}
		if b == '"' {
			break
		}
		resultStr += string(b)
	}

	return resultStr, nil
}

func parseValue(rd *bufio.Reader) (Value, error) {
	s, _ := parseString(rd)
	fmt.Printf("value: %s\n", s)
	return Value{str: s}, nil
}
