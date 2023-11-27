package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Here, we are using hardcoded secret key but in project store in server.
var jwtkey = []byte("my_secret_key")

var users = map[string]string{
	"user1": "Password1",
	"user2": "Password2",
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Create a struct that will be encoded to a JWT.
// We add jwt.RegisteredClaims as an embedded type, to provide fields like expiry time
type Claim struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func Login(w http.ResponseWriter, r *http.Request) {
	// Get the JSON body and decode into credentials. If the structure of the body is wrong, return an HTTP error
	fmt.Println("Welcome to login screen")
	var cred Credentials
	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Get the expected password from our in memory map
	expectedPassword, ok := users[cred.Username]

	/* If a password exists for the given user AND, if it is the same as the password we received, the we can move ahead
	if NOT, then we return an "Unauthorized" status*/

	if !ok || expectedPassword != cred.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Declare the expiration time of the token, Create the JWT claims, which includes the username and expiry time
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claim{
		Username: cred.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	// Declare the token with the algorithm used for signing, and the claims & create jwt string, If there is an error in creating the JWT return an internal server error
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	/* Finally, we set the client cookie for "token" as the JWT we just generated
	we also set an expiry time which is the same as the token itself */
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func Home(w http.ResponseWriter, r *http.Request) {
	// We can obtain the session token from the requests cookies, which come with every request, If the cookie is not set, return an unauthorized status
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := c.Value

	// Get the JWT string from the cookie & initialize a new instance of `Claim`
	claims := &Claim{}
	/*
		Parse the JWT string and store the result in `claims`. Note that we are passing the key in this method as well. This method will return an error
		if the token is invalid (if it has expired according to the expiry time we set on sign in), or if the signature does not match
	*/
	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(tokn *jwt.Token) (any, error) {
		return jwtkey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// Finally, return the welcome message to the user, along with their username given in the token
	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tknStr := c.Value
	claims := &Claim{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (any, error) {
		return jwtkey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	/*
		We ensure that a new token is not issued until enough time has elapsed In this case, a new token will only be issued if the old token is within
		30 seconds of expiry. Otherwise, return a bad request status
	*/
	if time.Until(claims.ExpiresAt.Time) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Now, create a new token for the current use, with a renewed expiration time
	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Set the new token as the users `token` cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Expires: time.Now(),
	})
}
