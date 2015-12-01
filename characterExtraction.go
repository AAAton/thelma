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
	var characterCountString, characterList string
	for name, count := range characterCount {
		if isCharacter(name, count) {
			characterCountString += name + "\t" + strconv.Itoa(count) + "\n"
			characterList += characterClassName(name) + "\n"
		}
	}

	ioutil.WriteFile("output/"+getStoryName(filename)+"_characters.html", []byte(characterList), 0777)
	ioutil.WriteFile("output/"+getStoryName(filename)+"_character_count.txt", []byte(characterCountString), 0777)
}
