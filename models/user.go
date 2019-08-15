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
	Passwd   string `json:"passwd" xorm:"-"` // 引入一个冗余字段，不写入数据库；目的是进行Json bind 使用
	Password string `json:"-" xorm:"notnull"`
	Salt     string `json:"-" xorm:"notnull"`
	Email    string `json:"email" xorm:"notnull"`
	Phone    string `json:"phone" xorm:"notnull"`
	IsSuper  bool   `json:"-" xorm:"notnull <-"` // 只从数据库读，不写
}

// TableName return struct database table name
func (User) TableName() string {
	return "users.users"
}

// Insert create a new record and insert into database
func (user *User) Insert() (int64, error) {
	return x.Insert(user)
}

// Exist 判断用户是否存在
func (user *User) Exist() (bool, error) {
	return x.Exist(user)
}

// BeforeInsert function
func (user *User) BeforeInsert() {
	user.Model.BeforeInsert()
	user.Salt = generatePasswordSalt()
	user.Password = calculatePassword(user.Passwd, user.Salt)
}

// IsPasswordRight check password is right
func (user *User) IsPasswordRight(passwd string) bool {
	return user.Password == calculatePassword(passwd, user.Salt)
}

// ResetPassword 重置密码
func (user *User) ResetPassword(password string) (int64, error) {
	user.Salt = generatePasswordSalt()
	user.Password = calculatePassword(password, user.Salt)
	return x.Id(user.ID).Update(user)
}

// GetAllUsers return set of User
func GetAllUsers() (users []User, err error) {
	err = x.Find(&users)
	if nil != err {
		return nil, err
	}

	return users, nil
}

// GetUserByID 通过user_id获取User记录
func GetUserByID(id string) (*User, error) {
	user := new(User)
	has, err := x.Id(id).Get(user)
	if nil != err {
		return nil, err
	}
	if !has {
		return nil, ErrUserNotExist{}
	}

	return user, nil
}

// GetUserByUsername 按照用户名获取用户信息
func GetUserByUsername(username string) (*User, error) {
	user := new(User)
	has, err := x.Where("username = ?", username).Get(user)
	if nil != err {
		return nil, err
	}
	if !has {
		return nil, ErrUserNotExist{}
	}

	return user, nil
}

// generateSalt 生成加密盐
func generatePasswordSalt() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(int(rand.Int31()))
}

// calculatePassword 根据给定原始密码以及加密盐生成密码
func calculatePassword(passwd, salt string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(passwd+salt)))
}
