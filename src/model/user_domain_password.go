package model

import (
	"encoding/hex"
	"github.com/kevinfjq/crud-golang/src/configuration/logger"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func (ud *userDomain) EncryptPassword() {
	hash, err := bcrypt.GenerateFromPassword([]byte(ud.password), 5)
	if err != nil {
		logger.Error("Error in generate password hash", err, zap.String("journey", "encrypt password"))
	}
	ud.password = hex.EncodeToString(hash)
}
