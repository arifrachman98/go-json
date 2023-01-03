package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FName   string
	MName   string
	LName   string
	Age     int
	Married bool
}

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

func TestJSONObj(t *testing.T) {
	cust := Customer{
		FName:   "Arif",
		MName:   "Rachman",
		LName:   "Hakim",
		Age:     24,
		Married: false,
	}

	bytes, err := json.Marshal(cust)
	errHandler(err)
	fmt.Println(string(bytes))

}

func TestDecode(t *testing.T) {
	JSONstring := `{
		"FName" : "Arif",
		"MName" : "Rachman",
		"LName" : "Hakim",
		"Age" : 24,
		"Married" : false
	}`
	JSONbytes := []byte(JSONstring)

	cust := &Customer{}

	err := json.Unmarshal(JSONbytes, cust)
	errHandler(err)

	fmt.Println(cust)
	fmt.Println(cust.FName)
	fmt.Println(cust.MName)
	fmt.Println(cust.LName)
	fmt.Println(cust.Age)
	fmt.Println(cust.Married)
}
