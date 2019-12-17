package infra

import "gitlab.talanlabs.com/okapi/OkapiPlanning/pkg/UseCase"

type DbPersistor struct {
	host_name string
	port      int
	content   []string
}

func NewDbPersitor(host_name string, port int) UseCase.IPersister {
	return &DbPersistor{host_name: host_name, port: port}
}

func (f *DbPersistor) Persist(something string) string {
	f.content = append(f.content, something)
	return something
}

func (f *DbPersistor) GetContent() []string {
	return f.content
}
