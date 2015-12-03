package main

import (
	"strconv"
	"strings"
)

//Character contains information about each character name and aliases
type Character struct {
	ID        int
	Name      string
	Count     int
	LinkedIDs []int
}

func (c Character) toString() string {
	return strconv.Itoa(c.ID) + "\t" + c.Name + "\t" + strconv.Itoa(c.Count)
}

func (c Character) toHTML() string {
	//class names can't begin with a number
	classes := "char" + strconv.Itoa(c.ID)
	for _, id := range c.LinkedIDs {
		classes += " char" + strconv.Itoa(id)
	}
	return "<span class=\"" + classes + "\">" + c.Name + "</span>"
}

func (c Character) equals(c2 Character) bool {
	return c.ID == c2.ID
}

func (c Character) isLinkedTo(c2 Character) bool {
	for _, id := range c.LinkedIDs {
		for _, id2 := range c2.LinkedIDs {
			if id == id2 {
				return true
			}
		}
	}
	return false
}

func link(c, c2 Character) {
	c.LinkedIDs = append(c.LinkedIDs, c2.ID)
	c2.LinkedIDs = append(c2.LinkedIDs, c2.ID)
}

func (c Character) isSimilarTo(c2 Character) bool {

	//Exactly the same is not similar
	if c.equals(c2) {
		return false
	}

	//One name is a substring of the other
	if strings.Contains(c.Name, c2.Name) || strings.Contains(c2.Name, c.Name) {
		return true
	}

	//One name is possesive form of the other
	if strings.HasSuffix(c.Name, "s") && c.Name[:len(c.Name)-2] == c2.Name {
		return true
	}
	if strings.HasSuffix(c2.Name, "s") && c2.Name[:len(c2.Name)-2] == c.Name {
		return true
	}

	return false
}

func (c Character) propegateLink(characters Characters, originalChar Character) {

	if !c.isLinkedTo(originalChar) {
		link(c, originalChar)
	}

	for _, ID := range c.LinkedIDs {
		linkedChar := characters.get(ID)
		if !linkedChar.isLinkedTo(originalChar) {
			linkedChar.propegateLink(characters, originalChar)
		}
	}
}
