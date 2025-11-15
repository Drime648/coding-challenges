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

		fmt.Println(obj)
		isComma, err := checkByte(rd, ',')
		if err != nil {
			return obj, err
		}
		if isComma {
			rd.ReadByte() // consume comma
			continue
		}

		clearWhitespace(rd)
		isEnd, err = checkByte(rd, '}')
		if err != nil {
			return obj, err
		}
		if isEnd {
			fmt.Println("found end of object")
			break
		}

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

	fmt.Printf("string:%s\n", resultStr)
	return resultStr, nil
}

func parseValue(rd *bufio.Reader) (Value, error) {
	val := Value{}
	valueType, err := getValueType(rd)
	if err != nil {
		return val, err
	}

	val.typ = valueType

	switch valueType {
	case TypeString:
		val.str, err = parseString(rd)
	case TypeArray:
		val.array, err = parseArray(rd)
	case TypeObject:
		val.obj, err = ParseObject(rd)
	}

	if err != nil {
		return val, err
	}

	return val, nil
}

func parseArray(rd *bufio.Reader) ([]Value, error) {
	arr := make([]Value, 0)
	isArray, err := checkByte(rd, '[')
	if err != nil {
		return arr, err
	}
	if !isArray {
		return arr, fmt.Errorf("invalid array format")
	}
	rd.ReadByte() // clear out the [

	err = clearWhitespace(rd)
	if err != nil {
		return arr, err
	}

	isEnd, err := checkByte(rd, ']')
	if err != nil {
		return arr, err
	}
	if isEnd {
		return arr, nil
	}

	for {
		clearWhitespace(rd)
		value, err := parseValue(rd)
		if err != nil {
			return arr, err
		}
		arr = append(arr, value)
		isComma, err := checkByte(rd, ',')
		if err != nil {
			return arr, err
		}
		if isComma {
			rd.ReadByte() // consume comma
			continue
		}

		clearWhitespace(rd)

		isEnd, err = checkByte(rd, ']')
		if err != nil {
			return arr, err
		}
		if isEnd {
			rd.ReadByte() // clear our ]
			fmt.Println("Found end of array")
			break
		}

		return arr, fmt.Errorf("invalid format for Object, needs , to separate values")
	}
	return arr, nil
}
