package utilities

import (
	"fmt"
	"testing"
	"time"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"
)

const (
	length = 14
)

func TestRandomCode_MakeBase2Code(t *testing.T) {
	res, _ := MakeBase2Code(time.Now().UnixNano(), length)
	t.Log("Result: " + res)
	assert.Equal(t, length, utf8.RuneCountInString(res), fmt.Sprint(res))
}

func TestRandomCode_MakeBase8Code(t *testing.T) {
	res, _ := MakeBase8Code(time.Now().UnixNano(), length)
	t.Log("Result: " + res)
	assert.Equal(t, length, utf8.RuneCountInString(res), fmt.Sprint(res))
}

func TestRandomCode_MakeBase10Code(t *testing.T) {
	res, _ := MakeBase10Code(time.Now().UnixNano(), length)
	t.Log("Result: " + res)
	assert.Equal(t, length, utf8.RuneCountInString(res), fmt.Sprint(res))
}

func TestRandomCode_MakeBase16Code(t *testing.T) {
	res, _ := MakeBase16Code(time.Now().UnixNano(), length)
	t.Log("Result: " + res)
	assert.Equal(t, length, utf8.RuneCountInString(res), fmt.Sprint(res))
}

func TestRandomCode_MakeBase32Code(t *testing.T) {
	res, _ := MakeBase32Code(time.Now().UnixNano(), length)
	t.Log("Result: " + res)
	assert.Equal(t, length, utf8.RuneCountInString(res), fmt.Sprint(res))
}

func TestRandomCode_MakeBase64Code(t *testing.T) {
	res, _ := MakeBase64Code(time.Now().UnixNano(), length)
	t.Log("Result: " + res)
	assert.Equal(t, length, utf8.RuneCountInString(res), fmt.Sprint(res))
}

func TestRandomCode_MakeEnAlphabetCaseSensitiveCode(t *testing.T) {
	res, _ := MakeEnAlphabetCaseSensitiveCode(time.Now().UnixNano(), length)
	t.Log("Result: " + res)
	assert.Equal(t, length, utf8.RuneCountInString(res), fmt.Sprint(res))
}

func TestRandomCode_MakeEnAlphabetCode(t *testing.T) {
	res, _ := MakeEnAlphabetCode(time.Now().UnixNano(), length)
	t.Log("Result: " + res)
	assert.Equal(t, length, utf8.RuneCountInString(res), fmt.Sprint(res))
}

func TestRandomCode_MakeEnAlphaNumericCaseSensitiveCode(t *testing.T) {
	res, _ := MakeEnAlphaNumericCaseSensitiveCode(time.Now().UnixNano(), length)
	t.Log("Result: " + res)
	assert.Equal(t, length, utf8.RuneCountInString(res), fmt.Sprint(res))
}

func TestRandomCode_MakeEnAlphaNumericCode(t *testing.T) {
	res, _ := MakeEnAlphaNumericCode(time.Now().UnixNano(), length)
	t.Log("Result: " + res)
	assert.Equal(t, length, utf8.RuneCountInString(res), fmt.Sprint(res))
}

func TestRandomCode_MakeFaAlphabetCode(t *testing.T) {
	res, _ := MakeFaAlphabetCode(time.Now().UnixNano(), length)
	t.Log("Result: " + res)
	assert.Equal(t, length, utf8.RuneCountInString(res), fmt.Sprint(res))
}

func TestRandomCode_MakeFaAlphaNumericCode(t *testing.T) {
	res, _ := MakeFaAlphaNumericCode(time.Now().UnixNano(), length)
	t.Log("Result: " + res)
	assert.Equal(t, length, utf8.RuneCountInString(res), fmt.Sprint(res))
}

func TestRandomCode_MakeFaNumericCode(t *testing.T) {
	res, _ := MakeFaNumericCode(time.Now().UnixNano(), length)
	t.Log("Result: " + res)
	assert.Equal(t, length, utf8.RuneCountInString(res), fmt.Sprint(res))
}
