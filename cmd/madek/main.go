package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/IAD-ZHDK/madek"
)

func main() {
	cmd := parseCommand()

	client := madek.NewClient(cmd.oAddress, cmd.oUsername, cmd.oPassword)

	if cmd.cFetch {
		fetch(client, cmd.aID)
	} else if cmd.cServer {
		server(client)
	}
}

func fetch(client *madek.Client, id string) {
	set, err := client.CompileSet(id, true)
	if err != nil {
		fmt.Printf("Error encountered: %s\n", err)
		return
	}

	bytes, err := json.MarshalIndent(set, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func server(client *madek.Client) {
	http.HandleFunc("/set.json", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("missing id query parameter"))
			return
		}

		set, err := client.CompileSet(id, true)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		bytes, err := json.MarshalIndent(set, "", "  ")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write(bytes)
	})

	http.ListenAndServe("0.0.0.0:8888", nil)
}
