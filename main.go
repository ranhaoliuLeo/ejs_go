package main

import (
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {

	m := map[string]string {
		"name": "ranhao",
		"age": "14",
	}
	TransTo("./html.ejs", "./", "index.html",m)
}


func TransTo(path string, desPath string, filename string,template map[string]string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	contentStr := string(content)
	if err != nil {
		panic(err)
	}
	var templateSlice []string
	reg := regexp.MustCompile(`<%=(.*?)=>`)
	match := reg.FindAllSubmatch(content, -1)
	for _, val := range match {
		for index ,val_ := range val {
			if index == 0 {
				continue
			} else {
				templateSlice = append(templateSlice, string(val_))
			}
		}
	}

	for _, key := range templateSlice {
		key = strings.TrimSpace(key)
		contentStr = strings.Replace(contentStr, "<%= "+ key +" =>", template[key], 1)
	}
	contentByte := []byte(contentStr)
	ioutil.WriteFile(desPath + filename, contentByte, 0666)
}