package carouscrapy

import (
	"fmt"
	"log"
	"strings"

	"github.com/antchfx/htmlquery"
)

// Item An Carousell item
type Item struct {
	Username    string
	Name        string
	Price       string
	Category    string
	Description string
	URL         string
	PublishedAt string
	Sold        bool
}

// LoadByURL loads item page and return info.
func LoadByURL(url string) {
	doc, err := htmlquery.LoadURL(url)

	if err != nil {
		log.Fatal(err)
	}

	item := &Item{URL: url}

	userNode := htmlquery.FindOne(doc, "/html/body/div[1]/div/div[3]/div[1]/div/div[2]/div[1]/div[1]/div[2]/a")
	item.Username = htmlquery.InnerText(userNode)

	priceNode := htmlquery.FindOne(doc, "/html/body/div[1]/div/div[3]/div[1]/div/div[2]/div[2]/h2")
	item.Price = htmlquery.InnerText(priceNode)

	titleNode := htmlquery.FindOne(doc, "/html/body/div[1]/div/div[3]/div[1]/div/div[2]/div[4]/h1")
	item.Name = htmlquery.InnerText(titleNode)

	categoryNode := htmlquery.FindOne(doc, "/html/body/div[1]/div/div[3]/div[1]/div/div[2]/section/div[1]/div[6]/div/div/div/a")
	item.Category = strings.TrimSpace(htmlquery.InnerText(categoryNode))
	// categoryUrl := htmlquery.SelectAttr(categoryNode, "href")

	descNode := htmlquery.FindOne(doc, "/html/body/div[1]/div/div[3]/div[1]/div/div[2]/section/div[1]/div[5]/div")
	item.Description = strings.TrimSpace(htmlquery.InnerText(descNode))

	item.PrintInfo()

	// fmt.Printf(">>> [%s] %s (%s)\n%s\n", category, title, price, desc)

	// for _, attr := range categoryNode.Attr {
	// fmt.Println("  >", attr.Key, "=", attr.Val)
	// }

	// fmt.Println(">>>>",
	// 	categoryNode,
	// 	htmlquery.SelectAttr(categoryNode, "href"),
	// 	htmlquery.SelectAttr(categoryNode, "outterText"),
	// 	htmlquery.SelectAttr(categoryNode, "hostname"),
	// )
}

// PrintInfo display summary of Item
func (i Item) PrintInfo() {
	fmt.Printf("[%s] %s (%s)\n%s\n%s\n%s\n", i.Category, i.Name, i.Price, i.Description, i.Username, i.URL)
}

// LoadByID loads item by ID
func LoadByID(id string) {
	url := fmt.Sprintf("https://sg.carousell.com/p/%s", id)
	LoadByURL(url)
}
