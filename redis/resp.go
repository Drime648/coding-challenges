package main

import (
	"bufio"
	"io"
	"strconv"
	"fmt"
)


type valueType byte

const (
	TypeArray valueType = '*'
	TypeBulk valueType = '$'
	TypeString valueType = '+'
	TypeInt valueType = ':'
	TypeError valueType = '-'
)

type Value struct {
	typ valueType
	str string
	num int
	bulk string
	array []Value
}

type Resp struct {
	reader *bufio.Reader
}

func NewResp(rd io.Reader) *Resp {
	return &Resp{reader: bufio.NewReader(rd),}
}


func (r *Resp) readLine() ([]byte, int, error) {
	n := 0
	line := []byte{}
	for {
		byte, err := r.reader.ReadByte()
		if err != nil {
			return nil, 0, err
		}
		if byte == '\r' {
			//need to clear out the \r\n
			r.reader.ReadByte()
			break
		}
		n ++
		line = append(line, byte)
	}
	return line, n, nil
}



func (r *Resp) readInteger() (int, int, error) {
	line, n, err := r.readLine()
	if err != nil {
		return 0, 0, err
	}

	num, err := strconv.Atoi(string(line))
	if err != nil {
		return 0, 0, err
	}
	return num, n, nil
}


func (r *Resp) Read() (Value, error) {
	msg_type, err := r.reader.ReadByte()
	if err != nil {
		return Value{}, err
	}

	switch msg_type {
	case byte(TypeArray):
		return r.readArray()
	case byte(TypeBulk):
		return r.readBulk()
	default:
		return Value{}, fmt.Errorf("Unknown RESP type: %v", msg_type)
	}
}

func (r *Resp) readBulk() (Value, error){
	strLen, _, err := r.readInteger()
	// fmt.Printf("Bulk string len: %d\n", strLen)
	if err != nil {
		return Value{}, err
	}

	val := Value{typ: TypeBulk}

	bulkStr := make([]byte, strLen)
	_, err = r.reader.Read(bulkStr)
	if err != nil {
		return Value{}, err
	}

	val.bulk = string(bulkStr)
	// fmt.Printf("Bulk string: %s\n", val.bulk)

	r.readLine()//clear out any leftover from the string
	//							ex: they give big str but smaller len,
	//							we only read up to smaller length. Also
	//							clear out the \r\n

	return val, nil

}

func (r *Resp)readArray() (Value, error){
	arrayLen, _, err := r.readInteger()
	if err != nil {
		return Value{}, err
	}

	arr := make([]Value, arrayLen)

	for i := 0; i < arrayLen; i++ {
		v, err := r.Read()
		if err != nil {
			return Value{}, err
		}
		arr[i] = v
	}
	val := Value{}
	val.typ = TypeArray
	val.array = arr
	return val, nil
}
