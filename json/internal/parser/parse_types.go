package parser

type Object struct {
	values map[string]Value
}

type ValueType string

const (
	TypeString ValueType = "string"
)


type Value struct {
	typ ValueType
	str string
}

type String struct {
	data string
}

