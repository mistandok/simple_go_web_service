package clients

type Logic interface {
	SayHello(userID string) (string, error)
}

type Logger interface {
	Log(message string)
}
