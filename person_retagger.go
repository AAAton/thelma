package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var characters map[string]int

func retagPersonsInConll(conllFile string, characterCount map[string]int) {

	fmt.Println("retagging persons")

	characters = characterCount
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	data := string(dat)
	rows := strings.Split(data, "\n")

	for index, row := range rows {
		words := strings.Split(row, "\t")

		if isInList(rows, index) {
			if words[11] != "person" {
				words[11] = "person"
				row = ""
				for _, word := range words {
					row += word + "\t"
				}
				rows[index] = row
			}
		}
	}

	rewriteConll(conllFile, rows)
}

func rewriteConll(filename string, rows []string) {

	ioutil.WriteFile(filename, []byte(""), 0777)
	var lastPrint int
	var fileString string
	for index, row := range rows {
		fileString += row + "\n"

		//Emptying buffer every 10 000 rows
		if index/10000 > lastPrint {
			lastPrint = index / 10000
			fmt.Print(".")
			appendText(fileString, filename)
			fileString = ""
		}
	}
	appendText(fileString, filename)
	fmt.Println()
}

func isInList(rows []string, index int) bool {
	for name := range characters {
		broken := false
		wordsInName := strings.Split(name, " ")
		i := index
		for _, wordInName := range wordsInName {
			word := getWordFromRow(rows[i])
			if word != wordInName {
				broken = true
				break
			}
			i++
		}
		if !broken {
			return true
		}
	}
	return false
}

func getWordFromRow(row string) string {
	words := strings.Split(row, "\t")
	return words[1]
}
