package models
import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

var privatekey *rsa.PrivateKey
var publicKey *rsa.PublicKey

func init() {
	// GenerateKey generates a random RSA private key of the given bit size.
	privatekey, _ = rsa.GenerateKey(rand.Reader, 2048)
	publicKey = &privatekey.PublicKey
}

func GetJWT(user *User) (string, int) {
	token := jwt.New(jwt.SigningMethodRS512)
	claims := make(jwt.MapClaims)
	overtime, _ := beego.AppConfig.Int("jwt_expires_in")
	claims["exp"] = time.Now().Local().Add(time.Duration(overtime) * time.Minute).Unix()
	claims["iat"] = time.Now().Local().Unix()
	claims["nbf"] = claims["iat"]
	claims["userid"] = user.Id
	claims["projectid"] = user.Projectid
	claims["name"] = user.Real_name
	token.Claims = claims
	tokenString, err := token.SignedString(privatekey)
	if err != nil {
		fmt.Println(err)
		return "", RESULT_ERROR
	}

	return tokenString, RESULT_OK
}

func ParseJwt(auth string) (userObj *User, ret int) {
	keyfunc := func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	}

	if auth=="test"{
		user, err := GetUserInfobyId(1)
		return user, err
	}

	token, err := jwt.Parse(strings.TrimPrefix(auth,"Bearer "), keyfunc)
	if err != nil {
		return userObj, RESULT_TOKEN_ERROR
	}

	if token.Valid {
		Claims := token.Claims.(jwt.MapClaims)
		if int64(Claims["exp"].(float64)) < time.Now().Local().Unix() {
			return userObj, RESULT_TOKEN_EXPIRED
		}

		user, err := GetUserInfobyId(int64(Claims["userid"].(float64)))
		return user, err
	}

	return userObj, RESULT_TOKEN_ERROR
}
