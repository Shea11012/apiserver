package model

import (
	"apiserver/pkg/auth"
	"apiserver/pkg/constvar"
	"fmt"
	"github.com/go-playground/validator/v10"
	"sync"
)

type UserInfo struct {
	Id        uint64 `json:"id"`
	Username  string `json:"username"`
	SayHello  string `json:"sayHello"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type UserList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*UserInfo
}

type User struct {
	BaseModel
	Username string `json:"username" gorm:"column:username;not null" validate:"required,min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" validate:"required,min=5,max=128"`
	Cards []Card
}

func (u *User) Create() error {
	return DB.mysql.Create(&u).Error
}

func DeleteUser(id uint64) error {
	user := User{}
	user.Id = id
	return DB.mysql.Delete(&user).Error
}

func (u *User) Update() error {
	return DB.mysql.Save(u).Error
}

func GetUser(username string) (*User, error) {
	u := &User{}
	d := DB.mysql.Where("username = ?", username).First(&u)
	return u, d.Error
}

func GetUserById(id uint64) (User,error) {
	u := User{}
	d := DB.mysql.Where("id = ?",id).First(&u)
	return u,d.Error
}

func ListUser(username string, offset, limit int) ([]*User, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	users := make([]*User, 0)
	var count uint64
	where := fmt.Sprintf("username like '%%%s%%'", username)
	if err := DB.mysql.Model(&User{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}

	if err := DB.mysql.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}

	return users, count, nil
}

func (u *User) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)
	return
}

func (u *User) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

// 测试gorm关联模型
func (u *User) GetCards() ([]Card,error) {
	cards := make([]Card,0)
	DB.mysql.Model(u).Related(&cards)
	return cards,nil
}
