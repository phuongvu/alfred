package main

import (
	"github.com/jarcoal/httpmock"
	"reflect"
	"testing"
)

var (
	baseUrl      = "test.com/products.json"
	expectedUrls = []string{"https://test.com/products.json?page=1", "https://test.com/products.json?page=2"}
	content      = `{"products": [
					{"id":2759196675,"title":"Computer1","vendor":"Hand, Olson","product_type":"Computer","tags":[],"variants":[{"id":8041889923,"title":"Komputer 2","option1":"Lavender","option2":null,"option3":null,"sku":"","requires_shipping":true,"taxable":true,"featured_image":null,"available":true,"price":"38.94","grams":4115,"compare_at_price":null,"position":1,"product_id":2759196675,"created_at":"2015-09-23T20:52:12-04:00","updated_at":"2015-09-23T20:52:12-04:00"}],"images":[{"id":5642059587,"created_at":"2015-09-23T20:52:12-04:00","position":1,"updated_at":"2015-09-23T20:52:12-04:00","product_id":2759196675,"variant_ids":[],"src":""}],"options":[{"name":"Title","position":1,"values":["Lavender","Yellow"]}]},
		    			{"id":2759194115,"title":"Keyboard2","vendor":"Schuppe Group","product_type":"Keyboard","tags":[],"variants":[{"id":8041885891,"title":"Keyboard 2","option1":"Sky blue","option2":null,"option3":null,"sku":"","requires_shipping":true,"taxable":true,"featured_image":null,"available":true,"price":"1.05","grams":5493,"compare_at_price":null,"position":1,"product_id":2759194115,"created_at":"2015-09-23T20:52:09-04:00","updated_at":"2015-09-23T20:52:09-04:00"}],"images":[{"id":5642059421,"created_at":"2015-09-23T20:52:12-04:00","position":1,"updated_at":"2015-09-23T20:52:12-04:00","product_id":2759196675,"variant_ids":[],"src":""}],"options":[{"name":"Title","position":1,"values":["Lavender","Yellow"]}]}
		    		]
		   	}`

	expectedObj = &Target{
		[]Product{
			Product{
				"Computer1",
				"Computer",
				[]Variant{
					Variant{
						Title: "Komputer 2",
						Price: "38.94",
						Grams: 4115,
					},
				},
			},
			Product{
				"Keyboard2",
				"Keyboard",
				[]Variant{
					Variant{
						Title: "Keyboard 2",
						Price: "1.05",
						Grams: 5493,
					},
				},
			},
		},
	}
)

type FakeCrawlerImpl struct{}

func TestCrawlerImpl_PreppingUrls(t *testing.T) {

	var c CrawlerImpl
	var urls []string

	results := c.PreppingUrls(baseUrl, 2, "https")

	for url := range results {
		urls = append(urls, url)
	}

	if !reflect.DeepEqual(urls, expectedUrls) {
		t.Errorf("Expected %v", expectedUrls)
		t.Errorf("     got %v", urls)
	}
}

func TestCraw(t *testing.T) {
	var c CrawlerImpl

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://test.com/products.json?page=1",
		httpmock.NewStringResponder(200, content))

	result := Craw(c, "test.com/products.json", 1, "https")

	for c := range result {
		if !reflect.DeepEqual(c, expectedObj) {
			t.Errorf("Expected %v", expectedObj)
			t.Errorf("     got %v", c)
		}
	}

}
