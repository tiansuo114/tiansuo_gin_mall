package util

//
//import (
//	"github.com/dgrijalva/jwt-go"
//	"time"
//)
//
//var jwtSecret = []byte("woshinantong")
//
//type Claims struct {
//	ID        uint   `json:"id"`
//	UserName  string `json:"user_name"`
//	Authority int    `json:"authority"`
//	jwt.StandardClaims
//}
//
//type EmailClaims struct {
//	UserID        uint   `json:"user_id"`
//	Email         string `json:"email"`
//	Password      string `json:"password"`
//	OperationType uint   `json:"operation_type"`
//	jwt.StandardClaims
//}
//
//// GenerateToken 签发token
//func GenerateToken(id uint, username string, authority int) (token string, err error) {
//	nowTime := time.Now()
//	expireTime := nowTime.Add(24 * time.Hour)
//	claims := Claims{
//		ID:        id,
//		UserName:  username,
//		Authority: authority,
//		StandardClaims: jwt.StandardClaims{
//			ExpiresAt: expireTime.Unix(),
//			Issuer:    "FanOne-Mall",
//		},
//	}
//	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	token, err = tokenClaims.SignedString(jwtSecret)
//	return
//}
//
////// ParseToken 验证用户token
////func ParseToken(token string) (*Claims, error) {
////	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
////		return jwtSecret, nil
////	})
////	if err != nil {
////		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
////			return claims, nil
////		}
////	}
////	return nil, err
////}
//
////ParseToken 验证用户token
//func ParseToken(token string) (*Claims, error) {
//	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
//		return jwtSecret, nil
//	})
//	if tokenClaims != nil {
//		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
//			return claims, nil
//		}
//	}
//	return nil, err
//}
//
//// GenerateEmailToken 签发EmailToken
//func GenerateEmailToken(userId, Operation uint, email, password string) (token string, err error) {
//	nowTime := time.Now()
//	expireTime := nowTime.Add(12 * time.Hour)
//	claims := EmailClaims{
//		UserID:        userId,
//		Email:         email,
//		Password:      password,
//		OperationType: Operation,
//		StandardClaims: jwt.StandardClaims{
//			ExpiresAt: expireTime.Unix(),
//			Issuer:    "FanOne-Mall",
//		},
//	}
//	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	token, err = tokenClaims.SignedString(jwtSecret)
//	return
//}
//
////// ParseEmailToken  解析EmailToken
////func ParseEmailToken(token string) (*EmailClaims, error) {
////	tokenClaims, err := jwt.ParseWithClaims(token, &EmailClaims{}, func(token *jwt.Token) (interface{}, error) {
////		return jwtSecret, nil
////	})
////	if tokenClaims != nil {
////		if claims, ok := tokenClaims.Claims.(*EmailClaims); ok && tokenClaims.Valid {
////			return claims, nil
////		}
////	}
////	return nil, err
////}
//
////ParseEmailToken 验证邮箱验证token
//func ParseEmailToken(token string) (*EmailClaims, error) {
//	tokenClaims, err := jwt.ParseWithClaims(token, &EmailClaims{}, func(token *jwt.Token) (interface{}, error) {
//		return jwtSecret, nil
//	})
//	if tokenClaims != nil {
//		if claims, ok := tokenClaims.Claims.(*EmailClaims); ok && tokenClaims.Valid {
//			return claims, nil
//		}
//	}
//	return nil, err
//}
