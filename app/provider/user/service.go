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
	if db.Where(&User{UserName: user.UserName}).First(userDB).Error != gorm.ErrRecordNotFound {
		return nil, errors.New("用户名已经被注册，请换一个用户名")
	}

	token := genToken(10)
	user.Token = token

	// 将请求注册进入redis，保存一天
	cacheService := u.container.MustMake(contract.CacheKey).(contract.CacheService)

	key := fmt.Sprintf("user:register:%v", user.Token)
	if err := cacheService.SetObj(ctx, key, user, 24*time.Hour); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserService) SendRegisterMail(ctx context.Context, user *User) error {
	logger := u.container.MustMake(contract.LogKey).(contract.Log)
	configer := u.container.MustMake(contract.ConfigKey).(contract.Config)

	// 配置服务中获取发送邮件需要的参数
	host := configer.GetString("app.smtp.host")
	port := configer.GetInt("app.smtp.port")
	username := configer.GetString("app.smtp.username")
	password := configer.GetString("app.smtp.password")
	from := configer.GetString("app.smtp.from")
	domain := configer.GetString("app.domain")

	// 实例化gomail
	d := gomail.NewDialer(host, port, username, password)

	// 组装message
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetAddressHeader("To", user.Email, user.UserName)
	m.SetHeader("Subject", "感谢您注册我们的hadecast")
	link := fmt.Sprintf("%v/user/register/verify?token=%v", domain, user.Token)
	m.SetBody("text/html", fmt.Sprintf("请点击下面的链接完成注册：%s", link))

	// 发送电子邮件
	if err := d.DialAndSend(m); err != nil {
		logger.Error(ctx, "send email error", map[string]interface{}{
			"err":     err,
			"message": m,
		})
		return err
	}
	return nil
}

func (u *UserService) VerifyRegister(ctx context.Context, token string) (bool, error) {
	//验证token
	cacheService := u.container.MustMake(contract.CacheKey).(contract.CacheService)
	key := fmt.Sprintf("user:register:%v", token)
	user := &User{}
	if err := cacheService.GetObj(ctx, key, user); err != nil {
		return false, err
	}
	if user.Token != token {
		return false, nil
	}

	//验证邮箱，用户名的唯一
	ormService := u.container.MustMake(contract.ORMKey).(contract.ORMService)
	db, err := ormService.GetDB()
	if err != nil {
		return false, err
	}
	userDB := &User{}
	if db.Where(&User{Email: user.Email}).First(userDB).Error != gorm.ErrRecordNotFound {
		return false, errors.New("邮箱已注册用户，不能重复注册")
	}
	if db.Where(&User{UserName: user.UserName}).First(userDB).Error != gorm.ErrRecordNotFound {
		return false, errors.New("用户名已经被注册，请换一个用户名")
	}

	// 验证成功将密码存储数据库之前需要加密，不能原文存储进入数据库
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return false, err
	}
	user.Password = string(hash)

	// 具体在数据库创建用户
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
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
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
