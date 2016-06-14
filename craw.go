package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type Crawler interface {
	PreppingUrls(string, int, string) chan string
	GetJson(string, interface{}) error
	FetchDataFromURL(string) interface{}
}

type CrawlerImpl struct{}

func (c CrawlerImpl) PreppingUrls(baseUrl string, page int, protocol string) chan string {

	output := make(chan string)
	u, err := url.Parse(baseUrl)

	if err != nil {
		log.Printf("Error fetching: %v", err)
	}

	//Prefer https over http
	u.Scheme = protocol

	go func() {
		for i := 1; i < page+1; i++ {
			q := u.Query()
			q.Set("page", strconv.Itoa(i))
			u.RawQuery = q.Encode()
			rawUrl := u.String()
			output <- rawUrl
		}
		close(output)
	}()

	return output
}

func (c CrawlerImpl) GetJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func (c CrawlerImpl) FetchDataFromURL(url string) interface{} {
	log.Printf("URL: %s", url)

	output := make(chan interface{})

	go func() {
		model := &Target{}
		err := c.GetJson(url, model)
		if err != nil {
			log.Fatal("Decode failed!", err)
		}
		output <- model
		close(output)
	}()

	return <-output
}

func Craw(c Crawler, shopUrl string, num int, protocol string) chan interface{} {
	output := make(chan interface{}, MaxWorkers)

	go func() {
		for url := range c.PreppingUrls(shopUrl, num, protocol) {
			output <- c.FetchDataFromURL(url)
		}
		close(output)
	}()

	return output
}
