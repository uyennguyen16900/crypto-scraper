package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/

// Crypto struct
type Crypto struct {
	Name   string
	Price  string
	Symbol string
}

// GetCoins return a string slice of coins that have price under given price
func GetCoins(price float64, coins []Crypto) []string {
	var res []string
	for _, coin := range coins {
		fmt.Println("price", coin.Price[1:])
		re := regexp.MustCompile(`[0-9.]+`)
		priceArr := re.FindAllString(coin.Price[1:], -1)
		coinPriceStr := strings.Join(priceArr[:], "")
		coinPrice, _ := strconv.ParseFloat(coinPriceStr, 64)

		if coinPrice <= price {
			res = append(res, coin.Symbol)
		}
	}
	return res
}

// GetLimitPrice returns the limit price input
func GetLimitPrice() float64 {
	fmt.Println("Enter the limit price: ")

	var price float64
	fmt.Scanln(&price)
	return price
}

func main() {
	fName := "cryptocoinmarketcap.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	writer.Write([]string{"Name", "Symbol", "Price (USD)"})

	// Instantiate default collector
	c := colly.NewCollector()

	var coins []Crypto
	// var jsonText = []byte(`[
	//     {"Name": "", "Symbol": "", "Price": ""}
	// ]`)

	c.OnHTML("tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			crypto := Crypto{
				Name:   el.ChildText("td:nth-child(2)"),
				Symbol: el.ChildText("td:nth-child(3)"),
				Price:  el.ChildText(".price___3rj7O"),
			}

			writer.Write([]string{
				crypto.Name,
				crypto.Symbol,
				crypto.Price,
			})

			// if err := json.Unmarshal([]byte(jsonText), &coins); err != nil {
			// 	log.Println(err)
			// }
			coins = append(coins, crypto)

		})
		cryptoJSON, _ := json.MarshalIndent(coins, "", "")
		_ = ioutil.WriteFile("output.json", cryptoJSON, 0644)

		fmt.Println(GetCoins(GetLimitPrice(), coins))

	})

	c.Visit("https://coinmarketcap.com/all/views/all/")

	log.Printf("Scraping finished, check file %q for results\n", fName)
}
