package main

import (
	"flag"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type TestInput struct {
	ToEncode string
	Encoded  string
}

func TestEncodeAndDecodeStruct(t *testing.T) {
	//given
	testInput := TestInput{
		"Hello",
		"JBSWY3DP",
	}

	assert := assert.New(t)
	assert.Equal(testInput.Encoded, encode([]byte(testInput.ToEncode)), "the encoded value is wrong")
	assert.Equal(testInput.ToEncode, decode([]byte(testInput.Encoded)), "the decoded value is wrong")

}

func TestEncodeAndDecodeAnonymous(t *testing.T) {

	testInput := struct {
		ToEncode string
		Encoded  string
	}{
		ToEncode: "Hello",
		Encoded:  "JBSWY3DP",
	}

	assert := assert.New(t)
	assert.Equal(testInput.Encoded, encode([]byte(testInput.ToEncode)), "the encoded value is wrong")
	assert.Equal(testInput.ToEncode, decode([]byte(testInput.Encoded)), "the decoded value is wrong")

}

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}
