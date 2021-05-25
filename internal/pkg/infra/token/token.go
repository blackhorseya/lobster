package token

import (
	"time"

	"github.com/blackhorseya/lobster/internal/pkg/entity/er"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// JwtTokenClaims declare custom claims
type JwtTokenClaims struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

// Factory declare token's factory
type Factory struct {
	o      *Options
	logger *zap.Logger
}

// Options declare token's configuration
type Options struct {
	Name      string `json:"name" yaml:"name"`
	Signature string `json:"signature" yaml:"signature"`
}

// NewOptions serve caller to create Options
func NewOptions(v *viper.Viper) (*Options, error) {
	var (
		err error
		o   = new(Options)
	)

	if err = v.UnmarshalKey("app", o); err != nil {
		return nil, err
	}

	return o, nil
}

// New serve caller to create zap.Logger
func New(o *Options, logger *zap.Logger) (*Factory, error) {
	var f = &Factory{
		o:      o,
		logger: logger,
	}

	return f, nil
}

// NewToken serve caller new a json web token
func (f *Factory) NewToken(id int64, email string) (string, error) {
	loc, _ := time.LoadLocation("UTC")

	claims := JwtTokenClaims{
		ID:    id,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().In(loc).Add(7 * 86400 * time.Second).Unix(),
			Issuer:    f.o.Name,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(f.o.Signature))
	if err != nil {
		return "", err
	}

	return ss, nil
}

// ValidateToken serve caller to given signed token to validate token
func (f *Factory) ValidateToken(signedToken string) (claims *JwtTokenClaims, err error) {
	token, err := jwt.ParseWithClaims(signedToken, &JwtTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(f.o.Signature), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JwtTokenClaims)
	if !ok || !token.Valid {
		return nil, er.ErrValidateToken
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, er.ErrExpiredToken
	}

	return claims, nil
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(New, NewOptions)
