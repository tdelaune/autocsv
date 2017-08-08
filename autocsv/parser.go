package autocsv

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)


// New Init feed
func New(file, separator, searchField, neededFieldsRaw string) *Feed {
	splitNeededFields := strings.Split(neededFieldsRaw, ",")
	var neededFields []*Index

	for _, v := range splitNeededFields {
		neededFields = append(neededFields, &Index{Name: v})
	}

	records := make(map[string]*Value)

	return &Feed{
		File:         file,
		Separator:    rune(separator[0]),
		SearchField:  &Index{Name: searchField},
		NeededFields: neededFields,
		records:      records,
	}
}

// Parse CSV file parser
func (f *Feed) Parse() {
	file, err := os.Open(f.File)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)
	reader.Comma = f.Separator

	f.parseHeader(reader)

	for {
		record := parseLine(reader)

		if record == nil {
			break
		}

		f.Store(record)
	}
}

func parseLine(r *csv.Reader) []string {
	record, err := r.Read()

	if err == io.EOF {
		return nil
	} else if err != nil {
		log.Fatal(err)
	}

	return record
}

func (f *Feed) parseHeader(r *csv.Reader) {
	headers := parseLine(r)

	for i, v := range headers {
		if f.SearchField.Name == v {
			f.SearchField.Position = i
		}
	}

	for _, field := range f.NeededFields {
		for i, v := range headers {
			if field.Name == v {
				field.Position = i
			}
		}
	}
}

// Store store csv record in memory
func (f *Feed) Store(record []string) {
	key := record[f.SearchField.Position]

	if _, ok := f.records[key]; ok {
		// return if key already presents in records
		return
	}

	formattedKey := strings.ToLower(removeAccent(key))

	f.records[formattedKey] = f.setFieldsAndOriginalKey(record, key)
}

func (f *Feed) setFieldsAndOriginalKey(record []string, originalKey string) *Value {
	value := &Value{OriginalKey: originalKey}

	for _, index := range f.NeededFields {
		indexValue := &FieldValue{Value: record[index.Position]}
		indexValue.Index = index
		value.FieldValues = append(value.FieldValues, indexValue)
	}

	return value
}

func removeAccent(keyword string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	s, _, _ := transform.String(t, keyword)
	return s
}
