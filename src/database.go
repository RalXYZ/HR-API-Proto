package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"os"
)

type Login struct {
	DbUser   string `json:"dbUser"`
	DbPasswd string `json:"dbPasswd"`
}

var db *gorm.DB

func getLoginInfo(file string) (string, string) {
	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Error: Requires a database configuration file!")
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var login Login
	json.Unmarshal([]byte(byteValue), &login)

	return login.DbUser, login.DbPasswd
}

func connectDatabase(dbName string)  {
	configDirectory := "conf/admin.json"
	username, password := getLoginInfo(configDirectory)
	dbArgs := username + ":" + password + "@(localhost)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open("mysql", dbArgs)
	if err != nil {
		panic(err)
	}
}

type User struct {
	Username string `gorm:"size:32"`
	Password []byte `gorm:"size:60"`
}

type Qscer struct {
	UID    uint  `gorm:"primary_key;auto_increment"`
	Zjuid string `gorm:"size:10;not_null"`
	Name  string `gorm:"size:10;not_null"`
	Qscid string `gorm:"size:10;not_null"`
	Birthday string `gorm:"size:10;not_null"`
}

func initDatabase() {
	connectDatabase("hr_proto")
	// db.LogMode(true)
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Qscer{})
}

func addUser(requestUser *User) error {
	if requestUser.Username == "" || requestUser.Password == nil {
		panic("request Username or Password not exists")
	}
	if db.First(&User{}, "username = ?", requestUser.Username).RecordNotFound() {
		hashedPassword, err := bcrypt.GenerateFromPassword(requestUser.Password, bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		requestUser.Password = hashedPassword
		db.Create(requestUser)
		return nil
	} else {
		return errors.New("user exists")
	}
}

func checkLoginInfo(requestUser *User) error {
	var dbUser User
	if result := db.First(&dbUser, "username = ?", requestUser.Username); result.RecordNotFound() {
		return errors.New("user not found")
	} else if bcrypt.CompareHashAndPassword(dbUser.Password, requestUser.Password) == bcrypt.ErrMismatchedHashAndPassword {
		return errors.New("wrong Password")
	} else {
		return nil
	}
}

func addQSCer(requestQSCer *Qscer) error {
	if db.First(&Qscer{}, "zjuid = ?", requestQSCer.Zjuid).RecordNotFound() {
		db.Create(requestQSCer)
		return nil
	} else {
		return errors.New("ZJU ID duplicated")
	}
}
