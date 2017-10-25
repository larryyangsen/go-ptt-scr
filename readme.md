# Go Ptt scrape
## go get -u github.com/asciimoo/colly
## go build
---
## helper
```sh
 ./go-scr -h
```
```
Usage of ./go-scr:
  -board string
        A ptt board name (default "beauty")
  -pages int
        How many pages to scrape (default 1)
  -start int
        What's the page start to scrape

```
## output json file

```sh 
./go-scr >output.json -pages=3 -board=gamesale
```