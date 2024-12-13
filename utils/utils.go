package utils

import (
	"crypto/rand"
	"fmt"
	"setokoapp/constants"
	"time"

	"github.com/google/uuid"
)

func GenerateCurrentTime() time.Time {
	if constants.MODE_UNIT_TEST {
		currentDateTime, _ := time.Parse(constants.DATE_TIME_FORMAT, "2006-01-02 15:04:05")
		return currentDateTime
	}
	return time.Now()
}

func GenerateNewUUID() uuid.UUID {
	if constants.MODE_UNIT_TEST {
		return uuid.MustParse("3dd95dca-f16c-4795-9aea-e077156b79d9")
	}

	return uuid.New()
}

func GenTransactionId() string {
	if constants.MODE_UNIT_TEST {
		return "06010212345"
	}

	currentTime := fmt.Sprintf(time.Now().Format("060102"))
	randomString, _ := GenerateRandomString(5)
	transactionID := currentTime + randomString

	return transactionID
}

func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789"
	bytes, err := GenerateRandomBytes(n)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes), nil
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}
