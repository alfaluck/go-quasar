package api

import "log"

type Auth struct {}

type Credentials struct {
	Email string `json:"email"`
	Password string `json:"password"`
	RememberMe bool `json:"rememberMe"`
}

type Result struct {
	Authorized bool
}

func (a *Auth) Login(c *Credentials, res *Result) (e error) {
	res.Authorized = true
	log.Print(`Tried to login`)
	return
}