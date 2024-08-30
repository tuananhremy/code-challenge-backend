package app

type (
	Handler struct {
		ds *DataStorage
	}
)

func NewHandler(ds *DataStorage, jwtSecret string) *Handler {
	return &Handler{
		ds: ds,
	}
}
