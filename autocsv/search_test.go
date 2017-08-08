package autocsv

import (
	"testing"
)

var feed *Feed

func init() {
	feed = New("liste.csv", ";", "stop_name", "stop_lat,stop_lon")
	feed.Parse()
}

func benchmarkSearch(key string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		feed.Search(key)
	}
}

func benchmarkSearchRegex(key string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		feed.SearchRegex(key)
	}
}

func BenchmarkSearchA(b *testing.B)           { benchmarkSearch("a", b) }
func BenchmarkSearchMai(b *testing.B)         { benchmarkSearch("mai", b) }
func BenchmarkSearchMairie(b *testing.B)      { benchmarkSearch("mairie", b) }
func BenchmarkSearchMairieUpper(b *testing.B) { benchmarkSearch("Mairie", b) }
func BenchmarkSearchEgalite(b *testing.B)     { benchmarkSearch("egalite", b) }
func BenchmarkSearchNotFound(b *testing.B)    { benchmarkSearch("zeghzoieg", b) }

func BenchmarkSearchRegexA(b *testing.B)           { benchmarkSearchRegex("a", b) }
func BenchmarkSearchRegexMai(b *testing.B)         { benchmarkSearchRegex("mai", b) }
func BenchmarkSearchRegexMairie(b *testing.B)      { benchmarkSearchRegex("mairie", b) }
func BenchmarkSearchRegexMairieUpper(b *testing.B) { benchmarkSearchRegex("Mairie", b) }
func BenchmarkSearchRegexEgalite(b *testing.B)     { benchmarkSearchRegex("egalite", b) }
func BenchmarkSearchRegexNotFound(b *testing.B)    { benchmarkSearchRegex("zeghzoieg", b) }
