package utils

import (
	"encoding/json"

	"github.com/ChocolateAceCream/telescope/backend/singleton"
	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
)

func Decoder(tokenStr string) (claimsJSON []byte, err error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenStr, jwt.MapClaims{})
	if err != nil {
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	claimsJSON, _ = json.MarshalIndent(claims, "", "  ")
	singleton.Logger.Info("decode result", zap.String("claims", string(claimsJSON)))
	return
}
