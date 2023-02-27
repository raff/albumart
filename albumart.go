package albumart

import (
	"fmt"
	"github.com/buger/jsonparser"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func GetArtwork(ArtistName string, AlbumSongName string) string {
	query := AlbumSongName
	if ArtistName != "" {
		query = ArtistName + " " + query
	}

	uri := GetArtworkFlags(query, "", true)
        if uri == "" {
            return "Not Found"
        }

        return uri
}

func GetArtworkFlags(query, entity string, large bool) string {
	if entity == "" {
		entity = "song"
	}

	artworkUrlToReturn := ""
	whitespaceRegEx, _ := regexp.Compile(`\s`)
	query = whitespaceRegEx.ReplaceAllString(query, "+")
	response, err := http.Get("https://itunes.apple.com/search?term=" + query + "&limit=1&entity=" + entity)
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		defer response.Body.Close()
		data, readErr := ioutil.ReadAll(response.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}
		resultCount, _ := jsonparser.GetInt(data, "resultCount")
		if resultCount > 0 {
			artworkUrlToReturn, _ = jsonparser.GetString(data, "results", "[0]", "artworkUrl100")
			if large {
				makeLarge, _ := regexp.Compile("100x100bb")
				artworkUrlToReturn = makeLarge.ReplaceAllString(artworkUrlToReturn, "1200x1200bb")
			}
		}
	}
	return artworkUrlToReturn
}
