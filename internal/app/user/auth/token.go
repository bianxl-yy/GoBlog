package auth

import (
	"fmt"
	"github.com/bianxl-yy/GoBlog/app/utils"
	"github.com/bianxl-yy/GoBlog/internal/app/user/base"
	"github.com/labstack/echo"
)

var tokens map[string]*Token

type Token struct {
	Value      string
	UserId     int
	CreateTime int64
	ExpireTime int64
}

// check token is valid or expired.
func (t *Token) IsValid() bool {
	if base.GetUserById(t.UserId) == nil {
		return false
	}
	return t.ExpireTime > utils.Now()
}

// create new token from user and context.
func CreateToken(u *base.User, c echo.Context, expire int64) *Token {
	t := new(Token)
	t.UserId = u.Id
	t.CreateTime = utils.Now()
	t.ExpireTime = t.CreateTime + expire
	t.Value = utils.Sha1(fmt.Sprintf("%s-%s-%d-%d", c.RealIP(), c.Request().UserAgent, t.CreateTime, t.UserId))
	tokens[t.Value] = t
	go SyncTokens()
	return t
}

// get token by token value.
func GetTokenByValue(v string) *Token {
	return tokens[v]
}

// get tokens of given user.
func GetTokensByUser(u *base.User) []*Token {
	ts := make([]*Token, 0)
	for _, t := range tokens {
		if t.UserId == u.Id {
			ts = append(ts, t)
		}
	}
	return ts
}

// remove a token by token value.
func RemoveToken(v string) {
	delete(tokens, v)
	go SyncTokens()
}

// clean all expired tokens in memory.
// do not write to json.
func CleanTokens() {
	for k, t := range tokens {
		if !t.IsValid() {
			delete(tokens, k)
		}
	}
}

// write tokens to json.
// it calls CleanTokens before writing.
func SyncTokens() {
	CleanTokens()
	Storage.Set("tokens", tokens)
}

// load all tokens from json.
func LoadTokens() {
	tokens = make(map[string]*Token)
	Storage.Get("tokens", &tokens)
}
