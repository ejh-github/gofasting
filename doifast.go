package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
	"golang.org/x/net/html"
)

func main() {
	response, err := http.Get("https://goarch.org")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	textTags := []string{
		"span", "h3",
	}
	tag := ""
	enter := false
	tokenizer := html.NewTokenizer(response.Body)
	for {
		tt := tokenizer.Next()
		token := tokenizer.Token()
		err := tokenizer.Err()
		if err == io.EOF {
			break
		}
		switch tt {
		case html.ErrorToken:
			log.Fatal(err)
		case html.StartTagToken, html.SelfClosingTagToken:
			enter = false
			tag = token.Data
			for _, ttt := range textTags {
				if tag == ttt {
					enter = true
					break
				}
			}
		case html.TextToken:
			if enter {
				data := strings.TrimSpace(token.Data)
				if len(data) > 0 {
					temp := strings.Split(data, "\n")
					for i := 0; i < len(temp); i++ {
						if strings.Contains(temp[i], "Fast") == true {
							fmt.Println(time.Now())
							fmt.Println(temp[i])
						}
					}
				}
			}
		}
	}
}
