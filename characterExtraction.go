package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func findCharacterNames(filename string) map[string]int {
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
			characterName += vals[2] + " "
		} else if characterName != "" {
			personMap[strings.Trim(characterName, " ")]++
			characterName = ""
		}
	}
	return personMap
}

func printCharacters(characterCount map[string]int) {
	var characterList string
	for name, count := range characterCount {
		if count > 4 {
			characterList += name + "\t" + strconv.Itoa(count) + "\n"
		}
	}

	ioutil.WriteFile("output/"+getStoryName(filename)+"_characters.txt", []byte(characterList), 0777)
}
