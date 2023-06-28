package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"golang.org/x/net/html"
)

type URLStructure struct {
	scheme string `default:"https"`
	host   string
	path   string
}

type FetchGroup struct {
	addr string
	root *html.Node
}

func Fetch(addr string) (*FetchGroup, error) {
	response, err := http.Get(addr)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	n, err := html.Parse(response.Body)

	if err != nil {
		return nil, err
	}

	return &FetchGroup{addr: addr, root: n}, nil
}

func TraverseNodeTree(fg *FetchGroup, c chan string, depth uint8) {
	n := fg.root
	// fmt.Printf("Depth!! %s\n", strconv.Itoa(int(depth)))
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				// fmt.Printf("[DATA] Found: %s \n", attr.Val)
				data, err := url.Parse(attr.Val)

				if err != nil {
					// fmt.Printf("[ERROR] URL cannot be parsed. Message: %s", err)
					continue
				}

				foundUrl := &URLStructure{
					scheme: data.Scheme,
					host:   data.Host,
					path:   data.Path,
				}
				if data.Host != "" {
					fmt.Printf("Depth: %s, %s://%s%s\n", strconv.Itoa(int(depth)), foundUrl.scheme, foundUrl.host, foundUrl.path)
					c <- fmt.Sprintf("%s://%s%s", foundUrl.scheme, foundUrl.host, foundUrl.path)
				}
				// fmt.Printf("[CHECKPOINT] 6 -- %s -- %s -- %s \n", foundUrl.scheme, foundUrl.host, foundUrl.path)
				// if data.Host != "" {
				// 	fmt.Printf("%s://%s%s\n\n", foundUrl.scheme, foundUrl.host, foundUrl.path)
				// 	c <- fmt.Sprintf("%s://%s%s", foundUrl.scheme, foundUrl.host, foundUrl.path)
				// }
			}
		}
	}

	if depth == 0 {
		return
	}

	for currChild := fg.root.FirstChild; currChild != nil; currChild = currChild.NextSibling {
		go TraverseNodeTree(&FetchGroup{addr: fg.addr, root: currChild}, c, depth-1)
	}

	return
}

func main() {
	visited := make(map[string]bool)
	queue := make(chan string)
	links := os.Args[1:]

	go func() {
		for _, link := range links {
			queue <- link
		}
	}()

	for newLink := range queue {
		if _, ok := visited[newLink]; ok {
			continue
		} else {
			visited[newLink] = true
		}

		fmt.Printf("[LOG] Working on a new link: %s\n", newLink)

		rootNode, err := Fetch(newLink)

		if err != nil {
			fmt.Printf("[ERROR] Cannot fetch url: %s, Message: %s\n", newLink, err)
			continue
		}

		go TraverseNodeTree(rootNode, queue, 3)
	}
}
