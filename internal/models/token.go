package models

import "github.com/golang-jwt/jwt"

// TODO to no forget change the password
const Secret = "something"

type TokenClaims struct {
	jwt.StandardClaims
	User
}

type User struct {
	Id   uint
	Role string
}

type auth string

var UserCtx auth = "models.UserCtx"
