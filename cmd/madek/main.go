package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/IAD-ZHDK/madek"
	"github.com/gin-gonic/gin"
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

func server(client *madek.Client, cacheEnabled bool) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	requestCache := cache.New(cache.NoExpiration, cache.NoExpiration)

	router.GET("/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.AbortWithError(http.StatusBadRequest, errors.New("missing id query parameter"))
			return
		}

		if ctx.Query("fresh") != "yes" {
			val, ok := requestCache.Get(id)
			if ok {
				ctx.JSON(http.StatusOK, val.(*madek.Set))
				return
			}
		}

		set, err := client.CompileSet(id)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		if cacheEnabled {
			requestCache.Set(id, set, cache.NoExpiration)
		}

		ctx.JSON(http.StatusOK, set)
	})

	fmt.Println("+------------------------------------------------------------+")
	fmt.Println("| Running server on http://0.0.0.0:8080...                   |")
	fmt.Println("| Load data using the following pattern:                     |")
	fmt.Println("| > http://0.0.0.0:8080/82108639-c4a6-412d-b347-341fe5284caa |")
	fmt.Println("+------------------------------------------------------------+")

	router.Run("0.0.0.0:8080")
}
