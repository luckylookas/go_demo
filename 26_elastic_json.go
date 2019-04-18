package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v6"
	"github.com/elastic/go-elasticsearch/v6/esapi"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Hit struct {
	Id     string                 `json:"_id"`
	Source map[string]interface{} `json:"_source"`
}

type HitsObject struct {
	Total int   `json:"total"`
	Hits  []Hit `json:"hits"`
}

type Result struct {
	Took float32    `json:"took"`
	Hits HitsObject `json:"hits"`
}


type Person struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	BirthDate string `json:"brithDate"`
	Addresses []Address `json:"addresses"`
}

type Address struct {
	Street string `json:"street"`
	Zip string		`json:"zip"`
}

/*
tuning
https://github.com/elastic/go-elasticsearch/blob/master/_examples/fasthttp/fasthttp.go
 */
func main() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	check(err)


	p := Person{
		LastName: "Bistakova",
		FirstName:"Ludmila",
		BirthDate: "1958-01-01",
		Addresses: []Address{{Street: "Bratislavastreet", Zip: "1000"}},
	}

	buf, _ := json.Marshal(p)

	fmt.Println(string(buf))

	return

	req := esapi.IndexRequest{
		DocumentType: "doc",
		Index:        "person",
		Body:         strings.NewReader(string(buf)),
		Refresh:      "true",
	}



	res, err := req.Do(context.Background(), es)
	check(err)
	res.Body.Close()

	res, err = es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithDocumentType("_doc"),
		es.Search.WithIndex("my_index"),
		es.Search.WithBody(strings.NewReader(`{"query" : { "match" : { "city" : {"query" : "garden", "fuzziness": 2} } }}`)),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	check(err)
	defer res.Body.Close()

	//gjson saves time parsing the json and is much faster if you only need 1 or 2 properties from the json in question
	bytes, err := ioutil.ReadAll(res.Body)
	check(err)
	cityInfo := gjson.GetBytes(bytes, `hits.hits.#._source.city`)

	fmt.Println("city", cityInfo)

	/*
	//you can unmarshal json into structures for full typed access
	var result Result
	err = json.NewDecoder(res.Body).Decode(&result)
	check(err)

	fmt.Println(result.Took)
	fmt.Println(result.Hits.Hits)

	/*
	// you could also unmarshal json without any structural info, just as nested maps using type assertions
	var response  map[string]interface{}

	err = json.NewDecoder(res.Body).Decode(&response)
	check(err)
	for _, hit := range response["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}
*/


}
