package parser

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func genReader(input string) *bufio.Reader {
	stringReader := strings.NewReader(input)
	rd := bufio.NewReader(stringReader)
	return rd
}

func TestBasicStringObj(t *testing.T) {
	input := "{\"name\": \"dhruv\", \"age\": \"20\"}"
	rd := genReader(input)

	output, err := ParseObject(rd)
	require.Nil(t, err)

	expected := Object{
		values: map[string]Value{
			"name": {typ: TypeString, str: "dhruv"},
			"age":  {typ: TypeString, str: "20"},
		},
	}
	require.Equal(t, expected, output)
}

func TestBasicArrayParse(t *testing.T) {
	input := "{\n  \"fruits\": [\n    \"apple\",\n    \"banana\",\n    \"orange\",\n    \"kiwi\"\n  ]\n}"
	rd := genReader(input)

	output, err := ParseObject(rd)
	require.Nil(t, err)

	expected := Object{
		values: map[string]Value{
			"fruits": {typ: TypeArray, array: []Value{
				{typ: TypeString, str: "apple"},
				{typ: TypeString, str: "banana"},
				{typ: TypeString, str: "orange"},
				{typ: TypeString, str: "kiwi"},
			}},
		},
	}
	require.Equal(t, expected, output)
}
