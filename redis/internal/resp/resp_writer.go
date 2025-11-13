package resp

import (
	"strconv"
)


func (r *Resp) Write(v Value) error {
	bytes := v.Marshal()

	_, err := r.writer.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}


func (v Value) Marshal() []byte {
	switch v.Typ {
	case TypeArray:
		return v.marshalArray()
	case TypeBulk:
		return v.marshalBulk()
	case TypeString:
		return v.marshalString()
	case TypeInt:
		return v.marshalInt()
	case TypeError:
		return v.marshalError()
	case TypeNull:
		return v.marshalNull()
	}
	return []byte{}
}


func (v Value) marshalString() []byte {
	bytes := []byte{byte(TypeString)}
	bytes = append(bytes, []byte(v.Str)...)
	bytes = append(bytes, '\r', '\n')
	return bytes
}

func (v Value) marshalInt() []byte {
	bytes := []byte{byte(TypeString)}
	if v.Num < 0 {
		bytes = append(bytes, '-')
	}
	numStr := strconv.Itoa(v.Num)
	bytes = append(bytes, []byte(numStr)...)
	bytes = append(bytes, '\r', '\n')
	return bytes
}


func (v Value) marshalBulk() []byte {
	bytes := []byte{byte(TypeBulk)}
	lenStr := strconv.Itoa(len(v.Bulk))
	bytes = append(bytes, []byte(lenStr)...)
	bytes = append(bytes, '\r', '\n')
	bytes = append(bytes, []byte(v.Bulk)...)
	bytes = append(bytes, '\r', '\n')
	return bytes
}

func (v Value) marshalArray() []byte {
	bytes := []byte{byte(TypeArray)}
	lenArr := strconv.Itoa(len(v.Array))
	bytes = append(bytes, []byte(lenArr)...)

	for _, val := range v.Array {
		bytes = append(bytes, val.Marshal()...)
	}
	return bytes
}

func (v Value) marshalNull() []byte {
	return []byte("$-1\r\n")
}

func (v Value) marshalError() []byte {
	bytes := []byte{byte(TypeError)}
	bytes = append(bytes, []byte(v.Str)...)
	bytes = append(bytes, '\r', '\n')
	return bytes
}
