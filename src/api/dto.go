package api

// UserResponse - структура ответа для информации о пользователе
type UserResponse struct {
	ID    int    `json:"id" example:"1"`
	Name  string `json:"name" example:"John Doe"`
	Email string `json:"email" example:"john.doe@example.com"`
}

type Message struct {
	Status  string `json:"status" example:"not ok"`
	Message string `json:"message" example:"Some message"`
}

type BadRequest struct {
	Detail Message `json:"detail"`
}
