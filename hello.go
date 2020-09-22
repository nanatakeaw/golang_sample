package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func main() {
	//hello()
	//a, b := twoReturn()
	//fmt.Println(a)
	//fmt.Println(b)

	//jsonString()

	jString := jsonString2()
	fmt.Println(jString)
}

func hello() {
	fmt.Println("Hello Go")
	fmt.Println("Narongsak Takaew")

	var x int = 15
	var s string = "Sawasdee"
	var f float32 = 10.56

	fmt.Println(x)
	fmt.Println(s)
	fmt.Println(f)
}
func twoReturn() (int, string) {

	return 1, "Super star"
}

type Person struct {
	Name string
	Age  int
}

func jsonString() {

	personJson1 := `{"name": "Narongsak","age": 39}`

	var p1 Person
	json.Unmarshal([]byte(personJson1), &p1)
	//fmt.Printf("Name: %s, Age: %d", p1.Name, p1.Age)

	personJson2 := `[{"name": "Narongsak","age": 39},{"name": "Coding Boy","age": 6}]`
	var p2 []Person
	json.Unmarshal([]byte(personJson2), &p2)
	//fmt.Printf("Person : %+v", p2)

	for i, data := range p2 {
		fmt.Println(i, data.Name)
	}
}
func jsonString2() string {

	var p [5]Person

	for i := 0; i < 5; i++ {
		p[i].Name = strconv.Itoa(i) + "XX"
		p[i].Age = i + 1
	}
	//fmt.Printf("Person : %+v", p)

	pJson, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(pJson)
}
