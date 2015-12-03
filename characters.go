package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

//Characters is a list of character
type Characters struct {
	List []*Character
}

func (characters Characters) get(ID int) *Character {
	var character *Character
	for _, c := range characters.List {
		if c.ID == ID {
			character = c
			break
		}
	}
	return character
}

func (characters Characters) remove(index int) {
	characters.List = append(characters.List[:index], characters.List[index+1:]...)
	characters.List[len(characters.List)-1] = nil
}

func createCharacterListFromMap(characterMap map[string]int) Characters {
	//Converting map into list of characters
	var characters Characters
	i := 0

	nonLetter := "\\P{L}"

	for name, count := range characterMap {
		i++
		r, _ := regexp.Compile(nonLetter + name + nonLetter)
		c := Character{i, name, count, []int{}, r}
		characters.List = append(characters.List, &c)
	}
	return characters
}

func (characters Characters) linkAliases() {
	characters.singleLink()
	characters.propegateLinks()
}

func (characters Characters) singleLink() {
	for _, c := range characters.List {
		for _, c2 := range characters.List {
			if (c.isSimilarTo(*c2) || c2.isSimilarTo(*c)) && !c.isLinkedTo(*c2) {
				link(c, c2)
			}
		}
	}
}

func (characters Characters) propegateLinks() {
	for _, character := range characters.List {
		for _, ID := range character.LinkedIDs {
			characters.get(ID).propegateLink(characters, character)
		}
	}
}

func (characters Characters) clean() {
	for index, character := range characters.List {
		if len(character.LinkedIDs) == 0 && character.Count < 3 {
			fmt.Println("removing", character.Name)
			characters.remove(index)
		}
	}
}

func (characters Characters) print(filename string) {
	var html string
	for _, c := range characters.List {
		html += c.toHTML() + "\n"
	}
	ioutil.WriteFile(filename, []byte(html), 0777)
}

//temp
func (characters Characters) controlCharacters() {
	for _, c := range characters.List {
		fmt.Println(c.LinkedIDs)
	}
}
