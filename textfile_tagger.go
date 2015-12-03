package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"thelma/qsort"
)

func tagCharactersInTextFile(filename string, characterCount map[string]int) {
	fmt.Println("tagging characters in txt file")
	filecontents, _ := ioutil.ReadFile(filename)

	nonLetter := "\\P{L}"
	sortedCharacters := sortCharacters(characterCount)
	replacements := 0

	for _, character := range sortedCharacters {

		taggedCharacter := characterClassName(character)
		r, _ := regexp.Compile(nonLetter + character + nonLetter)
		indexes := r.FindAllIndex(filecontents, -1)
		if indexes != nil {
			for i := len(indexes) - 1; i >= 0; i-- {
				ending := append([]byte(taggedCharacter), filecontents[indexes[i][1]-1:]...)
				filecontents = append(filecontents[:indexes[i][0]+1], ending...)
				replacements++
			}
		}
	}

	taggedFile := "output/" + getStoryName() + "_tagged.html"
	ioutil.WriteFile(taggedFile, []byte(filecontents), 0777)
	fmt.Println("Made", replacements, "replacements to", taggedFile)
}

//TODO remove and replace with character.toHTML
func characterClassName(character string) string {
	class := strings.ToLower(character)
	class = strings.Replace(class, " ", "_", -1)
	return "<span class=\"" + class + "\">" + character + "</span>"
}

//TODO remove and create function within characters
func sortCharacters(characterCount map[string]int) []string {
	var characters []string
	for character, count := range characterCount {
		if isCharacter(character, count) {
			characters = append(characters, character)
		}
	}
	return qsort.QuickSort(characters)
}

//TODO remove and clean up character list in characters instead
func isCharacter(character string, count int) bool {
	return count > 4 && len(character) > 1
}
