package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

var anagrams = readDict()

type stats struct {
	words int
	keys  int
}

func main() {
	listen()
}

func listen() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a word then enter: ")
	text, _ := reader.ReadString('\n')
	fmt.Println("Anagrams are: ", findAnagrams(text))
}

func findAnagrams(word string) []string {
	w := sortWord(strings.ToLower(word))

	if v, ok := anagrams[w]; ok {
		return v
	}
	return nil
}

func readDict() map[string][]string {
	file, err := os.Open("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	words := make(map[string][]string)
	scanner := bufio.NewScanner(file)
	s := stats{words: 0, keys: 0}

	for scanner.Scan() {
		s.words++
		w := strings.ToLower(scanner.Text())
		//Sort the word to find the key
		sorted := sortWord(w)
		words[sorted] = append(words[sorted], w)
	}
	fmt.Printf("Stats are %+v\n", s)
	return words
}

func sortWord(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.TrimSpace(strings.Join(s, ""))
}
