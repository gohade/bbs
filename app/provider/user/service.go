package user

import (
	"context"
	"fmt"
	"github.com/gohade/hade/framework"
	"github.com/gohade/hade/framework/contract"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

type UserService struct {
	container framework.Container
	logger    contract.Log
	configer  contract.Config
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func genToken(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (u *UserService) Register(ctx context.Context, user *User) (*User, error) {
	// 判断邮箱是否已经注册了
	ormService := u.container.MustMake(contract.ORMKey).(contract.ORMService)
	db, err := ormService.GetDB()
	if err != nil {
		return nil, err
	}
	userDB := &User{}
	if db.Where(&User{Email: user.Email}).First(userDB).Error != gorm.ErrRecordNotFound {
		return nil, errors.New("邮箱已注册用户，不能重复注册")
	}

	// 将请求注册进入redis，保存一天
	cacheService := u.container.MustMake(contract.CacheKey).(contract.CacheService)
	token := genToken(10)
	user.Token = token

	key := fmt.Sprintf("user:register:%v", user.Token)
	if err := cacheService.SetObj(ctx, key, user, 24*time.Hour); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserService) SendRegisterMail(ctx context.Context, user *User) error {
	host := u.configer.GetString("app.smtp.host")
	port := u.configer.GetInt("app.smtp.port")
	username := u.configer.GetString("app.smtp.username")
	password := u.configer.GetString("app.smtp.password")
	from := u.configer.GetString("app.smtp.from")
	d := gomail.NewDialer(host, port, username, password)
	//d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetAddressHeader("To", user.Email, user.UserName)
	m.SetHeader("Subject", "感谢您注册我们的hadecast")
	domain := u.configer.GetString("app.domain")
	link := fmt.Sprintf("%v/user/register/verify?token=%v", domain, user.Token)
	m.SetBody("text/html", fmt.Sprintf("请点击下面的链接完成注册：%s", link))

	// 发送电子邮件
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func (u *UserService) VerifyRegister(ctx context.Context, token string) (bool, error) {
	cacheService := u.container.MustMake(contract.CacheKey).(contract.CacheService)

	key := fmt.Sprintf("user:register:%v", token)
	user := &User{}
	if err := cacheService.GetObj(ctx, key, user); err != nil {
		return false, err
	}
	if user.Token != token {
		return false, nil
	}

	// 判断邮箱是否已经注册了
	ormService := u.container.MustMake(contract.ORMKey).(contract.ORMService)
	db, err := ormService.GetDB()
	if err != nil {
		return false, err
	}
	userDB := &User{}
	if db.Where(&User{Email: user.Email}).First(userDB).Error != gorm.ErrRecordNotFound {
		return false, errors.New("邮箱已注册用户，不能重复注册")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return false, err
	}
	user.Password = string(hash)

	if err := db.Create(user).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (u *UserService) Login(ctx context.Context, user *User) (*User, error) {
	ormService := u.container.MustMake(contract.ORMKey).(contract.ORMService)
	db, err := ormService.GetDB()
	if err != nil {
		return nil, err
	}

	userDB := &User{}
	if err := db.Where("username=?", user.UserName).First(userDB).Error; err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password)); err != nil {
		return nil, err
	}

	userDB.Password = ""
	// 缓存session
	cacheService := u.container.MustMake(contract.CacheKey).(contract.CacheService)
	token := genToken(10)
	key := fmt.Sprintf("session:%v", token)
	if err := cacheService.SetObj(ctx, key, userDB, 24*time.Hour); err != nil {
		return nil, err
	}

	userDB.Token = token
	return userDB, nil
}

func (u *UserService) Logout(ctx context.Context, user *User) error {
	cacheService := u.container.MustMake(contract.CacheKey).(contract.CacheService)
	userSession, err := u.VerifyLogin(ctx, user.Token)
	// 不需要做任何操作
	if err != nil || userSession.UserName != user.UserName {
		return nil
	}

	key := fmt.Sprintf("session:%v", user.Token)
	_ = cacheService.Del(ctx, key)
	return nil
}

func (u *UserService) VerifyLogin(ctx context.Context, token string) (*User, error) {
	cacheService := u.container.MustMake(contract.CacheKey).(contract.CacheService)
	key := fmt.Sprintf("session:%v", token)
	user := &User{}
	if err := cacheService.GetObj(ctx, key, user); err != nil {
		return nil, err
	}

	return user, nil
}

func NewUserService(params ...interface{}) (interface{}, error) {
	container := params[0].(framework.Container)
	logger := container.MustMake(contract.LogKey).(contract.Log)
	configer := container.MustMake(contract.ConfigKey).(contract.Config)
	return &UserService{container: container, logger: logger, configer: configer}, nil
}
