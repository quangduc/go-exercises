package link

// package main

// import (
// 	"bytes"
// 	"flag"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"strings"

// 	"golang.org/x/net/html"
// )

// func main() {
// 	url := flag.String("url", "https://www.prnewswire.com/news-releases/news-releases-list/", "url to parse html")
// 	flag.Parse()
// 	if *url == "" {
// 		panic("url cannot be empty")
// 	}
// 	resp, err := http.Get(*url)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	htmlBody := string(body)

// 	doc, err := html.Parse(strings.NewReader(string(htmlBody)))
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	var result []Link

// 	var collectText func(n *html.Node, buf *bytes.Buffer)
// 	collectText = func(n *html.Node, buf *bytes.Buffer) {
// 		if n.Type == html.TextNode {
// 			buf.WriteString(n.Data)
// 		}
// 		for c := n.FirstChild; c != nil; c = c.NextSibling {
// 			collectText(c, buf)
// 		}
// 	}

// 	var traverLinkNode func(*html.Node)
// 	traverLinkNode = func(n *html.Node) {
// 		if n.Type == html.ElementNode && n.Data == "a" {
// 			textBuf := &bytes.Buffer{}
// 			collectText(n, textBuf)
// 			var href string
// 			for _, attr := range n.Attr {
// 				if attr.Key == "href" {
// 					href = attr.Val
// 					break
// 				}
// 			}
// 			result = append(result, Link{href, textBuf.String()})
// 		}
// 		for c := n.FirstChild; c != nil; c = c.NextSibling {
// 			traverLinkNode(c)
// 		}
// 	}

// 	traverLinkNode(doc)
// 	fmt.Printf("%v", result)
// }
