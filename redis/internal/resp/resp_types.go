package resp

import (
	"io"
	"bufio"
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
