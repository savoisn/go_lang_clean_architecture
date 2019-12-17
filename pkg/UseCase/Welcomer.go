package UseCase

type Welcomer struct {
	persistor IPersister
	language  string
}
type IPersister interface {
	Persist(something string) string
}

func NewWelcomer(persistor IPersister) *Welcomer {
	welcomer := Welcomer{persistor: persistor}
	return &welcomer
}

func (s *Welcomer) SayHello(person_name string) string {
	s.persistor.Persist(person_name)
	return "Hello " + person_name
}
