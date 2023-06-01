package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
)

func main() {
	log.SetFlags(0)

	var (
		r map[string]interface{}
		// wg sync.WaitGroup
	)

	// Initialize a client with the default settings.
	//
	// An `ELASTICSEARCH_URL` environment variable will be used when exported.
	//
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://127.0.0.1:9200",
		},
		Username: "elasticsearch",
		Password: "*********",
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// 1. Get cluster info
	//
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	// Check response status
	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}
	// Deserialize the response into a map.
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print client and server version numbers.
	log.Printf("Client: %s", elasticsearch.Version)
	log.Printf("Server: %s", r["version"].(map[string]interface{})["number"])
	log.Println(strings.Repeat("~", 37))

	// 2. Index documents concurrently
	//
	// for i, title := range []string{"Test One", "Test Two"} {
	// 	wg.Add(1)
	//
	// 	go func(i int, title string) {
	// 		defer wg.Done()
	//
	// 		// Build the request body.
	// 		data, err := json.Marshal(struct{ Title string }{Title: title})
	// 		if err != nil {
	// 			log.Fatalf("Error marshaling document: %s", err)
	// 		}
	//
	// 		// Set up the request object.
	// 		req := esapi.IndexRequest{
	// 			Index:      "test",
	// 			DocumentID: strconv.Itoa(i + 1),
	// 			Body:       bytes.NewReader(data),
	// 			Refresh:    "true",
	// 		}
	//
	// 		// Perform the request with the client.
	// 		res, err := req.Do(context.Background(), es)
	// 		if err != nil {
	// 			log.Fatalf("Error getting response: %s", err)
	// 		}
	// 		defer res.Body.Close()
	//
	// 		if res.IsError() {
	// 			log.Printf("[%s] Error indexing document ID=%d", res.Status(), i+1)
	// 		} else {
	// 			// Deserialize the response into a map.
	// 			var r map[string]interface{}
	// 			if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
	// 				log.Printf("Error parsing the response body: %s", err)
	// 			} else {
	// 				// Print the response status and indexed document version.
	// 				log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
	// 			}
	// 		}
	// 	}(i, title)
	// }
	// wg.Wait()

	log.Println(strings.Repeat("-", 37))

	// 3. Search for the indexed documents
	//
	// Build the request body.
	/*
		"suggest": {
		    "title_suggest": {
		      "prefix": "二",
		      "completion": {
		        "field": "title",
		        "size": 10,
		        "contexts": {
		          "title_perid": [ "2" ],
		          "title_subject": [ "2" ],
		          "title_type":["1"]

		        }
		      }
		    }
		  }
	*/
	var buf bytes.Buffer
	query := map[string]interface{}{
		"suggest": map[string]interface{}{
			"title_suggest": map[string]interface{}{
				"prefix": "一",
				"completion": map[string]interface{}{
					"field": "title",
					"size":  10,
					"contexts": map[string]interface{}{
						// "title_perid":   []string{"2"},
						// "title_subject": []string{"2"},
						"title_type": []string{"-1", "2"},
					},
				},
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	// Perform the search request.
	res, err = es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("test_suggest_zpj_2"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}
	// fmt.Println(res.String())

	resp := new(SuggestSearchResp)

	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	suggestData := resp.Suggest.TitleSuggest[0].Options

	suggestStr := []string{}
	if len(suggestData) > 0 {
		for _, v := range suggestData {
			suggestStr = append(suggestStr, v.Text)
		}
	}
	fmt.Println(suggestStr)

	// suggestS := []string{}
	// if len(resp.Suggest.TitleSuggest) > 0 {
	// for _, _ := range resp.Suggest.TitleSuggest[0] {
	//
	// }

	// // Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d sugget; took: %dms",
		res.Status(),
		len(suggestData),
		resp.Took,
	)
	// // Print the ID and document source for each hit.
	// for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
	// 	log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	// }

	log.Println(strings.Repeat("=", 37))
}

type SuggestSearchResp struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	// Shards   Shards  `json:"_shards"`
	// Hits     Hits    `json:"hits"`
	Suggest SuggestD `json:"suggest"`
}

// type Shards struct {
// 	Total      int `json:"total"`
// 	Successful int `json:"successful"`
// 	Skipped    int `json:"skipped"`
// 	Failed     int `json:"failed"`
// }

// type Total struct {
// 	Value    int    `json:"value"`
// 	Relation string `json:"relation"`
// }

// type Hits struct {
// 	Total    Total         `json:"total"`
// 	MaxScore interface{}   `json:"max_score"`
// 	Hits     []interface{} `json:"hits"`
// }
type Contexts struct {
	TitlePerid   []string `json:"title_perid"`
	TitleSubject []string `json:"title_subject"`
	TitleType    []string `json:"title_type"`
}
type Title struct {
	Input    []string `json:"input"`
	Contexts Contexts `json:"contexts"`
}
type Source struct {
	Title Title `json:"title"`
}

type Options struct {
	Text string `json:"text"`
	// Index    string   `json:"_index"`
	// Type     string   `json:"_type"`
	ID     string  `json:"_id"`
	Score  float64 `json:"_score"`
	Source Source  `json:"_source"`
	// Contexts Contexts `json:"contexts,omitempty"`
}
type TitleSuggest struct {
	Text    string    `json:"text"`
	Offset  int       `json:"offset"`
	Length  int       `json:"length"`
	Options []Options `json:"options"`
}
type SuggestD struct {
	TitleSuggest []TitleSuggest `json:"title_suggest"`
}
