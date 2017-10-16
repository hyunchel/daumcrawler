// Package daumcrawler crawls on a keyword and
// saves on either JSON files.
package daumcrawler

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path"
	"encoding/json"
	"strings"
	"time"

	"github.com/hyunchel/daumapi"
)

var (
	buf bytes.Buffer
	logger = log.New(&buf, "INFO: ", log.Lshortfile)
)

// Combine JSON strings into a single JSON string.
// Use "totalCount", "title", "contents", and "datetime".
func combineJSON(contents []string) string {
	logger.Println("daumcrawler.combineJSON is running.")
	var documents []map[string]string
	var totalCount float64 = 0
	for _, c := range contents {
		m := decodeJSON(c).(map[string]interface{})
		meta := m["meta"].(map[string]interface{})
		totalCount += meta["total_count"].(float64)
		// Iterate through all documents and pick only needed attributes.
		for _, d := range m["documents"].([]interface{}) {
			detailedDoc := d.(map[string]interface{})
			shortDoc := map[string]string{}
			shortDoc["title"] = detailedDoc["title"].(string)
			shortDoc["contents"] = detailedDoc["contents"].(string)
			shortDoc["datetime"] = detailedDoc["datetime"].(string)
			// FIXME: This reinitializes the array every time.
			documents = append(documents, shortDoc)
		}
	}	
	combined := map[string]interface{}{
		"total_count": totalCount,
		"documents": documents,
	}
	return encodeJSON(combined)
}

// Decode arbitrary type of contents into an interface.
func decodeJSON(contents string) interface{} {
	logger.Println("daumcrawler.decodeJSON is running.")
	var f interface{}
	json.Unmarshal([]byte(contents), &f)
	return f
}

// Encode arbitrary type of contents into JSON string format.
func encodeJSON(contents interface{}) string {
	logger.Println("daumcrawler.encodeJSON is running.")
	b, err := json.Marshal(contents)
	if err != nil {
		log.Fatalf("Error occured: %v", err)
	}
	return string(b)
}

// Save JSON string into a file.
func saveInJSON(filename string, contents string) {
	logger.Println("daumcrawler.saveInJSON is running.")
	f := createAFile(filename)
	defer f.Close()

	if f != nil {
		logger.Printf("f: %v", f)
	}

	f.WriteString(contents)
}

// Make a string camelCase.
func convertToCamelCase(text string) string {
	words := strings.Split(text, " ")
	for i, t := range words[1:] {
		words[i + 1] = strings.Title(t)
	}
	return strings.Join(words, "")
}

// Make filename with time and keyword.
func makeFilename(keyword string) string {
	logger.Println("daumcrawler.makeFilename is running.")
	currentTime := time.Now().Unix()
	camelCasedKeyword := convertToCamelCase(keyword)
	return fmt.Sprintf("%v-%v", currentTime, camelCasedKeyword)
}
// Create and return a file.
func createAFile(filename string) *os.File {
	logger.Println("daumcrawler.createAFile is running.")
	base := path.Base("")
	name := path.Join(base, filename)
	logger.Printf("Opening file %q", name)
	f, err := os.Create(name)	
	if err != nil {
		logger.Fatalf("Error on opening file %q: %v", filename, err)
	}
	return f
}

// Crawl on the given keyword and save in JSON file.
func crawl(kakaoappkey string, keyword string) {
	logger.Println("daumcrawler.crawl is running.")
	results := []string{
		daumapi.Web(kakaoappkey, keyword),
		daumapi.Blog(kakaoappkey, keyword),
		daumapi.Tip(kakaoappkey, keyword),
		daumapi.Cafe(kakaoappkey, keyword),
	}
	filename := makeFilename(keyword)
	result := combineJSON(results)
	saveInJSON(filename, result)
}

func Run(kakaoappkey string, keyword string) {
	logger.Println("daumcrawler.Run is running.")
	crawl(kakaoappkey, keyword)
}

func RunWithLogging(kakaoappkey string, keyword string) {
	logger.Println("daumcrawler.RunWithLogging is running.")
	crawl(kakaoappkey, keyword)
	// Print logs.
	fmt.Print(&buf)
}