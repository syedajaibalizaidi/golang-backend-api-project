package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "secretkey"

// for generating tokens
func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"expire": time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

// for verifying the token , if the token is invalid. get the token as the input or parameter. and in there we want to verify that token.
// now for this jwt package also offers parse method which parses a received token and extracts the information that stored in that token. and tells weather its a valid token.
// first it wants a token then secretkey for which we pass anonymous function to parse.
// *jwt.SigningMethodHMAC type checking syntax where we can access a type like property with a dot on a field to check the value stored in that field is of that type.
func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) { // anonymous function
		_, ok := token.Method.(*jwt.SigningMethodHMAC) // returning secret , checking token either it signed or not.

		// checking if the method used for signing the token is not ok a diff method not the method we used in that case we return an error.
		if !ok {
			return nil, errors.New("Unexpected error") // nil for secretkey as we got a invalid token. so sharing secert key is not good. instead here create a new error with error package.
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("token parsing failed.") // return 0 whenever we get a error.
	}

	// checking if its valid for parsing the token data.
	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return 0, errors.New("Invalid token") // return 0 whenever we get a error.
	}

	// if we get a valid token, we can use our parsed token. using type-checking syntax below where we add a dot and paranthesis
	// here we wanna check whether the claims we got for this token is of the jwt.MapClaims() type
	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("Invalid token claims") // return 0 whenever we get a error.
	}

	// after above claim if check we will be able to access data in that claims variable here
	// email, ok := claims["email"].(string)  // using these brackets to get hold of the email key. again checking the type.
	userId, ok := claims["userId"].(int64) // using these brackets to get hold of the userId key.type conversion as it stored in float64

	return userId, nil
}
