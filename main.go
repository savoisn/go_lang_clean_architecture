package main

import (
	"fmt"

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
}
