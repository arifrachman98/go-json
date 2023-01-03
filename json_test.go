package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func errHandler(err error) {
	if err != nil {
		panic(err)
	}
}

func logJSON(data interface{}) {
	bytes, err := json.Marshal(data)
	errHandler(err)

	fmt.Println(string(bytes))
}

func TestEndcode(t *testing.T) {
	logJSON("Arif")
	logJSON(1)
	logJSON(true)
	logJSON([]string{"Arif", "Rachman", "Hakim"})
}
