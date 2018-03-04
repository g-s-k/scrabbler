package main

import "os"
import "fmt"
import "bufio"
import "strings"
import "io/ioutil"
import "encoding/json"

const scoresFile = "scores.json"

func main() {
	// load score data from file
	rawScores, err := ioutil.ReadFile(scoresFile)
	scoresByLang := make(map[string](map[string]int))
	if err != nil {
		panic(err)
	}
	json.Unmarshal(rawScores, &scoresByLang)

	// load text from file
	readFile := os.Args[1]
	rawText, err := ioutil.ReadFile(readFile)
	if err != nil {
		panic(err)
	}
	var words []string
	scanner := bufio.NewScanner(strings.NewReader(string(rawText)))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	fmt.Println(words)
}
