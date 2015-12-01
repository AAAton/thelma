package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"thelma/qsort"
)

var filename, conllFile string

func init() {
	flag.StringVar(&filename, "f", "", "help message for flagname")
	flag.StringVar(&conllFile, "c", "", "help message for flagname")
	flag.Parse()
}

func main() {

	if filename == "" {
		fmt.Println("You need to set filename with -f flag")
		os.Exit(0)
	}

	filename = cleanUpSymbols(filename)

	if conllFile == "" {
		conllFile = stagger(filename)

		namedEntityImprover(conllFile)
	}
	characterCount := findCharacterNames(conllFile)

	printCharacters(characterCount)

	tagCharactersInTextFile(filename, characterCount)

}

func stagger(filename string) string {
	fmt.Println("Running stagger...")

	conllFilename := "output/conll/" + getStoryName(filename) + ".conll"
	print := "java -jar stagger.jar -modelfile models/swedish.bin -tag " + filename + " > " + conllFilename
	ioutil.WriteFile("runStagger.sh", []byte(print), 0777)

	cmd := exec.Command("/bin/sh", "runStagger.sh")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	return conllFilename
}

func cleanUpSymbols(originalFile string) string {
	//TODO clean up symbols

	return originalFile
}

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

	taggedFile := "output/" + getStoryName(filename) + "_tagged.html"
	ioutil.WriteFile(taggedFile, []byte(filecontents), 0777)
	fmt.Println("Made", replacements, "replacements to", taggedFile)
}

func characterClassName(character string) string {
	class := strings.ToLower(character)
	class = strings.Replace(class, " ", "_", -1)
	return "<span class=\"" + class + "\">" + character + "</span>"
}

func sortCharacters(characterCount map[string]int) []string {
	var characters []string
	for character, count := range characterCount {
		if isCharacter(character, count) {
			characters = append(characters, character)
		}
	}

	return qsort.QuickSort(characters)
}

func getStoryName(filename string) string {
	storyName := strings.Replace(filename, ".txt", "", -1)
	indexOfSlash := strings.Index(storyName, "/")
	storyName = string([]byte(storyName)[indexOfSlash:])
	return storyName
}

func isCharacter(character string, count int) bool {
	return count > 4 && len(character) > 1
}
