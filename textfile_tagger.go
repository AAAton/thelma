package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"thelma/structs"
)

func tagCharactersInTextFile(filename string, characters structs.Characters) {
	fmt.Println("tagging characters in txt file")
	filecontents, _ := ioutil.ReadFile(filename)

	nonLetter := "\\P{L}"

	replacements := 0

	for _, character := range characters.List {

		taggedCharacter := character.ToTag()
		r, _ := regexp.Compile(nonLetter + character.Name + nonLetter)
		indexes := r.FindAllIndex(filecontents, -1)
		if indexes != nil {
			for i := len(indexes) - 1; i >= 0; i-- {
				ending := append([]byte(taggedCharacter), filecontents[indexes[i][1]-1:]...)
				filecontents = append(filecontents[:indexes[i][0]+1], ending...)
				replacements++
			}
		}
	}

	taggedFile := "output/" + getStoryName() + ".html"
	ioutil.WriteFile(taggedFile, []byte(filecontents), 0777)
	fmt.Println("Made", replacements, "replacements to", taggedFile)
}
