package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Generic struct {
	Name string
	Cool bool
	Rank int
}

func main() {
	generic := Generic{"Golang", true, 100}
	fileJson, _ := json.Marshal(generic)
	err := ioutil.WriteFile("output.json", fileJson, 0644)
	if err != nil {
		fmt.Printf("WriteFileJson ERROR: %+v", err)
	}

	var fileGZ bytes.Buffer
	zipper := gzip.NewWriter(&fileGZ)
	defer zipper.Close()
	_, err = zipper.Write([]byte(string(fileJson)))
	if err != nil {
		fmt.Printf("zipper.Write ERROR: %+v", err)
	}
}
