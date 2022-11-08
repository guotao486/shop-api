/*
 * @Author: GG
 * @Date: 2022-10-26 16:10:30
 * @LastEditTime: 2022-11-07 14:46:43
 * @LastEditors: GG
 * @Description: jwt
 * @FilePath: \shop-api\utils\jwt\jwt_helper.go
 *
 */
package jwt

import (
	"encoding/json"
	"log"

	"github.com/dgrijalva/jwt-go"
)

// token解码后结构体
type DecodeToken struct {
	Iat      int    `json:"iat"`
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Iss      string `json:"iss"`
	IsAdmin  bool   `json:"isAdmin"`
}

/**
 * @description: 生成token
 * @param {*jwt.Token} claims
 * @param {string} secret
 * @return {*}
 */
func GenerateToken(claims *jwt.Token, secret string) (token string, err error) {
	hmacSecret := []byte(secret)
	return claims.SignedString(hmacSecret)
}

/**
 * @description: 解析token
 * @param {string} token
 * @param {string} secret
 * @return {*}
 */
func VerifyToken(token string, secret string) *DecodeToken {
	hmacSecretString := secret
	hmacSecret := []byte(hmacSecretString)

	decode, err := jwt.Parse(
		token, func(t *jwt.Token) (interface{}, error) {
			return hmacSecret, nil
		})

	// 解析错误
	if err != nil {
		return nil
	}

	// 判断token是否失效
	if !decode.Valid {
		return nil
	}

	// 类型断言
	decodeClaims := decode.Claims.(jwt.MapClaims)
	log.Print(decodeClaims)

	// 封装至结构体
	var decodeToken DecodeToken
	jsonString, _ := json.Marshal(decodeClaims)
	jsonErr := json.Unmarshal(jsonString, &decodeToken)
	if jsonErr != nil {
		log.Print(jsonErr)
		return nil
	}
	return &decodeToken
}
