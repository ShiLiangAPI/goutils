package pwdutils

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"strings"
)

/////////////////////////// pbkdf2_sha256加密 ///////////////////////////

// PasswordHash pbkdf2_sha256加密,与python同步
func (obj *PwdUtil) PasswordHash(password string) (string, error) {

	// pbkdf2加密 <--- 关键
	hash := pbkdf2.Key([]byte(password), []byte(obj.salt), obj.iterations, sha256.Size, sha256.New)

	// base64编码成为固定长度的字符串
	b64SaltHash := strings.TrimRight(base64.StdEncoding.EncodeToString([]byte(obj.salt)), "=")
	b64PasswordHash := strings.TrimRight(base64.StdEncoding.EncodeToString(hash), "=")

	// 最终字符串拼接成pbkdf2_sha256密钥格式
	pwdHash := fmt.Sprintf("$%s$%d$%s$%s", "pbkdf2-sha256", obj.iterations, b64SaltHash, b64PasswordHash)

	return pwdHash, nil
}

// PasswordVerify 密码校验
func (obj *PwdUtil) PasswordVerify(pwd, hash string) bool {
	pwdHash, _ := obj.PasswordHash(pwd)

	return pwdHash == hash
}
