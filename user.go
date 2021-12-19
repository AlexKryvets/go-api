package go_api

type User struct {
	id   int    `json:"-"`
	name string `json:"name"`
}
