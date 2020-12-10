package database

import (
	"HR-API-proto/database/models"
	"HR-API-proto/utils"
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/crypto/bcrypt"
)

var DB *gorm.DB

func ConnectDatabase(dbName string) {
	configDirectory := "conf/admin.json"
	username, password := utils.GetLoginInfo(configDirectory)
	dbArgs := username + ":" + password + "@(localhost)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open("mysql", dbArgs)
	if err != nil {
		panic(err)
	}
}

func Init() {
	ConnectDatabase("hr_proto")
	// DB.LogMode(true)
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Qscer{})
}

func AddUser(requestUser *models.User) error {
	if requestUser.Username == "" || requestUser.Password == nil {
		panic("request Username or Password not exists")
	}
	if DB.First(&models.User{}, "username = ?", requestUser.Username).RecordNotFound() {
		hashedPassword, err := bcrypt.GenerateFromPassword(requestUser.Password, bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		requestUser.Password = hashedPassword
		DB.Create(requestUser)
		return nil
	} else {
		return errors.New("user exists")
	}
}

func VerifyLoginCredential(requestUser *models.User) error {
	var dbUser models.User
	if result := DB.First(&dbUser, "username = ?", requestUser.Username); result.RecordNotFound() {
		return errors.New("user not found")
	} else if bcrypt.CompareHashAndPassword(dbUser.Password, requestUser.Password) == bcrypt.ErrMismatchedHashAndPassword {
		return errors.New("wrong Password")
	} else {
		return nil
	}
}

func AddQSCer(requestQSCer *models.Qscer) error {
	if DB.First(&models.Qscer{}, "zjuid = ?", requestQSCer.Zjuid).RecordNotFound() {
		DB.Create(requestQSCer)
		return nil
	} else {
		return errors.New("ZJU ID duplicated")
	}
}
