package models

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// User is struct
type User struct {
	Model    `xorm:"extends"`
	Username string `json:"username" xorm:"notnull"`
	Passwd   string `json:"passwd" xorm:"-"`
	Password string `json:"-" xorm:"notnull"`
	Salt     string `json:"-" xorm:"notnull"`
	Email    string `json:"email" xorm:"notnull"`
	Phone    string `json:"phone" xorm:"notnull"`
}

// TableName return struct database table name
func (User) TableName() string {
	return "users.users"
}

// Insert create a new record and insert into database
func (user *User) Insert() (affected int64, err error) {
	return x.Insert(user)
}

// BeforeInsert function
func (user *User) BeforeInsert() {
	user.Model.BeforeInsert()

	rand.Seed(time.Now().UnixNano())
	user.Salt = strconv.Itoa(int(rand.Int31()))

	fmt.Printf("Password: %s, salt: %s\n", user.Passwd, user.Salt)
	user.Password = fmt.Sprintf("%x", md5.Sum([]byte(user.Passwd+user.Salt)))
}

// GetAllUsers return set of User
func GetAllUsers() (users []User, err error) {
	err = x.Find(&users)
	if nil != err {
		return nil, err
	}

	return users, nil
}
