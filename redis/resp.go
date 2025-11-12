package main

import (
	"bufio"
	"io"
	"strconv"
)

const (
	STRING = '+'
	ERROR = '-'
	INTEGER = ':'
	BULK = '$'
	ARRAY = '*'
)

type Value struct {
	typ string
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
