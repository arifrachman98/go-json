package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

type Customer struct {
	FName     string
	MName     string
	LName     string
	Age       int
	Married   bool
	Hobbies   []string
	Addresses []Address
}

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"url"`
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

func TestEndcode(t *testing.T) {
	logJSON("Arif")
	logJSON(1)
	logJSON(true)
	logJSON([]string{"Arif", "Rachman", "Hakim"})
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

func TestJSONArrayEncode(t *testing.T) {
	cust := Customer{
		FName:   "Arif",
		MName:   "Torenoo",
		LName:   "Hakim",
		Age:     25,
		Married: true,
		Hobbies: []string{"Gimang", "Geming", "Gemung"},
	}

	bytes, err := json.Marshal(cust)
	errHandler(err)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	JSONstring := `{
		"FName" : "Arif",
		"MName" : "Rachman",
		"LName" : "Hakim",
		"Age" : 24,
		"Married" : false,
		"Hobbies" : ["Gem","Makan","Code"]
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
	fmt.Println(cust.Hobbies)
}

func TestJSONArrayComplexEncode(t *testing.T) {
	cust := Customer{
		FName:   "Arif",
		MName:   "Torenoo",
		LName:   "Hakim",
		Age:     25,
		Married: true,
		Hobbies: []string{"Gimang", "Geming", "Gemung"},
		Addresses: []Address{
			{
				Street:     "Jl. Bengawan",
				Country:    "Indonesia",
				PostalCode: "171745",
			},
			{
				Street:     "Lomere",
				Country:    "Koba",
				PostalCode: "100203",
			},
			{
				Street:     "Poerto",
				Country:    "Ricco",
				PostalCode: "14045",
			},
		},
	}

	bytes, err := json.Marshal(cust)
	errHandler(err)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	JSONstring := `{
		"FName" : "Arif",
		"MName" : "Rachman",
		"LName" : "Hakim",
		"Age" : 24,
		"Married" : false,
		"Hobbies" : ["Gem","Makan","Code"],
		"Addresses" : [
			{
				"Street" : "Jalanan",
				"Country" : "Lomestry",
				"PostalCode" : "44697"
			},{
				"Street" : "Jalanan yang jauh",
				"Country" : "Panama",
				"PostalCode" : "569778"
			},{
				"Street" : "Jalanan yang rusak",
				"Country" : "Armedia",
				"PostalCode" : "5489486"
			}]
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
	fmt.Println(cust.Hobbies)
	fmt.Println(cust.Addresses)
}

func TestJSONOnlyArrayComplexEncode(t *testing.T) {
	addr := []Address{
		{
			Street:     "Lomere",
			Country:    "Koba",
			PostalCode: "100203",
		},
		{
			Street:     "Poerto",
			Country:    "Ricco",
			PostalCode: "14045",
		},
	}

	bytes, err := json.Marshal(addr)
	errHandler(err)
	fmt.Println(string(bytes))
}

func TestJSONOnlyArrayComplexDecode(t *testing.T) {
	JSONstring := `[{
		"Street" : "Jalanan",
		"Country" : "Lomestry",
		"PostalCode" : "44697"
	},{
		"Street" : "Jalanan yang jauh",
		"Country" : "Panama",
		"PostalCode" : "569778"
	},{
		"Street" : "Jalanan yang rusak",
		"Country" : "Armedia",
		"PostalCode" : "5489486"
	}]`
	JSONbytes := []byte(JSONstring)

	addr := &[]Address{}
	err := json.Unmarshal(JSONbytes, addr)
	errHandler(err)

	fmt.Println(addr)
}

func TestJSONtagEncode(t *testing.T) {
	prod := Product{
		Id:       "P0001",
		Name:     "Laptop Gimang",
		ImageURL: "http://loremipsum.com",
	}

	bytes, err := json.Marshal(prod)
	errHandler(err)
	fmt.Println(string(bytes))
}

func TestJSONtagDecode(t *testing.T) {
	JSONstring := `{
		"id" : "D0001",	
		"name" : "Laptop buah",	
		"url" : "www.lingdua2.com"
	}`
	JSONbyte := []byte(JSONstring)

	prod := &Product{}
	err := json.Unmarshal(JSONbyte, prod)
	errHandler(err)

	fmt.Println(string(prod.Id))
	fmt.Println(string(prod.Name))
	fmt.Println(string(prod.ImageURL))
}

func TestMapEncode(t *testing.T) {
	prod := map[string]interface{}{
		"id":    "K0001",
		"name":  "MetaVersion 1.1.0",
		"price": 10000000,
	}

	bytes, err := json.Marshal(prod)
	errHandler(err)
	fmt.Println(string(bytes))
}

func TestMapDecode(t *testing.T) {
	JSONstring := `{
		"id" : "L0001",
		"name" : "Hape Terlalu Geming",
		"price" : "7000000"
	}`
	JSONbytes := []byte(JSONstring)

	var res map[string]interface{}
	err := json.Unmarshal(JSONbytes, &res)
	errHandler(err)

	fmt.Println(res)
	fmt.Println(res["id"])
	fmt.Println(res["name"])
	fmt.Println(res["price"])
}
