package model

type Flash struct {
	Type           string
	Message        string
	FormInputValue map[string]string
	FormErrorMap   map[string]string
}
