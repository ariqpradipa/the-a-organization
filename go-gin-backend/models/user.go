package models

import (
	"errors"
	"fmt"
	"html"
	"strings"

	"bookweb/utils/token"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	// Id uint `json:"id" gorm:"primary_key"`
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;unique" json:"password"`
}

func (user *User) SaveUser() (*User, error) {
	var err error
	err = DB.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave() error {
	// hashing password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashPassword)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}

// Verify password correct or not
func VerifyPassword(password, hasedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hasedPassword), []byte(password))
}

// Check user exist or not and return token
func LoginCheck(username string, password string) (string, error) {
	var err error
	user := User{}
	err = DB.Model(User{}).Where("username = ?", username).Take(&user).Error
	if err != nil {
		return "", err
	}
	// typed password compare with user's database password
	err = VerifyPassword(password, user.Password)
	
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	token, err := token.GenerateToken(user.ID)
	fmt.Println("verify token",token,err)
	if err != nil {
		return "", err
	}
	return token, nil
}

// Get user by ID

func GetUserByID(uid uint)(User,error){
	var user User
	if err:= DB.First(&user,uid).Error;err!=nil{
		return user,errors.New("User not found")
	}
	user.PrepareGive()
	return user,nil
}

func (user *User) PrepareGive(){
	user.Password = ""
}


