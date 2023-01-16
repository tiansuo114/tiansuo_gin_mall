package util

import (
	"encoding/json"
	"errors"
	"gin_mall_tmp/cache"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Claims struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Authority int    `json:"authority"`
}

func (c Claims) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"ID":        c.ID,
		"UserName":  c.UserName,
		"Authority": c.Authority,
	})
}

type EmailClaims struct {
	UserID        uint   `json:"user_id"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	OperationType uint   `json:"operation_type"`
}

func (c EmailClaims) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"UserID":        c.UserID,
		"Email":         c.Email,
		"Password":      c.Password,
		"OperationType": c.OperationType,
	})
}

// GenerateToken 签发token
func GenerateToken(id uint, username string, authority int) (token string, err error) {
	claims := Claims{
		ID:        id,
		UserName:  username,
		Authority: authority,
	}
	claimString, _ := json.Marshal(claims)
	hash, err := hashString([]byte(claimString))
	if err != nil {
		LogrusObj.Infoln("token generate api", err)
		return "", err
	}
	err = cache.RedisClient.SetNX(hash, claims, 24*time.Hour).Err()
	if err != nil {
		LogrusObj.Infoln("token generate api", err)
		cache.RedisClient.Del(hash)
		return "", err
	}
	return hash, nil
}

// ParseToken 验证用户token
func ParseToken(token string) (*Claims, error) {
	val := &Claims{}
	err := cache.RedisClient.Get(token).Scan(val)
	if err != nil {
		LogrusObj.Infoln("token parse api", err)
		return nil, err
	} else if val == nil {
		LogrusObj.Infoln("token parse api", err)
		return nil, errors.New("获取token信息失败")
	}
	return val, nil
}

// GenerateEmailToken 签发EmailToken
func GenerateEmailToken(userId, Operation uint, email, password string) (token string, err error) {
	claims := EmailClaims{
		UserID:        userId,
		Email:         email,
		Password:      password,
		OperationType: Operation,
	}
	claimString, _ := json.Marshal(claims)
	hash, err := hashString([]byte(claimString))
	if err != nil {
		LogrusObj.Infoln("email token generate api", err)
		return "", err
	}
	err = cache.RedisClient.SetNX(hash, claims, 12*time.Hour).Err()
	if err != nil {
		LogrusObj.Infoln("email token generate api", err)
		cache.RedisClient.Del(hash)
		return "", err
	}
	return hash, nil
}

// ParseEmailToken  解析EmailToken
func ParseEmailToken(token string) (*EmailClaims, error) {
	val := &EmailClaims{}
	err := cache.RedisClient.Get(token).Scan(val)
	if err != nil {
		LogrusObj.Infoln("email token parse api", err)
		return nil, err
	} else if val == nil {
		LogrusObj.Infoln("email token parse api", err)
		return nil, errors.New("获取token信息失败")
	}
	return val, nil
}

func hashString(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
