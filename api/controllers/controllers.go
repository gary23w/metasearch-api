package controllers

import (
	"context"
	enginesearch "github.com/gary23w/metasearch_api/internal/engine"
	_ "github.com/gary23w/metasearch_api/internal/providers/all"
	"github.com/gary23w/metasearch_api/internal/search"

	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	//test to see if the api is running
	c.JSON(200, gin.H{
		"helth": "online",
	})
}

func Search(c *gin.Context) {
	//get query parameters from post
	query := c.Param("query")
	ctx := context.Background()
	s, err := enginesearch.NewEngine(ctx)
	if err != nil {
		c.JSON(500, gin.H{
			"error":   err,
			"details": "error creating engine",
		})
	}
	limit := 10
	it := s.Search(ctx, search.Request{
		Query: query,
	})
	defer it.Close()
	//build a json response
	var results []interface{}
	for i := 0; i < limit && it.Next(ctx); i++ {
		r := it.Result()
		//fmt.Printf("%s - %q (%T)\n\n", r.GetURL(), r.GetTitle(), r)
		results = append(results, r)
	}
	if err := it.Err(); err != nil {
		c.JSON(500, gin.H{
			"error":   err,
			"details": "error getting results",
		})
	}
	tok := it.Token()
	//fmt.Println("\n\ntoken:", hex.EncodeToString([]byte(tok)))
	results = append(results, tok)
	if err := it.Err(); err != nil {
		c.JSON(500, gin.H{
			"error":   err,
			"details": "error getting token results",
		})
	}

	c.JSON(200, gin.H{
		"results": results,
	})
}

func Autocomplete(c *gin.Context) {
	//get query parameters from post
	query := c.Param("query")
	ctx := context.Background()
	s, err := enginesearch.NewEngine(ctx)
	if err != nil {
		c.JSON(500, gin.H{
			"error":   err,
			"details": "error creating engine",
		})
	}
	list, err := s.AutoComplete(ctx, query)
	if err != nil {
		c.JSON(500, gin.H{
			"error":   err,
			"details": "error getting autocomplete results",
		})
	}
	var results []interface{}
	for _, r := range list {
		results = append(results, r)
	}
	c.JSON(200, gin.H{
		"results": results,
	})
}
