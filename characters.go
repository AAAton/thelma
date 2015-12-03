package main

import "io/ioutil"

//Characters is a list of character
type Characters struct {
	List []Character
}

func (characters Characters) get(ID int) Character {
	var character Character
	for _, c := range characters.List {
		if c.ID == ID {
			character = c
			break
		}
	}
	return character
}

func createCharacterListFromMap(characterMap map[string]int) Characters {
	//Converting map into list of characters
	var characters Characters
	i := 0
	for name, count := range characterMap {
		i++
		characters.List = append(characters.List, Character{i, name, count, []int{}})
	}
	return characters
}

func (characters Characters) linkAliases() {
	characters.singleLink()
	characters.propegateLinks(characters.List[0])
}

func (characters Characters) singleLink() {
	for _, c := range characters.List {
		for _, c2 := range characters.List {
			if c.isSimilarTo(c2) && !c.isLinkedTo(c2) {
				link(c, c2)
			}
		}
	}
}

func (characters Characters) propegateLinks(character Character) {
	for _, ID := range character.LinkedIDs {
		characters.get(ID).propegateLink(characters, character)
	}
}

func (characters Characters) print(filename string) {
	var html string
	for _, c := range characters.List {
		html += c.toHTML() + "\n"
	}
	ioutil.WriteFile(filename, []byte(html), 0777)
}
