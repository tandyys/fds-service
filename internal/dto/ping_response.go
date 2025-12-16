package dto

type PingResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type PingRandomNumberResponse struct {
	Status       int `json:"status"`
	RandomNumber int `json:"random_number"`
}
