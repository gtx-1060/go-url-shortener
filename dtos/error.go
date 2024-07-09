package dtos

type Error struct {
	Message string `json:"message"`
}

func ErrorResponse(msg string) Error {
	return Error{msg}
}
