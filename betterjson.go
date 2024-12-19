package betterj

import (
	"bytes"
	"encoding/json"
)

// MinifyJ minifies a JSON string
func MinifyJ(j string) (string, error) {

	//empty buffer created
	var buffer bytes.Buffer //I used the bytes.buffer because its dynamic and works good with compact and indent

	//I have used the json.Compact here to remove the whitespace and the new lines
	//[]byte(j) here is a mutable byte slice
	if err := json.Compact(&buffer, []byte(j)); err != nil {
		return "", err
	}
	return buffer.String(), nil //converted the buffer content to string
}

// BeautifyJ beautifies a JSON string
func BeautifyJ(j string, indentWith string) (string, error) {
	var buffer bytes.Buffer
	//I have used the json.Indent here to format the JSON with custom indentation
	if err := json.Indent(&buffer, []byte(j), "", indentWith); err != nil {
		return "", err
	}
	return buffer.String(), nil
}
