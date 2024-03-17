package utils

import (
	"crypto/sha1"
	"crypto/sha512"
	"encoding/hex"
)

func ParseStringToSha512(str string) string {
	hasher := sha512.New()
	hasher.Write([]byte(str))
	hashBytes := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}

func ParseStringToSha1(str string) string {
	hasher := sha1.New()
	hasher.Write([]byte(str))
	hashBytes := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}
