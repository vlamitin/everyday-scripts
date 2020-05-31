package okko_fims_parser

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type jsonResponse struct {
	Element jsonResponseElement `json:"element"`
}

type jsonResponseElement struct {
	CollectionItems jsonResponseCollectionItems `json:"collectionItems"`
}

type jsonResponseCollectionItems struct {
	TotalSize int                          `json:"totalSize"`
	Items     []jsonResponseCollectionItem `json:"items"`
}

type jsonResponseCollectionItem struct {
	Element jsonResponseFilm `json:"element"`
}

type jsonResponseFilmCovers struct {
	Items []jsonResponseFilmCoverItem `json:"items"`
}

type coverImageType string

const (
	cover    coverImageType = "COVER"
	portrait coverImageType = "PORTRAIT"
)

type jsonResponseFilmCoverItem struct {
	Url       string         `json:"url"`
	ImageType coverImageType `json:"imageType"`
}

type jsonResponseFilm struct {
	Name         string                 `json:"name"`
	OriginalName string                 `json:"originalName"`
	Type         string                 `json:"type"`
	Alias        string                 `json:"alias"`
	BasicCovers  jsonResponseFilmCovers `json:"basicCovers"`
}

type Film struct {
	Name           string
	Url            string
	PicCoverUrl    string
	PicPortraitUrl string
}

type OkkoCollection string

const (
	OptimumCollection  OkkoCollection = "msvod_all_optimum"
	NewPromoCollection OkkoCollection = "new-promo"
)

func GetFilms(collection OkkoCollection) []Film {
	fmt.Println("collection", collection)
	firstResponse := getResponse(collection, 10)
	fmt.Printf("firstResponse films found: %d\n", len(firstResponse.Element.CollectionItems.Items))
	fmt.Println("firstResponse.Element.CollectionItems.TotalSize", firstResponse.Element.CollectionItems.TotalSize)
	secondResponse := getResponse(collection, firstResponse.Element.CollectionItems.TotalSize)

	fmt.Printf("secondResponse films found: %d\n", len(secondResponse.Element.CollectionItems.Items))

	films := []Film{}

	for _, parsedItem := range secondResponse.Element.CollectionItems.Items {
		coverItemIndex := searchIndex(parsedItem.Element.BasicCovers.Items, func(item jsonResponseFilmCoverItem) bool {
			return item.ImageType == cover
		})
		portraitItemIndex := searchIndex(parsedItem.Element.BasicCovers.Items, func(item jsonResponseFilmCoverItem) bool {
			return item.ImageType == portrait
		})

		var picCoverUrl string = ""
		var picPortraitUrl string = ""

		if coverItemIndex != -1 {
			picCoverUrl = parsedItem.Element.BasicCovers.Items[coverItemIndex].Url
		}
		if portraitItemIndex != -1 {
			picPortraitUrl = parsedItem.Element.BasicCovers.Items[portraitItemIndex].Url
		}

		films = append(films, Film{
			Name:           parsedItem.Element.Name,
			Url:            strings.ToLower(fmt.Sprintf("https://okko.tv/%s/%s", parsedItem.Element.Type, parsedItem.Element.Alias)),
			PicCoverUrl:    picCoverUrl,
			PicPortraitUrl: picPortraitUrl,
		})
	}

	return films
}

func getResponse(collection OkkoCollection, limit int) jsonResponse {
	client := &http.Client{}

	// https://ctx.playfamily.ru/screenapi/v1/noauth/collection/web/1?elementAlias=msvod_all_optimum&elementType=COLLECTION&limit=16&offset=0&withInnerCollections=true
	req, err := http.NewRequest(
		"GET",
		"https://ctx.playfamily.ru/screenapi/v1/noauth/collection/web/1"+"?"+url.Values{
			"elementAlias":         []string{string(collection)},
			"elementType":          []string{"COLLECTION"},
			"limit":                []string{strconv.Itoa(limit)},
			"offset":               []string{"0"},
			"withInnerCollections": []string{"true"},
		}.Encode(),
		nil)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var result jsonResponse

	json.NewDecoder(resp.Body).Decode(&result)

	return result
}

func searchIndex(items []jsonResponseFilmCoverItem, predicate func(item jsonResponseFilmCoverItem) bool) int {
	for i, item := range items {
		if predicate(item) {
			return i
		}
	}

	return -1
}
