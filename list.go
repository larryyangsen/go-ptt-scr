package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/asciimoo/colly"
)

const titleSelector = ".title a"
const titleLinkSelector = ".title a"
const authorSelector = ".meta .author"
const dateSelector = ".meta .date"
const pushContentSelector = ".nrec"

const categoryPatten = `\[(.+)\]`
const preNumPatten = `index(\d+).html`

var listCount int

func list(e *colly.HTMLElement) {
	dom := e.DOM
	athor := dom.Find(authorSelector).Text()
	title := dom.Find(titleSelector).Text()
	push := dom.Find(pushContentSelector).Text()
	date := dom.Find(dateSelector).Text()
	if category := getCategory(title); category != "" {
		item.Category = category
	}
	if link, exist := dom.Find(titleLinkSelector).Attr("href"); exist {
		item.Link = fmt.Sprintf("https://www.ptt.cc%s", link)
	}

	item.Athor = athor
	item.Title = title
	item.Push = push
	item.Date = date
	output.Items = append(output.Items, item)
}
func prePage(e *colly.HTMLElement) {
	link := e.Attr("href")
	output.PrePage = fmt.Sprintf("%s%s", mainURL, link)
	output.PrePageNumber = getPrePageNum(link)
	log.Println(e.Request.URL)
	listCount++
	log.Print("Page Count", listCount)
	e.Request.Visit(output.PrePage)
}
func getCategory(title string) string {
	if title != "" {
		r, _ := regexp.Compile(categoryPatten)
		return r.FindString(title)
	}
	return ""

}
func getPrePageNum(prePage string) int {
	if prePage != "" {
		r, _ := regexp.Compile(preNumPatten)
		if num, err := strconv.Atoi(r.FindStringSubmatch(prePage)[1]); err == nil {
			return num
		}
	}
	return 0
}
