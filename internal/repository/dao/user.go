package dao

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	ErrDuplicateEmail = errors.New("邮箱已经被注册")
	ErrRecordNotFound = gorm.ErrRecordNotFound
)

type User struct {
	Id            int64          `gorm:"primaryKey,autoIncrement"`
	Email         sql.NullString `gorm:"unique"`
	Password      string
	Nickname      string `gorm:"type=varchar(128)"`
	Birthday      int64
	AboutMe       string         `gorm:"type=varchar(4096)"`
	Phone         sql.NullString `gorm:"unique"`
	WechatOpenId  sql.NullString `gorm:"unique"`
	WechatUnionId sql.NullString
	Ctime         int64
	Utime         int64
}
type UserDao interface {
	Insert(ctx context.Context, u User) error
	FindByEmail(ctx context.Context, email string) (User, error)
	UpdateById(ctx context.Context, entity User) error
	FindById(ctx context.Context, uid int64) (User, error)
}
type GormUserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) UserDao {
	return &GormUserDao{
		db: db,
	}
}
func (d *GormUserDao) Insert(ctx context.Context, u User) error {
	now := time.Now().UnixMilli()
	u.Ctime = now
	u.Utime = now
	err := d.db.WithContext(ctx).Create(&u).Error
	var me *mysql.MySQLError
	if errors.As(err, &me) {
		const duplicateErr = 1062
		if me.Number == duplicateErr {
			//邮箱已经被注册
			return ErrDuplicateEmail
		}
	}
	return err
}
func (d *GormUserDao) FindById(ctx context.Context, uid int64) (User, error) {
	var u User
	err := d.db.WithContext(ctx).Where("id = ?", uid).First(&u).Error
	return u, err
}

func (d *GormUserDao) UpdateById(ctx context.Context, entity User) error {
	return d.db.WithContext(ctx).Model(&entity).Where("id = ?", entity.Id).Updates(map[string]any{
		"utime":    time.Now().UnixMilli(),
		"nickname": entity.Nickname,
		"birthday": entity.Birthday,
		"about_me": entity.AboutMe,
	}).Error
}

func (d *GormUserDao) FindByEmail(ctx context.Context, email string) (User, error) {
	var u User
	err := d.db.WithContext(ctx).Where("email = ?", email).First(&u).Error
	return u, err
}
