package encrypt

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/spf13/viper"
)

type Service interface {
	Password(password string) string
}

type service struct {
	secret string
}

func New() Service {
	return &service{
		secret: viper.GetString("PASSWORD_SALT"),
	}
}

func (s service) Password(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(s.secret))
	hasher.Write([]byte(password))

	hashedBytes := hasher.Sum(nil)

	return hex.EncodeToString(hashedBytes)
}
