package structs

import (
	"fmt"
	"io/ioutil"
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

func (characters Characters) getByName(name string) *Character {
	var character *Character
	for _, c := range characters.List {
		if c.Name == name {
			character = c
			break
		}
	}
	return character
}

func (characters Characters) remove(index int) Characters {
	characters.List = append(characters.List[:index], characters.List[index+1:]...)
	return characters
}

//CreateCharacterListFromMap ...
func CreateCharacterListFromMap(characterMap map[string]int) Characters {
	//Converting map into list of characters
	var characters Characters
	i := 0

	for name, count := range characterMap {
		i++
		c := NewCharacter(i, name, count)
		characters.List = append(characters.List, &c)
	}

	characters = characters.clean()

	characters.linkAliases()
	characters.propegateLinks()

	return characters
}

func (characters Characters) linkAliases() {
	characters.singleLink()
	//characters.propegateLinks()
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
		character.setChainID(character.ID)
		for i := 0; i < len(character.LinkedIDs); i++ {
			ID := character.LinkedIDs[i]
			c := characters.get(ID)
			c.propegateLink(characters, character)
		}
	}
}

func (characters Characters) clean() Characters {
	deleted := true
	for deleted {
		deleted = false
		for index, character := range characters.List {
			//TODO Find a smarter way to do this condition
			if character.Count < 3 {
				characters = characters.remove(index)
				deleted = true
				break
			}
		}
	}
	return characters
}

//Print ...
func (characters Characters) Print(filename string) {
	var html string
	mainAliases := characters.getMainAliases()
	for _, c := range mainAliases.List {
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

func (characters Characters) getMainAliases() Characters {
	var mainAliases Characters
	for _, character := range characters.List {
		if !mainAliases.containsChain(character.ChainID) {
			mainAliases.List = append(mainAliases.List, characters.findMainAlias(character.ChainID))
		}
	}
	return mainAliases
}

func (characters Characters) containsChain(id int) bool {
	for _, character := range characters.List {
		if character.ChainID == id {
			return true
		}
	}
	return false
}

func (characters Characters) findMainAlias(chainID int) *Character {
	mainAlias := NewCharacter(-1, "", 0)

	for _, c := range characters.List {
		if c.ChainID == chainID && c.Count > mainAlias.Count {
			mainAlias = *c
		}
	}
	return &mainAlias
}
