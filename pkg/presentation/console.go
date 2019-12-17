package presentation

import (
	"fmt"

	db "gitlab.talanlabs.com/okapi/OkapiPlanning/pkg/infra/db"
	"gitlab.talanlabs.com/okapi/OkapiPlanning/pkg/use_case"
)

func Console() {
	// file_persistor := infra.NewFilePersistor("/home/nico/myfile")
	db_persistor := db.NewDbPersitor("host_name", 1234)

	// toto := Toto{}
	welcomer := use_case.NewWelcomer(db_persistor)
	hello := welcomer.SayHello("Diane")

	fmt.Println(hello)
}

type Toto struct {
	name string
}

func (t Toto) Persist(something string) string {
	return "toto"
}
