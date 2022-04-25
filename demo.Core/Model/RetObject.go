package Model

type RetObject struct {
	Success bool   `json:"success"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}
