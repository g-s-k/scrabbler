package main

import "os"
import "fmt"
import "bufio"
import "strings"
import "regexp"
import "io/ioutil"
import "encoding/json"

const scoresFile = "scores.json"
const lang = "en"

func main() {
	// load score data from file
	rawScores, err := ioutil.ReadFile(scoresFile)
	scoresByLang := make(map[string](map[string]int))
	if err != nil {
		panic(err)
	}
	json.Unmarshal(rawScores, &scoresByLang)

	// scan stdin - TODO: get the words from here
	//s := bufio.NewScanner(os.Stdin)
	//for s.Scan() {
	//fmt.Println(s.Text())
	//}

	// load text from file
	readFile := os.Args[1]
	rawText, err := ioutil.ReadFile(readFile)
	if err != nil {
		panic(err)
	}

	// split into words
	var words []string
	scanner := bufio.NewScanner(strings.NewReader(string(rawText)))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	// clean up our strings
	// TODO: remove diacritical marks
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		panic(err)
	}
	for ind, el := range words {
		words[ind] = strings.ToUpper(reg.ReplaceAllString(el, ""))
	}

	// remove duplicates from list
	var wordsUnique []string
	for _, el := range words {
		if !isInList(el, wordsUnique) {
			wordsUnique = append(wordsUnique, el)
		}
	}

	// sort by score in descending order
	wordsSorted := wordsUnique
	for ind1, _ := range wordsUnique {
		smallest := ind1
		for ind2 := ind1; ind2 < len(wordsSorted); ind2++ {
			if scrabble(wordsSorted[smallest], scoresByLang[lang]) < scrabble(wordsSorted[ind2], scoresByLang[lang]) {
				smallest = ind2
			}
		}
		if smallest != ind1 {
			tmp := wordsSorted[ind1]
			wordsSorted[ind1] = wordsSorted[smallest]
			wordsSorted[smallest] = tmp
		}
	}

	for _, word := range wordsSorted {
		fmt.Printf("%d %s\n", scrabble(word, scoresByLang[lang]), word)
	}

	// print words with scores

	// compute total scrabble score

	// make game grid (?)
}

// function to test membership in a list of strings
func isInList(thing string, book []string) bool {
	for _, el := range book {
		if thing == el {
			return true
		}
	}
	return false
}

// function to add up letter scores
func scrabble(word string, scores map[string]int) int {
	score := 0
	for _, letter := range word {
		score += scores[string(letter)]
	}
	return score
}
