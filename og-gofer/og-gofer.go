package ogGofer

import (
	"fmt"
	"strings"
	"net/http"
	"golang.org/x/net/html"
)

type ResponseObject struct {
	Success bool `json:"success"`
	Payload map[string]interface{} `json:"payload"`
}

// Takes an HTML token, returning 'ok' == true if the token contains a valid OG
// tag, along with the tag's property and content
func getOgTag(token html.Token) (ok bool, property string, content string) {
	ok = false
	for _, a := range token.Attr {
		switch a.Key {
		case "property":
			if strings.Index(a.Val, "og:") == 0 {
				if content != "" {
					ok = true
				}
				property = a.Val
			}
		case "content":
			if property != "" {
				ok = true
			}
			content = a.Val
		}
	}
	return
}

// Takes a URL (string), and scrapes the page for OG tags. Returns a map of
// OG tag strings, keyed by property
func scrape(url string) (ogTags map[string][]string, err error) {
	ogTags = make(map[string][]string)

	response, err := http.Get(url)

	if err != nil {
		err = HttpError{message: err.Error(), url: url}
		return
	} else if response.StatusCode != http.StatusOK {
		err = HttpError{message: http.StatusText(response.StatusCode), url: url}
		return
	}
	defer response.Body.Close()

	tokenizer := html.NewTokenizer(response.Body)
	for {
		token_type := tokenizer.Next()
		if token_type == html.ErrorToken {
			break
		} else if token_type == html.StartTagToken || token_type == html.SelfClosingTagToken {
			token := tokenizer.Token()
			if token.Data == "meta"  {
				ok, property, content := getOgTag(token)
				if ok {
					ogTags[property] = append(ogTags[property], content)
				}
			}
		}
	}

	return
}

// Scrapes the URL, and builds and returns the response object
func Gofer(url string) (responseObject ResponseObject) {
	responseObject.Payload = make(map[string]interface{})

	if url == "" {
		responseObject.Success = false
		responseObject.Payload["error"] = "No URL provided"
		return
	}

	responseObject.Payload["url"] = url

	ogTags, err := scrape(url)

	if err != nil {
		responseObject.Success = false
		responseObject.Payload["error"] = err.Error()
		return
	}

	responseObject.Success = true
	responseObject.Payload["ogTags"] = ogTags

	tagsFound := 0
	uniqueTagsFound := 0
	for _, contents := range ogTags {
		tagsFound += len(contents)
		uniqueTagsFound += 1
	}
	responseObject.Payload["tagsFound"] = tagsFound
	responseObject.Payload["uniqueTagsFound"] = uniqueTagsFound

	return
}


//  ==========
//  = ERRORS =
//  ==========
type HttpError struct {
	url     string
	message string
}
func (e HttpError) Error() string {
	if e.message != "" {
		return fmt.Sprintf("HTTP Error retrieving URL %s (%s)", e.url, e.message)
	} else {
		return fmt.Sprintf("HTTP Error retrieving URL %s", e.url)
	}
}
