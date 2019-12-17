package infra

import "gitlab.talanlabs.com/okapi/OkapiPlanning/pkg/entity"

type InMemoryPersister struct {
	plannings []entity.Planning
}

func NewInMemoryPersister() InMemoryPersister {
	return InMemoryPersister{}
}

func (i *InMemoryPersister) SavePlanning(planning entity.Planning) (error, entity.Planning) {
	i.plannings = append(i.plannings, planning)
	return nil, planning
}

func (i *InMemoryPersister) GetPlannings() (error, []entity.Planning) {
	return nil, i.plannings
}
