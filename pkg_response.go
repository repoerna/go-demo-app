package main

type Response[T any] struct {
	Status  bool   `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
}

func NewResponse(msg string, data any) Response[any] {
	return Response[any]{
		Status:  true,
		Message: msg,
		Data:    data,
	}
}

func NewResponseError(msg string) Response[any] {
	return Response[any]{
		Status:  false,
		Message: msg,
	}
}
