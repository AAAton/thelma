package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var personMap, placeMap, instMap map[string]int

func namedEntityImprover(filename string) {
	var persons, places, insts string

	personMap = make(map[string]int)
	placeMap = make(map[string]int)
	instMap = make(map[string]int)

	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	data := string(dat)
	rows := strings.Split(data, "\n")
	//foreach row in rows

	fmt.Println("parsing rows")
	for _, row := range rows {
		if strings.Contains(row, "person") {
			persons = persons + row + "\n"
			addToMap(personMap, row)
		} else if strings.Contains(row, "place") {
			places = places + row + "\n"
			addToMap(placeMap, row)
		} else if strings.Contains(row, "inst") {
			insts = insts + row + "\n"
			addToMap(instMap, row)
		}
	}
	fmt.Println("Removing unwanted persons/places")
	for person, count := range personMap {
		if placeMap[person] > count {
			delete(personMap, person)
		} else {
			delete(placeMap, person)
		}

		if instMap[person] > count {
			delete(personMap, person)
		} else {
			delete(instMap, person)
		}
	}

	fmt.Println("Creating list of characters")
	var personList string
	for person, count := range personMap {
		personList += person + "\t" + strconv.Itoa(count) + "\n"
	}

	fmt.Println("Creating list of places")

	var placeList string
	for place, count := range placeMap {
		placeList += place + "\t" + strconv.Itoa(count) + "\n"
	}

	correctSelma(filename, rows)
}

func addToMap(namedMap map[string]int, row string) {
	words := strings.Split(row, "\t")
	if len(words) > 1 {
		namedMap[words[2]]++
	}
}

func correctSelma(filename string, rows []string) {

	fmt.Println("Correcting Selma")

	ioutil.WriteFile(filename, []byte(""), 0777)
	var lastPrint int
	var fileString string
	for index, row := range rows {
		if strings.Contains(row, "person") || strings.Contains(row, "place") || strings.Contains(row, "inst") {
			words := strings.Split(row, "\t")
			if len(words) > 1 {
				name := words[2]
				tag := words[11]
				if personMap[name] > 0 {
					if tag != "person" {
						words[11] = "person"
					}
				}
				if placeMap[name] > 0 {
					if tag != "place" {
						words[11] = "place"
					}
				}
				if tag == "inst" {
					words[11] = "_"
				}
			}

			//Correcting row
			row = ""
			for _, word := range words {
				row += word + "\t"
			}
		}
		fileString += row + "\n"

		//Emptying buffer every 10 000 rows
		if index/10000 > lastPrint {
			lastPrint = index / 10000
			fmt.Println(index)
			append(fileString, filename)
			fileString = ""
		}
	}
	append(fileString, filename)
}

func append(text, filename string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}
}
