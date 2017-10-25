package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/asciimoo/colly"
)

const mainURL = "https://www.ptt.cc"
const over18URL = "https://www.ptt.cc/ask/over18"

const prePageSelector = ".btn-group-paging a:nth-child(2)"
const listSelector = ".r-ent"
const mainContentSelector = "div#main-content"

var (
	listColllector   *colly.Collector
	contentCollector *colly.Collector
	item             Item
	output           Output
	contents         []*Content
	boardName        = flag.String("board", "beauty", "A ptt board name")
	start            = flag.Int("start", 0, "What's the page start to scrape")
	pages            = flag.Int("pages", 1, "How many pages to scrape")
)

func init() {
	listColllector = colly.NewCollector()
	contentCollector = colly.NewCollector()
	storeSession(listColllector)
	storeSession(contentCollector)
	flag.Parse()

	output = Output{}
	item = Item{}

}

func main() {
	boardURL := ""
	if *start == 0 {
		boardURL = fmt.Sprintf("%s/bbs/%s/index.html", mainURL, *boardName)

	} else {
		boardURL = fmt.Sprintf("%s/bbs/%s/index%d.html", mainURL, *boardName, *start)

	}
	listColllector.MaxDepth = *pages

	listColllector.OnHTML(listSelector, list)
	listColllector.OnHTML(prePageSelector, prePage)
	listColllector.Visit(boardURL)
	contentCollector.OnHTML(mainContentSelector, content)
	print()
}
func storeSession(ctrl *colly.Collector) {
	form := map[string]string{}
	form["from"] = "/bbs/Gossiping/index.html"
	form["yes"] = "yes"
	err := ctrl.Post(over18URL, form)
	if err != nil {
		panic(err)
	}

}

func print() {
	listJSON, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("{\n\"list\":")
	fmt.Println(string(listJSON), ",")

	for _, item := range output.Items {
		contentCollector.Visit(item.Link)
	}
	// for _, content := range contents {
	// 	fmt.Println(content)
	// }
	contentJSON, err := json.MarshalIndent(contents, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\"contens\":")
	fmt.Println(string(contentJSON))
	fmt.Print("}")

}
