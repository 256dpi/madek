package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/256dpi/madek"
)

var address = flag.String("address", "https://medienarchiv.zhdk.ch", "The address of the Madek instance.")
var username = flag.String("username", "", "The username for authentication.")
var password = flag.String("password", "", "The password for authentication.")

func main() {
	// parse flags
	flag.Parse()

	// get id
	id := flag.Arg(0)

	// prepare client
	client := madek.NewClient(*address, *username, *password)

	// compile collection
	coll, err := client.CompileCollection(id)
	if err != nil {
		fmt.Printf("Error encountered: %s\n", err)
		return
	}

	// encode
	bytes, err := json.MarshalIndent(coll, "", "  ")
	if err != nil {
		panic(err)
	}

	// print
	fmt.Println(string(bytes))
}
