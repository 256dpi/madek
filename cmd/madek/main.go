package main

import (
	"github.com/IAD-ZHDK/madek"
	"github.com/kr/pretty"
)

func main() {
	cmd := parseCommand()

	client := madek.NewClient(cmd.oAddress, cmd.oUsername, cmd.oPassword)

	if cmd.cSet {
		getSet(client, cmd.aID)
	}
}

func getSet(client *madek.Client, id string) {
	set, err := client.GetSet(id)
	pretty.Println(set, err)
}
