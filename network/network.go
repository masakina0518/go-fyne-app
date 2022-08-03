package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	p := "https://tuyano-dummy-data.firebaseio.com/mydata.json"

	re, er := http.Get(p)

	if er != nil {
		panic(er)
	}

	defer re.Body.Close()

	s, er := ioutil.ReadAll(re.Body)
	if er != nil {
		panic(er)
	}

	var items []Mydata
	er = json.Unmarshal(s, &items)
	if er != nil {
		panic(er)
	}

	for i, im := range items {
		fmt.Println(i, im.Str())
	}
}

type Mydata struct {
	Name string
	Mail string
	Tel  string
}

func (m *Mydata) Str() string {
	return "<\"" + m.Name + "\" " + m.Mail + ", " + m.Tel + ">"
}
