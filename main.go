package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	infradb "gitlab.talanlabs.com/okapi/OkapiPlanning/pkg/infra/file"
	"gitlab.talanlabs.com/okapi/OkapiPlanning/pkg/presentation"
)

func main() {
	fmt.Println("Okapi V2")

	fp := infradb.NewFilePersistor("Coucou")
	fp.Persist("Coucou")
	fp.Persist("Coucou")
	fp.Persist("Coucou")
	fp.Persist("Coucou")
	fp.Persist("Coucou")

	for _, s := range fp.GetContent() {
		fmt.Println(s)
	}
	presentation.Console()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		// ControllerBoundary(RequestModel) -> UseCase -> Reply Result
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
