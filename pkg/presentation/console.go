package presentation

import (
	"fmt"

	"gitlab.talanlabs.com/okapi/OkapiPlanning/pkg/UseCase"
	infra "gitlab.talanlabs.com/okapi/OkapiPlanning/pkg/infra/db"
)

func Console() {
	// file_persistor := infra.NewFilePersistor("/home/nico/myfile")
	db_persistor := infra.NewDbPersitor("host_name", 1234)

	// toto := Toto{}
	welcomer := UseCase.NewWelcomer(db_persistor)
	hello := welcomer.SayHello("Diane")

	fmt.Println(hello)
}

type Toto struct {
	name string
}

func (t Toto) Persist(something string) string {
	return "toto"
}
