package main

import (
	"regexp"
	"strconv"
	"strings"
)

//Character contains information about each character name and aliases
type Character struct {
	ID        int
	Name      string
	Count     int
	LinkedIDs []int
	regex     *regexp.Regexp
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

	if c.equals(c2) {
		return true
	}

	for _, id := range c.LinkedIDs {
		if c2.ID == id {
			return true
		}
	}
	for _, id := range c2.LinkedIDs {
		if c.ID == id {
			return true
		}
	}
	return false
}

func link(c, c2 *Character) {
	c.LinkedIDs = append(c.LinkedIDs, c2.ID)
	c2.LinkedIDs = append(c2.LinkedIDs, c.ID)
}

func (c Character) isSimilarTo(c2 Character) bool {

	//Exactly the same is not similar
	if c.equals(c2) {
		return false
	}

	//One name is a part of the other name
	if c.regex.MatchString(" " + c2.Name + " ") {
		return true
	}

	//One name is possesive form of the other
	if strings.HasSuffix(c.Name, "s") && c.Name[:len(c.Name)-2] == c2.Name {
		return true
	}

	return false
}

func (c Character) propegateLink(characters Characters, originalChar *Character) {

	if !c.isLinkedTo(*originalChar) {
		link(&c, originalChar)
	}

	for _, ID := range c.LinkedIDs {
		linkedChar := characters.get(ID)
		if !linkedChar.isLinkedTo(*originalChar) {
			linkedChar.propegateLink(characters, originalChar)
		}
	}
}
