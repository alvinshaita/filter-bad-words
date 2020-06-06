package main

import (
	"C"
	"fmt"
	"encoding/json"
	"os"
	"io/ioutil"
	"strconv"
)

var json_data map[string]map[string]int

//export initialize_go
func initialize_go(){
	//collects data.json content
	fmt.Println("Initializing server..")
	jsonFile, err := os.Open("data.json")
    if err != nil {}
    defer jsonFile.Close()
    jsonFile_data, _ := ioutil.ReadAll(jsonFile)
    json.Unmarshal([]byte(jsonFile_data), &json_data)
}

//export get_topics
func get_topics(word *C.char) *C.char {
	//finds the topics in each word received
	var word_topics = "{"
	for _key, _value := range json_data[C.GoString(word)] { 
	    word_topics+=`"`+_key+`": `+strconv.Itoa(_value)+`, `
	}
	
	if len(word_topics)>2{
		word_topics = word_topics[:len(word_topics)-2]
	}
	word_topics+="}"

	//returns topics in each word
	return C.CString(word_topics)
}


func main(){
}
