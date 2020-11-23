package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func main() {
	buf, err := ioutil.ReadFile("sample_data.yml")
	if err != nil {
		fmt.Println("Error 1")
		return
	}

	// convert to struct of map
	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(buf, &m)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", m["a"])

	fmt.Printf("---------------------------------------\n")	

	// convert struct of map to yaml
	y, err := yaml.Marshal(m)
	if err != nil {
		fmt.Println("Error")
		return
	}

	fmt.Printf(string(y))

	f_out(string(y))
}