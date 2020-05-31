package main

import (
	"fmt"
	"github.com/vlamitin/everyday-scripts/internal/okko_fims_parser"
)

var (
	Collection string
	BuildDate  string
)

func main() {
	fmt.Printf("Executing okko_films_to_csv script, build date: %s, collection: %s\n", BuildDate, Collection)
	films := okko_fims_parser.GetFilms(okko_fims_parser.OkkoCollection(Collection))
	okko_fims_parser.WriteCsv(films, okko_fims_parser.OkkoCollection(Collection))
}
