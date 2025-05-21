package token

import "time"

// Maker is an interface for managing tokens
type Maker interface {

	//创建Token
	CreateToken(username string, duration time.Duration) (string, error)

	//验证Token
	VerifyToken(token string) (*Payload, error)
}
