// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ItemInput struct {
	Name string `json:"name"`
}

type NoopInput struct {
	ClientMutationID *string `json:"clientMutationId"`
}

type NoopPayload struct {
	ClientMutationID *string `json:"clientMutationId"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserInput struct {
	Name string `json:"name"`
}
