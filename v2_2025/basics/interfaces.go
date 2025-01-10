package main

import (
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(username, password string) (string, error)
	Signup(username, password string) (string, error)
}

//diff version of auth with simple and jwt

type SimpleAuthService struct{}

// methods with pointer receivers

func (s *SimpleAuthService) Login(username, password string) (string, error) {
	buffer := bytes.Buffer{}

	buffer.WriteString(password)
	buffer.WriteString(username)

	return buffer.String(), nil
}

func (s *SimpleAuthService) Signup(username, password string) (string, error) {
	buffer := bytes.Buffer{}

	buffer.WriteString(username)
	buffer.WriteString(password)

	return buffer.String(), nil
}

func NewSimpleAuthService() AuthService {
	return &SimpleAuthService{}
}

type JWTAuthService struct{}

func (s *JWTAuthService) Login(username, password string) (string, error) {
	return "", nil
}
func (s *JWTAuthService) Signup(username, password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	//generate JWT token

	return string(hashedPassword), nil
}

func NewJWTAuthService() AuthService {
	return &JWTAuthService{}
}

//user controller receives request

type UserController struct {
	authService AuthService
}

func (c *UserController) SignUpHandler(name, pass string) (string, error) {
	//validate data
	if name == "" && pass == "" {
		return "", errors.New("invalid user data")
	}

	result, err := c.authService.Signup(name, pass)
	if err != nil {
		return "", errors.New("something went wrong")
	}

	return result, nil
}

func NewUserController(authService AuthService) *UserController {
	return &UserController{authService: authService}
}

func main() {
	//received data somewhere
	userName := "max"
	password := "test123"

	authService := NewSimpleAuthService()

	//eventually we might need to change auth service, but we already have realisation below and it might be huge.
	//exactly for this reason I used AuthService Interface. I only have to make new JWT Auth Service and proceed it further.

	//The day has come and new AuthService occurs.
	//authService = NewJWTAuthService()

	userController := NewUserController(authService)

	result, err := userController.SignUpHandler(userName, password)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

//The showcase how struct embedding lead to interface inheritance
//Child class implements ChildInterface and by embedding it Parent does it too.

type ParentStruct struct {
	ChildStruct
}

type ChildStruct struct{}

func (c ChildStruct) testFunc() {
	fmt.Println("testFunc")
}

type ChildInterface interface {
	testFunc()
}
