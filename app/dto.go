package app

type (
	RegisterRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	BookSeatRequest struct {
		SeatNumber string `json:"seat_number"`
		UserID     uint   `json:"user_id"`
	}
)
