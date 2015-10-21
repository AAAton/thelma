package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	tokens := Tokenize("Selma.txt")
	tokenCount := make(map[string]int)
	session := GetSession()
	c := session.DB("thelma").C("tokens")

	for _, token := range tokens {
		tokenCount[token]++
	}
	fmt.Println(len(tokenCount) / len(tokens))
	len := len(tokenCount)

	i := 0
	counter := 0
	c.DropCollection()
	for token, count := range tokenCount {
		if i%(len/100) == 0 {
			fmt.Println(counter, "%")
			counter++
		}
		c.Insert(bson.M{"token": token, "count": count})
		i++
	}

	/*err := c.Insert(bson.M{"token": token, "count": 0})
	if err != nil {
		c.Update(bson.M{"token": token}, bson.M{"$inc": bson.M{"count": 1}})
	}*/

}

//Tokenize splits text file into tokens
func Tokenize(filename string) []string {
	dat, _ := ioutil.ReadFile(filename)
	data := string(dat)
	data = strings.Replace(data, "\n", " ", -1)
	data = strings.Replace(data, "\t", " ", -1)
	data = strings.Replace(data, "-", " ", -1)
	data = strings.Replace(data, ".", "", -1)
	data = strings.Replace(data, ",", "", -1)
	data = strings.Replace(data, "!", "", -1)
	data = strings.Replace(data, "?", "", -1)

	for strings.Index(data, "  ") > -1 {
		data = strings.Replace(data, "  ", " ", -1)
	}

	data = strings.ToLower(data)
	tokens := strings.Split(data, " ")
	return tokens
}

//GetSession returns db session. Remember to defer
func GetSession() *mgo.Session {
	session, err := mgo.Dial("localhost")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	session.SetMode(mgo.Monotonic, true)
	return session
}
