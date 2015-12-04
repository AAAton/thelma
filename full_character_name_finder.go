package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func findFullCharacterNames(filename string) map[string]int {
	var characterName string
	personMap := make(map[string]int)

	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	data := string(dat)
	rows := strings.Split(data, "\n")
	//foreach row in rows

	fmt.Println("extracting characters")
	for _, row := range rows {
		if strings.Contains(row, "person") {
			vals := strings.Split(row, "\t")
			characterName += vals[1] + " "
		} else if characterName != "" {
			personMap[strings.Trim(characterName, " ")]++
			characterName = ""
		}
	}
	return personMap
}
