package src

import (
	"encoding/base64"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDb() {
	// Create the database tables
	var err error
	db, err = gorm.Open(postgres.Open(Envars.DbUrl), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Accounts{})
	if err != nil {
		panic(err)
	}
	log.Println("Database connected")
}

func AddAccount(userID int64, ownerId int64, appid int32, apphash, stringSession string) error {
	session, err := encryptAES(stringSession)
	if err != nil {
		fmt.Println(err)
		return err
	}
	account := Accounts{
		UserID:        userID,
		OwnerId:       ownerId,
		StringSession: base64.StdEncoding.EncodeToString(session),
		AppId:         appid,
		AppHash:       apphash,
	}
	err = db.Create(&account).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func FetchAccounts(userID int64) ([]Accounts, error) {
	var accounts []Accounts
	err := db.Where("owner_id = ?", userID).Find(&accounts).Error
	if err != nil {
		return nil, err
	}
	for i := range accounts {
		session, err := base64.StdEncoding.DecodeString(accounts[i].StringSession)
		if err != nil {
			return nil, err
		}
		accounts[i].StringSession, err = decryptAES(session)
		if err != nil {
			return nil, err
		}
	}
	return accounts, nil
}
