package helper

import (
	"encoding/base64"
	"encoding/hex"
	"math/rand"
	"time"

	errHelper "zayyid-go/domain/shared/helper/error"

	"github.com/thanhpk/randstr"
	"golang.org/x/crypto/bcrypt"
)

type ResponseGenerate struct {
	NewApiKey    string `json:"new_api_key"`
	NewSecretKey string `json:"new_secret_key"`
}

func generateSecreteKey(randomKey string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(randomKey), 14)
	return string(bytes), err
}

func decodeKey(companyId string, key string, length int) string {
	stringKey := companyId + randomChar(length) + key + randomChar(length)

	var encodedString = base64.StdEncoding.EncodeToString([]byte(stringKey))
	src := []byte(encodedString)
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	return string(dst)
}

func randomChar(strLength int) string {
	var charset = "!@#$%^&*+_=abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, strLength)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GenerateApiAndSecretKey(companyId string, apiKey string) (response ResponseGenerate, err error) {

	if apiKey == "" {
		apiKey = randstr.String(7)
	}

	newApiKey := decodeKey(companyId, apiKey, 5)
	newSecretKey, err := generateSecreteKey(decodeKey(companyId, apiKey, 7))
	if err != nil {
		err = errHelper.NewIntegrationError(errHelper.ERROR_INTERNAL_SERVER, "Generate Key : "+err.Error())
		return
	}

	response = ResponseGenerate{
		NewApiKey:    newApiKey,
		NewSecretKey: newSecretKey,
	}

	return
}
