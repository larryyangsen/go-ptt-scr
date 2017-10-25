package main

type Output struct {
	PrePage       string
	PrePageNumber int
	Items         []Item
}

type Item struct {
	Title    string `json:"title"`
	Category string `json:"category"`
	Link     string `json:"link"`
	Athor    string `json:"athor"`
	Push     string `json:"push"`
	Date     string `json:"date"`
}
type Content struct {
	Athor     string   `json:"athor"`
	Title     string   `json:"title"`
	Category  string   `json:"category"`
	Link      string   `json:"link"`
	Datetime  string   `json:"datetime"`
	Urls      []string `json:"urls"`
	Content   string   `json:"content"`
	PublishIP string   `json:"publishIP"`
	EditedIP  string   `json:"editedIP"`
	Push      []*Reply `json:"push"`
	Boo       []*Reply `json:"boo"`
	Neutral   []*Reply `json:"neutral"`
}

type Reply struct {
	Userid  string `json:"userid"`
	Content string `json:"content"`
	IP      string `json:"ip"`
	Time    string `json:"time"`
}
