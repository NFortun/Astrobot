package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	res, err := http.DefaultClient.Get("https://www.lameteoagricole.net/meteo-heure-par-heure/Cesson-Sevigne-35510-j1.html")
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	node, err := html.Parse(res.Body)
	if err != nil {
		panic(err)
	}

	table := getTable(node)
	header := getHeader(table)
	//fmt.Printf("header: %s", header.Data)
	fmt.Println(extractHeaderInformations(header))
}

func getTable(n *html.Node) *html.Node {
	if n.Type == html.ElementNode && n.Data == "table" {
		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == "heures-table" {
				return n
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if result := getTable(c); result != nil {
			return result
		}
	}

	return nil
}

func extractHeaderInformations(n *html.Node) string {
	var result string
	if n.Type == html.TextNode && n.Data != "" {
		return strings.TrimSpace(n.Data)
	}
	if n.Type == html.ElementNode && n.Data == "th" {
		for _, attr := range n.Attr {
			if attr.Key == "data-field" {

			}
		}
		result = fmt.Sprintf("%s\n", result)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result = fmt.Sprintf("%s %s", result, extractHeaderInformations(c))
	}

	return result
}

func getHeader(n *html.Node) *html.Node {
	if n.Data == "thead" {
		return n
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if result := getHeader(c); result != nil {
			return result
		}
	}

	return nil
}
