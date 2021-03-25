package main

import (
	"bytes"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/view", viewHandler)
	_ = http.ListenAndServe("localhost:8182", nil)
}

func getMongodbVersion(address string) []byte {
	// Set client options
	clientOptions := options.Client().ApplyURI(address)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	result := client.Database("admin").RunCommand(context.Background(), bson.D{{"buildInfo",1}})

	str, err := result.DecodeBytes()
	if err != nil {
		log.Fatal(err)
	}

	return preparePrettyJSON([]byte(str.String()))
}

func preparePrettyJSON(body []byte) []byte {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, body, "", "\t")
	if err != nil {
		log.Fatalf("JSON parse error: ", err)
	}

	return prettyJSON.Bytes()
}

const (
	defDBAddress     = "mongodb://localhost:27017"
	defResourcesPath = "/data/view/"
)

type Page struct {
	Body []byte
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	var address string
	if address == "" || len(os.Args[1]) < 2 {
		address = defDBAddress
	}

	var pagePath string
	if pagePath == "" || len(os.Args[1]) < 3 {
		pagePath = defResourcesPath
	}

	t, _ := template.ParseFiles("/data/view/view.html")
	err := t.Execute(w, &Page{Body: getMongodbVersion(address)})
	if err != nil {
		log.Fatal()
	}
}



