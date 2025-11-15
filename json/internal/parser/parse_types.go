package parser

type Object struct {
	values map[string]Value
}

type ValueType string

const (
	TypeString  ValueType = "string"
	TypeArray   ValueType = "array"
	TypeNumber  ValueType = "number"
	TypeObject  ValueType = "object"
	TypeBoolean ValueType = "boolean"
	TypeNull    ValueType = "null"
)

type Value struct {
	typ     ValueType
	str     string
	number  Number
	obj     Object
	array   []Value
	boolean bool
}

// TODO: Implement the Number struct.

type Number struct {
	data int
}
