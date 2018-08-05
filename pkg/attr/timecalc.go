package attr

import (
	"log"
	"strings"
)

const wpm = 250

type Text struct {
	Value string `json:"text"`
}

func (t *Text) TimeToRead() int {
	words := strings.Split(strings.Join(strings.Fields(t.Value), string(' ')), string(' '))
	for _, word := range words {
		log.Println(len(word))
	}
	time := (len(words) * 60) / wpm
	return time
}
