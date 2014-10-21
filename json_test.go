package json

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnmarshalObject(t *testing.T) {
	type testcase struct {
		input  []byte
		output ObjectT
		err    bool
	}
	cases := []testcase{
		{[]byte(`{"test":"hello"}`), ObjectT{"test": Value("hello")}, false},
		{[]byte(`["123", "testtest"]`), nil, true},
	}

	for i, c := range cases {
		obj, err := UnmarshalObject(c.input)
		assert.Equal(t, c.output, obj, "case :%d", i)
		if c.err {
			assert.NotNil(t, err, "case :%d", i)
		} else {
			assert.Nil(t, err, "case :%d", i)
		}
	}
}

func TestUnmarshalArray(t *testing.T) {
	type testcase struct {
		input  []byte
		output ArrayT
		err    bool
	}
	cases := []testcase{
		{[]byte(`{"test":"hello"}`), nil, true},
		{[]byte(`["123", "testtest"]`), ArrayT{Value("123"), Value("testtest")}, false},
	}

	for i, c := range cases {
		obj, err := UnmarshalArray(c.input)
		assert.Equal(t, c.output, obj, "case :%d", i)
		if c.err {
			assert.NotNil(t, err, "case :%d", i)
		} else {
			assert.Nil(t, err, "case :%d", i)
		}
	}
}

func TestObjectMarshal(t *testing.T) {
	type testcase struct {
		input ObjectT
		expect []byte
	}
	cases := []testcase{
		{ObjectT{"test":Value("hello"), "test2": Value("testtest") }, []byte(`{"test":"hello","test2":"testtest"}`)},
		{ObjectT{"test":Value("hello"), "test2": Value(nil) }, []byte(`{"test":"hello","test2":null}`)},
		{ObjectT{"test":Value("hello"), "test2": CheckedValue{Value:"testtest", IsValid:false}}, []byte(`{"test":"hello"}`)},
		{ObjectT{"test":Value(ObjectT{"test":Value("test")})}, []byte(`{"test":{"test":"test"}}`)},
		{ObjectT{"test":Value(ArrayT{Value("hoge"),Value("hoge2")})}, []byte(`{"test":["hoge","hoge2"]}`)},
	}

	for i, c := range cases {
		out := c.input.Marshal()
		assert.Equal(t, c.expect, out, "case: %s", i)
	}
}

func TestArrayMarshal(t *testing.T) {
	type testcase struct {
		input ArrayT
		expect []byte
	}
	cases := []testcase{
		{ArrayT{Value("hello"), Value("testtest")}, []byte(`["hello","testtest"]`)},
		{ArrayT{Value("hello"), Value(nil)}, []byte(`["hello",null]`)},
		{ArrayT{Value("hello"), CheckedValue{Value:"testtest", IsValid:false}}, []byte(`["hello"]`)},
		{ArrayT{Value(ObjectT{"hello":Value("hihi")}), CheckedValue{Value:"testtest", IsValid:false}}, []byte(`[{"hello":"hihi"}]`)},
		{ArrayT{Value(ArrayT{Value("hihi")}), CheckedValue{Value:"testtest", IsValid:false}}, []byte(`[["hihi"]]`)},
	}

	for i, c := range cases {
		out := c.input.Marshal()
		assert.Equal(t, c.expect, out, "case: %s", i)
	}
}
