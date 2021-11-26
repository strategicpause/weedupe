package mapper

import (
	"regexp"
	"strings"
	"weedupe/common"
)

type Mapper interface {
	Map(string) []*common.Pair
}

type WordCountMapper struct {
	splitRe *regexp.Regexp
	replRe *regexp.Regexp
}

func NewWordCountMapper() Mapper {
	return &WordCountMapper {
		splitRe: regexp.MustCompile("\\s+"),
		replRe: regexp.MustCompile("\\W+"),
	}
}

func (m *WordCountMapper) Map(content string) []*common.Pair {
	words := m.splitRe.Split(content, -1)
	pairs := []*common.Pair{}
	for _, word := range words {
		word = m.replRe.ReplaceAllString(word, "")
		word = strings.TrimSpace(word)
		if word != "" {
			pairs = append(pairs, &common.Pair{word, 1})
		}
	}
	return pairs
}