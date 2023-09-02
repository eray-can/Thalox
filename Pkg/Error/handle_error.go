package Error

import (
	"fmt"
	"time"
)

type CustomError struct {
	DeveloperMessage error
	UserMessage      string
	Timestamp        string
	Code             int
	UserInput        string
}

var Codes = map[int]string{
	0: "Sistem hatası",
	1: "Gerçek bir kullanıcı değilsiniz",
	2: "Komut hatası",
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Üzgünüz, [%s] için işlem yapılamıyor: %s", e.UserInput, e.UserMessage)
}

func NewError(code int, userInput string) *CustomError {
	return &CustomError{
		DeveloperMessage: DeveloperMessage,
		UserMessage:      Codes[code],
		Timestamp:        time.Now().Format(time.RFC3339),
		Code:             code,
		UserInput:        userInput,
	}
}

func UpdateError(code int, err *CustomError) *CustomError {
	err.Code = code
	err.UserMessage = Codes[code]
	err.Timestamp = time.Now().Format(time.RFC3339)
	return err
}
