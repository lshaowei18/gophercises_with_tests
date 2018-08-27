package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Data struct {
	Path string
	URL  string
}

func main() {
	//Read json file
	dataByte, err := ioutil.ReadFile("paths.json")
	if err != nil {
		log.Fatalln(err)
	}
	//Unmarshal
	var datas []Data
	err = json.Unmarshal(dataByte, &datas)
	if err != nil {
		log.Fatalln(err)
	}
}
