package autocsv

// Feed contains document and search criterias
type Feed struct {
	File         string
	Separator    rune
	SearchField  *Index
	NeededFields []*Index
	records      map[string]*Value
}

// Value represents non formatted key and fields
type Value struct {
	OriginalKey string
	FieldValues []*FieldValue
}

// Index represents indexed keys
type Index struct {
	Name     string
	Position int
}

// FieldValue for specified Index
type FieldValue struct {
	*Index
	Value string
}
