package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
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
		conllFile = tagWithStagger(filename)
		namedEntityImprover(conllFile)
	}

	characterCount := findFullCharacterNames(conllFile)

	characters := createCharacterListFromMap(characterCount)
	characters.linkAliases()
	characters.clean()
	characters.print("output/characters/" + getStoryName() + ".html")

	//tagCharactersInTextFile(filename, characterCount)

}

func tagWithStagger(filename string) string {
	fmt.Println("Running stagger...")

	conllFilename := "output/conll/" + getStoryName() + ".conll"
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

func getStoryName() string {
	storyName := strings.Replace(filename, ".txt", "", -1)
	indexOfSlash := strings.Index(storyName, "/")
	storyName = string([]byte(storyName)[indexOfSlash:])
	return storyName
}
