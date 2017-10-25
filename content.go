package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/asciimoo/colly"
)

const contentAuthorSelector = "div#main-content .article-metaline:nth-child(1) .article-meta-value"
const contentTitleSelector = "div#main-content .article-metaline:nth-child(3) .article-meta-value"
const timeSelector = "div#main-content .article-metaline:nth-child(4) .article-meta-value"

const pushSelector = "div#main-content div.push"
const spanF2Selector = "div#main-content span.f2"
const contentLinkSelector = "div#main-content a"
const ipReg = `(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})`
const timeReg = `(\d{1,2}\/\d{1,2}\s\d{1,2}:\d{1,2})`

var count int

func content(e *colly.HTMLElement) {
	log.Println(e.Request.URL)
	count++
	log.Print("Content count", count)
	var urls = []string{}
	content := &Content{}
	content.Athor = e.DOM.Find(contentAuthorSelector).Text()
	content.Title = e.DOM.Find(contentTitleSelector).Text()
	if category := getCategory(content.Title); category != "" {
		content.Category = category
	}
	content.Datetime = e.DOM.Find(timeSelector).Text()

	e.DOM.Find(contentLinkSelector).Map(func(i int, s *goquery.Selection) string {
		if s.Text() != "" {
			urls = append(urls, s.Text())
		}
		return ""
	})
	e.DOM.Find(spanF2Selector).Contents().Map(func(i int, s *goquery.Selection) string {
		// fmt.Println(s.Text())
		if text := s.Text(); text != "" && strings.Contains(text, "發信站") {
			if t, err := getIP(text); err == nil {
				content.PublishIP = t
			} else {
				fmt.Println(err)
			}
		}
		if text := s.Text(); text != "" && strings.Contains(text, "編輯") {
			if t, err := getIP(text); err == nil {
				content.EditedIP = t
			} else {
				fmt.Println(err)
			}
		}

		if link, exist := e.DOM.Find("a").Attr("href"); exist {
			content.Link = link
		}
		return ""
	})
	e.DOM.Find(pushSelector).Map(func(i int, s *goquery.Selection) string {
		reply := &Reply{}

		tag := strings.Trim(s.Children().Filter(".push-tag").Text(), " ")
		reply.Userid = strings.Trim(s.Children().Filter(".push-userid").Text(), " ")
		reply.Content = strings.Trim(strings.Replace(s.Children().Filter(".push-content").Text(), ":", "", 1), " ")
		ipdatetime := strings.Trim(s.Children().Filter(".push-ipdatetime").Text(), " ")
		if ip, err := getIP(ipdatetime); err == nil {
			reply.IP = ip
		}
		if time, err := getReplyTime(ipdatetime); err == nil {
			reply.Time = time
		}
		switch tag {
		case "推":
			content.Push = append(content.Push, reply)
		case "→":
			content.Neutral = append(content.Neutral, reply)
		case "噓":
			content.Boo = append(content.Boo, reply)
		}

		return ""
	})

	content.Content = e.DOM.Children().Remove().End().Text()
	content.Urls = urls
	contents = append(contents, content)
	// fmt.Println(content)
}

func getIP(text string) (string, error) {
	r, err := regexp.Compile(ipReg)
	if err != nil {
		return "", err
	}
	return r.FindString(text), nil
}

func getReplyTime(text string) (string, error) {
	r, err := regexp.Compile(timeReg)
	if err != nil {
		return "", err
	}
	return r.FindString(text), nil
}
