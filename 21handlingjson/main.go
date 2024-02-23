package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"` //omitempty omits the field in the struct if empty
}

func main() {
	fmt.Println("Welcome to the tutorial on handling JSON")
	// EncodeJson()
	DecodeJson()
}
func EncodeJson() {
	courses := []course{
		{"ReactJS Bootcamp", 288, "LearnCodeOnline.com", "abcd123", []string{"web-dev", "js", "reactJS"}},
		{"Mern Bootcamp", 300, "LearnCodeOnline.com", "abcd123", []string{"fullstack", "js", "reactJS", "mongo"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.com", "abcd123", nil},
	}

	//Package this data as json

	// finalJson, err := json.Marshal(courses)
	finalJson, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s \n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
			"coursename": "ReactJS Bootcamp",
			"price": 288,
			"website": "LearnCodeOnline.com",
			"tags": ["web-dev","js","reactJS"]
		}
	`)

	var course course

	isValidJson := json.Valid(jsonDataFromWeb)
	if isValidJson {
		fmt.Println("Json is valid")
		err := json.Unmarshal(jsonDataFromWeb, &course)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%#v\n", course)
	} else {
		fmt.Println("INVALID JSON")
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is %T\n", k, v, v)
	}
}
