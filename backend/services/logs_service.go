package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/pump-p/solidithai-assignment-2/backend/config"
)

func QueryLogs(query string) ([]map[string]interface{}, error) {
	esClient := config.ESClient

	// Construct search request
	var buf bytes.Buffer
	searchQuery := fmt.Sprintf(`{
        "query": {
            "multi_match": {
                "query": "%s",
                "fields": ["sender", "content"]
            }
        }
    }`, query)
	buf.WriteString(searchQuery)

	// Execute search request
	res, err := esClient.Search(
		esClient.Search.WithContext(context.Background()),
		esClient.Search.WithIndex("streaming_logs"),
		esClient.Search.WithBody(&buf),
		esClient.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return nil, fmt.Errorf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("Error: %s", res.String())
	}

	// Parse response
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, fmt.Errorf("Error parsing response body: %s", err)
	}

	var logs []map[string]interface{}
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		logs = append(logs, hit.(map[string]interface{})["_source"].(map[string]interface{}))
	}
	return logs, nil
}
