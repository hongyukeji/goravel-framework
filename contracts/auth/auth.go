package auth

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/goravel/framework/contracts/http"
)

//go:generate mockery --name=Auth
type Auth interface {
	Guard(name string) Auth
	Parse(ctx http.Context, token string) (*Payload, error)
	ParseToken(ctx http.Context, token string) (jwt.Token, error)
	User(ctx http.Context, user any) error
	Login(ctx http.Context, user any) (token string, err error)
	LoginUsingID(ctx http.Context, id any) (token string, err error)
	Refresh(ctx http.Context) (token string, err error)
	Logout(ctx http.Context) error
}

type Payload struct {
	Guard    string
	Key      string
	ExpireAt time.Time
	IssuedAt time.Time
}
