package util

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"strings"
	"time"
)

const JWT_SERCRET_KEY = "my_secret"

func GenerateToken(userId string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"userId": userId,
		//"Username": req.Username,
		//"iss":      "issue",
		"sub": "login",
		//"aud": "aud",
		"exp": jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
		//"nbf" : jwt.NewNumericDate(time.Now().Add(30 * time.Minute)), //NotBefore
		"jti": uuid.NewString(), //unique id jwt
	})

	tokenString, err := token.SignedString([]byte(JWT_SERCRET_KEY))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(Token string) (*jwt.Token, error) {
	if Token == "" {
		return nil, errors.New("unauthorized")
	}

	auth := strings.TrimPrefix(Token, "Bearer ")

	t, err := jwt.Parse(auth, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(JWT_SERCRET_KEY), nil
	})

	if t.Valid {
		return t, nil
	} else if errors.Is(err, jwt.ErrInvalidKey) {
		return nil, jwt.ErrInvalidKey
	} else if errors.Is(err, jwt.ErrInvalidKeyType) {
		return nil, jwt.ErrInvalidKeyType
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		return nil, jwt.ErrTokenMalformed
	} else if errors.Is(err, jwt.ErrTokenUnverifiable) {
		return nil, jwt.ErrTokenUnverifiable
	} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
		return nil, jwt.ErrTokenSignatureInvalid
	} else if errors.Is(err, jwt.ErrTokenRequiredClaimMissing) {
		return nil, jwt.ErrTokenRequiredClaimMissing
	} else if errors.Is(err, jwt.ErrTokenExpired) {
		return nil, jwt.ErrTokenExpired
	} else if errors.Is(err, jwt.ErrTokenInvalidSubject) {
		return nil, jwt.ErrTokenInvalidSubject
	} else if errors.Is(err, jwt.ErrTokenInvalidId) {
		return nil, jwt.ErrTokenInvalidId
	} else if errors.Is(err, jwt.ErrTokenInvalidClaims) {
		return nil, jwt.ErrTokenInvalidClaims
	} else if errors.Is(err, jwt.ErrInvalidType) {
		return nil, jwt.ErrInvalidType
		/*} else if errors.Is(err, jwt.ErrTokenInvalidAudience) {
		return nil, jwt.ErrTokenInvalidAudience*/
		/*} else if errors.Is(err, jwt.ErrTokenNotValidYet) {
		return nil, jwt.ErrTokenNotValidYet*/
		/*}  else if errors.Is(err, jwt.ErrHashUnavailable) {
		return nil, jwt.ErrHashUnavailable*/
		/*} else if errors.Is(err, jwt.ErrTokenUsedBeforeIssued) {
		return nil, jwt.ErrTokenUsedBeforeIssued*/
		/*} else if errors.Is(err, jwt.ErrTokenInvalidIssuer) {
		return nil, jwt.ErrTokenInvalidIssuer*/
	} else {
		return nil, err
	}
}
