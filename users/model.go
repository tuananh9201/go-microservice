package main

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Role       string `json:"role"`
	DeleteFlag bool   `json:"delete_flag"`
}

type UserFilter struct {
	Email      string `json:"email"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	DeleteFlag bool   `json:"delete_flag"`
}

type Claims struct {
	Role string `json:"role"`
	jwt.StandardClaims
}
