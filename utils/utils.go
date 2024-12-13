package utils

import (
	"crypto/rand"
	"errors"
	"fmt"
	"reflect"
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

func ValidateRequestData(request interface{}) error {
	dataStruct := reflect.TypeOf(request)
	dataValue := reflect.ValueOf(request)

	for i := 0; i < dataStruct.NumField(); i++ {
		field := dataStruct.Field(i)
		fieldValue := dataValue.Field(i)

		if field.Tag.Get("validate") == "required" {
			if err := validateFieldValue(fieldValue, field); err != nil {
				return err
			}
		}
	}

	return nil
}

func validateFieldValue(fieldValue reflect.Value, field reflect.StructField) error {
	if fieldValue.IsValid() {
		if fieldValue.IsZero() {
			return errors.New(fmt.Sprintf("%v is Mandatory", field.Name))
		}
	} else {
		return errors.New(fmt.Sprintf("%v is Mandatory", field.Name))
	}

	return nil
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
