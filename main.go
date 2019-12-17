package main

import (
	"gitlab.talanlabs.com/okapi/OkapiPlanning/pkg/UseCase"
	db "gitlab.talanlabs.com/okapi/OkapiPlanning/pkg/infra/db"
	memory "gitlab.talanlabs.com/okapi/OkapiPlanning/pkg/infra/memory"
	"gitlab.talanlabs.com/okapi/OkapiPlanning/pkg/infra/web"
	planning "gitlab.talanlabs.com/okapi/OkapiPlanning/pkg/planning"
)

func main() {

	db_persistor := db.NewDbPersitor("host_name", 1234)
	welcomer := UseCase.NewWelcomer(db_persistor)
	//web := web.NewGinServer(welcomer)

	inMemoryPersister := memory.NewInMemoryPersister()
	planningUseCase := planning.NewPlanningUseCase(&inMemoryPersister)
	web := web.NewGinServer(welcomer, planningUseCase)

	web.StartServer()

}
