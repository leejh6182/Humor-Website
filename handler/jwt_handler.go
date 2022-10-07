package handler

import (
    "github.com/golang-jwt/jwt"
    "time"
)

func CreateJWT(userId string, name string, email string) (string, error) {
    signKey := []byte("TEST")

    token := jwt.New(jwt.SigningMethodHS256)
    
    claims := token.Claims.(jwt.MapClaims)
    claims["userid"] = userId
    claims["name"] = name
    claims["email"] = email
    claims["admin"] = false
    claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

    tk, error := token.SignedString(signKey)
    if error != nil {
        return "", error
    }

    return tk, nil
}


