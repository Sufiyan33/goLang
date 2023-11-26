package main

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

var jwtkey = []byte("secret key")

var users = map[string]string{
	"user1": "Password1",
	"user2": "Password2",
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claim struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func Login(w http.ResponseWriter, r *http.Request) {}

func Home(w http.ResponseWriter, r *http.Request) {}

func Refresh(w http.ResponseWriter, r *http.Request) {}

func Logout(w http.ResponseWriter, r *http.Request) {}
