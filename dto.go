package main

import "github.com/google/uuid"

type AddUserRequestBody struct {
	Name string `json:"name"`
}

type AddUserResponseBody struct {
	ID uuid.UUID `json:"id"`
}
