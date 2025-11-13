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
	TypeNull valueType = 0
)

type Value struct {
	Typ valueType
	Str string
	Num int
	Bulk string
	Array []Value
}

type Resp struct {
	reader *bufio.Reader
	writer io.Writer
}

func NewResp(rd io.ReadWriteCloser) *Resp {
	return &Resp{
		reader: bufio.NewReader(rd),
		writer: rd}
}
