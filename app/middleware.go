package app

const (
	contextKeyUserID = "userID"
)

type Middleware struct {
}

func NewMiddleware(jwtSecret string) *Middleware {
	return &Middleware{}
}
