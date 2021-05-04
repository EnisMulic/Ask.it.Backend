package constants

import "errors"

const (
	ErrMsgUnableToParseJSONBody string = "Unable to parse JSON body"
	ErrMsgUnableToConvertId string = "Unable to convert id"
	ErrMsgUnableToConvertUserId string = "Unable to convert user id"
	ErrMsgUnableToMarshalJson string = "Unable to marshal json"
	ErrMsgUnableToParseQueryParametars string = "Unable to parse query parametars"
)

var (
	ErrGeneric = errors.New("an error occurred")
	ErrEmailIsTaken = errors.New("the email is taken")
	ErrWrongPassword = errors.New("wrong password")
	ErrUserNotFound = errors.New("user not found")
	ErrQuestionNotFound = errors.New("question not found")
	ErrAnswerNotFound = errors.New("answer not found")
	ErrUnauthorized = errors.New("unauthorized")
)