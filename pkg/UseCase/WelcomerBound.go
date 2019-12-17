package UseCase

type WelcomerBound interface {
	SayHelloRequest(request WelcomerRequest) string
}
