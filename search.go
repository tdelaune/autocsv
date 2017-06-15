package main

import (
	"regexp"
	"sort"
	"strings"
)

// Result from search
type Result struct {
	Key    string            `json:"key"`
	Fields map[string]string `json:"fields"`
}

// Results slice of Result
type Results []*Result

// Search find keyword in feed records
func (f *Feed) Search(keyword string) Results {
	var results Results
	keyword = strings.ToLower(removeAccent(keyword))

	for key, value := range f.records {
		if hasKeyword(key, keyword) {
			result := &Result{Key: value.OriginalKey, Fields: setFields(value)}
			results = append(results, result)
		}
	}

	sort.Sort(results)

	return results
}

// SearchRegex only for benchmarking
func (f *Feed) SearchRegex(keyword string) Results {
	var results Results
	keyword = strings.ToLower(removeAccent(keyword))
	regexMatcher := "^" + keyword + "|\\W" + keyword

	var regex = regexp.MustCompile(regexMatcher)

	for key, value := range f.records {
		if regex.MatchString(key) {
			result := &Result{Key: value.OriginalKey, Fields: setFields(value)}
			results = append(results, result)
		}
	}

	sort.Sort(results)

	return results
}

func setFields(value *Value) map[string]string {
	fields := make(map[string]string)

	for _, v := range value.FieldValues {
		fields[v.Name] = v.Value
	}

	return fields
}

func hasKeyword(key, keyword string) bool {
	if strings.HasPrefix(key, keyword) {
		return true
	}

	if strings.Contains(key, " "+keyword) {
		return true
	}

	if strings.Contains(key, "'"+keyword) {
		return true
	}

	return false
}

func (s Results) Len() int {
	return len(s)

}

func (s Results) Less(i, j int) bool {
	return s[i].Key < s[j].Key

}

func (s Results) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]

}
