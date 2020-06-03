package okko_fims_parser

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func WriteCsv(films []Film, collection OkkoRequestElementAlias) {
	currentTime := time.Now()

	file, err := os.Create(fmt.Sprintf("films_%s_%s.csv", collection, currentTime.Format("02_01_2006T15_04_05")))
	if err != nil {
		log.Fatal("failed create file", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writeRow(writer, []string{
		"ID",
		"Item title",
		"Final URL",
		"Image URL",
		"Item subtitle",
		"Item description",
		"Sale price",
		"Price",
	})

	for _, film := range films {
		urlWords := strings.Split(film.PicCoverUrl, "/")
		writeRow(writer, []string{
			urlWords[len(urlWords)-1],
			film.Name,
			film.Url,
			film.PicCoverUrl,
			film.Name,
			"Купить фильм за 1₽",
			"",
			"1 RUB",
		})
	}
}

func writeRow(writer *csv.Writer, row []string) {
	err := writer.Write(row)
	if err != nil {
		log.Fatal("failed to write to file", err)
	}
}
