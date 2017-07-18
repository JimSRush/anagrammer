package main

import "fmt"
import "os"
import "bufio"
import "log"
import "strings"
import "sort"

var anagrams map[string][]string

type stats struct {
	words int
}

func main() {
	anagrams = readDict()
	listen()
}

func listen() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a word then enter: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(findAnagrams(text))
}

func findAnagrams(word string) []string {
	fmt.Print("Finding anagram of ", word)
	w := strings.ToLower(word)
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

	s := stats{words: 0}

	for scanner.Scan() {
		s.words++
		w := strings.ToLower(scanner.Text())
		//Sort the word to find the collision
		sorted := sortWord(w)
		if v, ok := words[sorted]; ok {
			//If the key exists, append the word to the slice
			v = append(v, w)
		} else {
			words[sorted] = append(words[sorted], w)
		}
	}

	return words
}

func sortWord(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
