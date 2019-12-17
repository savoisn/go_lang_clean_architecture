package UseCase

func (s *Welcomer) SayHelloRequest(request WelcomerRequest) string {
	person_name := request.NameToWelcome
	return s.SayHello(person_name)
}
