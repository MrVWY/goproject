package model
//Customize errors is based on business logic
import (
	"errors"
)

var (
	//Login
	ERROR_USER_NOTEXISTS = errors.New("user dose not exsist")
	ERROR_USER_EXISTS = errors.New("user already exsist")
	ERROR_USER_PWD = errors.New("user's PWD is error ")
	//Register
	ERROR_USER_RegisterIsFail = errors.New("The Register is fail")
	ERROR_USER_RegisterUserIDIsExsist = errors.New("The Register's UserID is exsist")
)