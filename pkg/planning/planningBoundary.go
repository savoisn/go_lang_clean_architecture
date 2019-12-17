package planning

import "gitlab.talanlabs.com/okapi/OkapiPlanning/pkg/entity"

type PlanningBoundary interface {
	SavePlanning(planning entity.Planning) (error, entity.Planning)
	GetPlannings() (error, []entity.Planning)
}
