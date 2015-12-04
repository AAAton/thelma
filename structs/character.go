package structs

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
	ChainID   int
	regex     *regexp.Regexp
}

//NewCharacter is a constructor Contructor
func NewCharacter(id int, name string, count int) Character {
	nonLetter := "\\P{L}"
	r, _ := regexp.Compile("(?i)" + nonLetter + name + nonLetter)
	c := Character{id, name, count, []int{}, -1, r}
	return c
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
	return "<span id=\"char" + strconv.Itoa(c.ID) + "\" class=\"" + classes + "\" count=\"" + strconv.Itoa(c.Count) + "\">" + c.Name + "</span>"
}

//ToTag creates a simple tag to put in text
func (c Character) ToTag() string {
	return "<span class=\"char" + strconv.Itoa(c.ID) + "\">" + c.Name + "</span>"
}

func (c Character) equals(c2 Character) bool {
	return c.ID == c2.ID
}

func areLinked(c, c2 Character) bool {
	return c.isLinkedTo(c2) && c2.isLinkedTo(c)
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

	return false
}

func (c *Character) linkTo(c2 *Character) {
	c.LinkedIDs = append(c.LinkedIDs, c2.ID)
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
	if strings.HasSuffix(c.Name, "s") && c.Name[:len(c.Name)-1] == c2.Name {
		return true
	}

	return false
}

func (c *Character) propegateLink(characters Characters, originalChar *Character) {

	c.setChainID(originalChar.ChainID)

	if !originalChar.isLinkedTo(*c) {
		originalChar.linkTo(c)
	}

	for _, ID := range c.LinkedIDs {
		linkedChar := characters.get(ID)
		if linkedChar != nil && !originalChar.isLinkedTo(*linkedChar) {
			linkedChar.propegateLink(characters, originalChar)
		}
	}
}

func (c *Character) setChainID(id int) {
	if c.ChainID < 0 {
		c.ChainID = id
	}
}
