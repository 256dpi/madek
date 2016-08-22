package main

import (
	"encoding/json"
	"fmt"

	"github.com/IAD-ZHDK/madek"
)

func main() {
	cmd := parseCommand()

	client := madek.NewClient(cmd.oAddress, cmd.oUsername, cmd.oPassword)

	if cmd.cSet {
		getSet(client, cmd.aID)
	}
}

func getSet(client *madek.Client, id string) {
	set, err := client.CompileSet(id)
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
