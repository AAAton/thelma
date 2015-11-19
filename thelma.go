package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var filename string
var conllFile string

func init() {
	flag.StringVar(&filename, "f", "", "help message for flagname")
	flag.StringVar(&conllFile, "c", "", "help message for conll")
	flag.Parse()
}

func main() {

	if filename == "" {
		fmt.Println("You need to set filename with -f flag")
		os.Exit(0)
	}
	if conllFile == "" {
		fmt.Println("You need to set conllFilename with -c flag")
		os.Exit(0)
	}

	filename = cleanUpSymbols(filename)

	//conllFile := stagger(filename)

	namedEntityImprover(conllFile)

	characterCount := findCharacterNames(conllFile)

	printCharacters(characterCount)

	tagCharactersInTextFile(filename, characterCount)

}

func stagger(filename string) string {
	//TODO create conll file with stagger
	conllFilename := "tmp.conll"
	cmd := "java -jar stagger.jar -modelfile models/swedish.bin > output/" + conllFilename
	err := exec.Command(cmd).Run()

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
	//TODO tag original textfile
}
