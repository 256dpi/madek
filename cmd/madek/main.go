package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/IAD-ZHDK/madek"
	"github.com/patrickmn/go-cache"
)

func main() {
	cmd := parseCommand()

	client := madek.NewClient(cmd.oAddress, cmd.oUsername, cmd.oPassword)
	client.LogRequests = true

	if cmd.cFetch {
		fetch(client, cmd.aID)
	} else if cmd.cServer {
		server(client, cmd.oCache)
	}
}

func fetch(client *madek.Client, id string) {
	coll, err := client.CompileCollection(id)
	if err != nil {
		fmt.Printf("Error encountered: %s\n", err)
		return
	}

	bytes, err := json.MarshalIndent(coll, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func server(client *madek.Client, cacheEnabled bool) {
	mux := http.NewServeMux()

	requestCache := cache.New(cache.NoExpiration, cache.NoExpiration)

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get id
		id := strings.Trim(r.URL.Path, "/")
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// prepare coll
		var collection *madek.Collection

		// read from cache if allowed
		if cacheEnabled && r.URL.Query().Get("fresh") != "yes" {
			val, ok := requestCache.Get(id)
			if ok {
				collection = val.(*madek.Collection)
			}
		}

		// fetch not already there
		if collection == nil {
			coll, err := client.CompileCollection(id)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// save in cache if allowed
			if cacheEnabled {
				requestCache.Set(id, coll, cache.NoExpiration)
			}

			collection = coll
		}

		// marshal
		bytes, err := json.Marshal(collection)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// write
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	}))

	fmt.Println("+------------------------------------------------------------+")
	fmt.Println("| Running server on http://0.0.0.0:8080...                   |")
	fmt.Println("| Load data using the following pattern:                     |")
	fmt.Println("| > http://0.0.0.0:8080/82108639-c4a6-412d-b347-341fe5284caa |")
	fmt.Println("+------------------------------------------------------------+")

	http.ListenAndServe("0.0.0.0:8080", mux)
}
