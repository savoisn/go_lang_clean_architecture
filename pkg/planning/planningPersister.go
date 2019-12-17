package planning

import (
	"gitlab.talanlabs.com/okapi/OkapiPlanning/pkg/entity"
)

type PlanningPersister interface {
	SavePlanning(planning entity.Planning) (error, entity.Planning)
	GetPlannings() (error, []entity.Planning)
}
