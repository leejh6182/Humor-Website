package handler

import (
    "github.com/golang-jwt/jwt"
    "time"
    . "fmt"
    "strconv"
)

var SIGNKEY = []byte("TEST")

func CreateToken(claim authTokenClaim) (string, error) {

    token := jwt.New(jwt.SigningMethodHS256)
    
    claims := token.Claims.(jwt.MapClaims)
    claims["userid"] = claim.UserId
    claims["email"] = claim.Email
    claims["level"] = strconv.Itoa(claim.Level)
    claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

    tk, error := token.SignedString(SIGNKEY)
    if error != nil {
        return "", error
    }

    return tk, nil
}

func ValidateToken(tokenString string) (*jwt.Token, error) {

    //Parse token and validate signature
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, Errorf("Unexpected signing method: %v", token.Header["alg"])
        }

        return SIGNKEY, nil
    })
}


