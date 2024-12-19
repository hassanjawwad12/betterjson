package betterj

import (
	"testing"
)

//We defined the input and output format for the function we created
//as mentioned here: [Creating go module](https://dev.to/kingkunte_/go-modules-beginners-guide-4a7p#:~:text=In%20the%20Go%20programming%20language,in%20the%20module's%20root%20directory.)

func TestMinifyJ(t *testing.T) {
	input := `{
		"name": "John",
		"age": 30,
		"city": "New York"
	}`
	want := `{"name":"John","age":30,"city":"New York"}`

	t.Run(input, func(t *testing.T) {
		got, err := MinifyJ(input)
		if err != nil {
			t.Errorf("MinifyJ failed: %v", err)
		}
		if got != want {
			t.Errorf("expected %v, got %v", want, got)
		}
	})

}

func TestBeautifyJ(t *testing.T) {
	input := `{"name":"John","age":30,"city":"New York"}`
	indent := "  "
	want := `{
  "name": "John",
  "age": 30,
  "city": "New York"
}`

	t.Run(input, func(t *testing.T) {
		got, err := BeautifyJ(input, indent)
		if err != nil {
			t.Errorf("BeautifyJ failed: %v", err)
		}
		if got != want {
			t.Errorf("expected %v, got %v", want, got)
		}
	})

}
