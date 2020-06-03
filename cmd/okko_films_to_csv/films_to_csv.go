package main

import (
	"fmt"
	"github.com/vlamitin/everyday-scripts/internal/okko_fims_parser"
)

var (
	ElementAlias string
	ElementType  string
	BuildDate    string
)

func main() {
	fmt.Printf(
		"Executing okko_films_to_csv script, build date: %s, elementAlias: %s, elementType: %s\n",
		BuildDate,
		ElementAlias,
		ElementType,
	)
	films := okko_fims_parser.GetFilms(
		okko_fims_parser.OkkoRequestElementAlias(ElementAlias),
		okko_fims_parser.OkkoRequestElementType(ElementType),
	)
	okko_fims_parser.WriteCsv(films, okko_fims_parser.OkkoRequestElementAlias(ElementAlias))
}
