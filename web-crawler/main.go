package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"golang.org/x/net/html"
)

func TraverseNodeTree(*html.Node) error {
	return nil
}

func main() {
	links := os.Args[1]

	response, err := http.Get(links)

	if err != nil {
		fmt.Printf("[ERROR] Cannot fetch url: %s, Message: %s\n", links, err)
		os.Exit(1)
	}

	defer response.Body.Close()

	rootNode, err := html.Parse(response.Body)

	if err != nil {
		fmt.Printf("[ERROR] Html cannot be parsed. Message: %s", err)
		os.Exit(2)
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {

			for _, attr := range n.Attr {
				if attr.Key == "href" {
					// fmt.Printf("[DATA] Found: %s \n", attr.Val)
					data, err := url.Parse(attr.Val)

					if err != nil {
						fmt.Printf("[ERROR] URL cannot be parsed. Message: %s", err)
						os.Exit(3)
					}

					if data.Host != "" {
						fmt.Printf("[DATA] Found data domain: %s\n", data.Host)
						fmt.Printf("------ Found data path: %s\n\n", data.Path)
					} else {
						fmt.Println("[DATA] No domain")
						fmt.Printf("------ Found data path: %s\n\n", data.Path)
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(rootNode)
}
