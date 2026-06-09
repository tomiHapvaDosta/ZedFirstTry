package main

type requestParamsUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type requestParamsPost struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
