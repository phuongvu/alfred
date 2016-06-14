package main

import (
	"flag"
	"os"
	"runtime"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func mapper(input interface{}, output chan interface{}) {

	results := map[string][]Variant{}

	for _, product := range input.(*Target).Products {
		switch product.ProductType {
		case "Computer":
			for _, v := range product.Variants {
				v.Title = product.Title + " (" + v.Title + ")"
				results["Computer"] = append(results["Computer"], v)
			}
		case "Keyboard":
			for _, v := range product.Variants {
				v.Title = product.Title + " (" + v.Title + ")"
				results["Keyboard"] = append(results["Keyboard"], v)
			}
		}
	}

	output <- results
	close(output)
}

func reducer(input chan interface{}, output chan interface{}) {

	results := map[string][]Variant{}

	for item := range input {
		for key, value := range item.(map[string][]Variant) {
			_, exists := results[key]

			if !exists {
				results[key] = value
			} else {
				results[key] = append(results[key], value...)
			}
		}
	}

	//sort price in ascending order
	price := func(p1, p2 *Variant) bool {
		price1, _ := strconv.ParseFloat(p1.Price, 64)
		price2, _ := strconv.ParseFloat(p2.Price, 64)
		return price1 < price2
	}

	for key, _ := range results {
		//sort by price
		By(price).Sort(results[key])
	}

	output <- results
	close(output)
}

func addToCart(input interface{}) interface{} {
	var cart Cart
	var minLength int
	var totalWeight int
	var totalPrice float64
	var keyboards, computers []Variant

	//Casting input to map[string][]Variant
	keyboards = input.(map[string][]Variant)["Keyboard"]
	computers = input.(map[string][]Variant)["Computer"]

	//Making sure that the number of keyboards is equal to the number of computers
	if len(keyboards) <= len(computers) {
		minLength = len(keyboards)
	} else {
		minLength = len(computers)
	}

	for i := 0; i < minLength; i++ {
		cart.items = append(cart.items, keyboards[i])
		cart.items = append(cart.items, computers[i])
	}

	for _, item := range cart.items {
		price, _ := strconv.ParseFloat(item.Price, 64)
		totalPrice += price
		totalWeight += item.Grams
	}

	cart.TotalPrice = strconv.FormatFloat(totalPrice, 'f', 2, 64)
	cart.TotalWeight = strconv.Itoa(totalWeight)

	return cart
}

func showCart(input interface{}) {
	var data [][]string
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Price", "Weight (in grams)"})
	table.SetFooter([]string{"", "Total weight in grams", input.(Cart).TotalWeight})

	for _, value := range input.(Cart).items {
		weight := strconv.Itoa(value.Grams)
		data = append(data, []string{value.Title, value.Price, weight})
	}

	table.AppendBulk(data)
	table.Render()
}

const (
	MaxWorkers      = 10
	defaultUrl      = "shopicruit.myshopify.com/products.json"
	defaultPages    = 5
	defaultProtocol = "https"
)

func main() {
	//Utilize every single core
	runtime.GOMAXPROCS(runtime.NumCPU())

	var c CrawlerImpl

	//Setup command-line flag
	shopURL := flag.String("url", defaultUrl, "Shop's url")
	numberOfPages := flag.Int("pages", defaultPages, "Number of pages to craw")
	preferredProtocol := flag.String("protocol", defaultProtocol, "Communication protocol")

	flag.Parse()

	//Params: base URL, number of page, protocol, go type def
	input := Craw(c, *shopURL, *numberOfPages, *preferredProtocol)

	//map: filter products by product type and reduce
	output := mapReduce(mapper, reducer, input)

	//produce cart
	cart := addToCart(<-output)

	showCart(cart)
}
