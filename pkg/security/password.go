package security

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	SALT_CHARS                = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	SALT_LENGTH               = 8
	DEFAULT_PBKDF2_ITERATIONS = 50000
)

func GenSalt(length int) (string, error) {
	var salt string
	if length <= 0 {
		return "", errors.New("生成salt长度出错")
	}
	rand.Seed(time.Now().Unix())
	for i := 0; i < length; i++ {
		salt += string(SALT_CHARS[rand.Intn(len(SALT_CHARS))])
	}
	return salt, nil
}

func GeneratePasswordHash(password string) string {
	salt, _ := GenSalt(SALT_LENGTH)
	method := "pbkdf2:sha256"
	h, actualMethod := hashInternal(method, salt, password)
	return fmt.Sprintf("%s$%s$%s", actualMethod, salt, h)
}

func CheckPasswordHash(pwhash string, password string) bool {
	if strings.Count(pwhash, "$") < 2 {
		return false
	}
	ss := strings.SplitN(pwhash, "$", 3)
	if len(ss) != 3 {
		return false
	}
	method, salt, hashval := ss[0], ss[1], ss[2]
	rv, _ := hashInternal(method, salt, password)
	return rv == hashval
}

func hashInternal(method string, salt string, password string) (string, string) {
	if method == "plain" {
		return password, method
	}
	rv := ""
	isPbkdf2 := false
	actualMethod := method
	iterations := 0
	if strings.HasPrefix(method, "pbkdf2:") {
		args := strings.Split(method[7:], ":")
		if len(args) == 1 {
			method = args[0]
			iterations = DEFAULT_PBKDF2_ITERATIONS
		} else if len(args) == 2 {
			method = args[0]
			iterations, _ = strconv.Atoi(args[1])
		}
		isPbkdf2 = true
		actualMethod = fmt.Sprintf("pbkdf2:%s:%d", method, iterations)
	}
	if isPbkdf2 {
		rv = hex.EncodeToString(pbkdf2.Key([]byte(password), []byte(salt), iterations, 32, sha256.New))
	} else if salt != "" {
		h := hmac.New(sha256.New, []byte(salt))
		h.Write([]byte(password))
		// Get result and encode as hexadecimal string
		rv = hex.EncodeToString(h.Sum(nil))
	}
	return rv, actualMethod
}
