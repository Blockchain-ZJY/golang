package main

import "fmt"

type UserConfig struct {
	Name     string
	Password string
	Age      int
	Email    string
}

type option func(*UserConfig)

func WithName(name string) option {
	return func(uc *UserConfig) {
		uc.Name = name
	}
}
func WithPassword(Password string) option {
	return func(uc *UserConfig) {
		uc.Password = Password
	}
}
func WithEmail(Email string) option {
	return func(uc *UserConfig) {
		uc.Email = Email
	}
}
func WithAge(age int) option {
	return func(uc *UserConfig) {
		uc.Age = age
	}
}

func NewUserConfig(opts ...option) *UserConfig {
	s := &UserConfig{
		Name:     "default",
		Password: "default",
		Age:      0,
		Email:    "default",
	}

	for _, opt := range opts {
		opt(s)
	}
	return s
}

func main() {
	userconfig := NewUserConfig(WithName("xiaoming"), WithPassword("<PASSWORD>"), WithEmail("<EMAIL>"), WithAge(18))
	fmt.Println(userconfig)
}
