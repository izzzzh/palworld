// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUser = "user"

// User 用户
type User struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string    `gorm:"column:name;not null;comment:姓名" json:"name"`                                         // 姓名
	Password  string    `gorm:"column:password;not null;comment:密码" json:"password"`                                 // 密码
	Role      int32     `gorm:"column:role;not null;default:1;comment:角色" json:"role"`                               // 角色
	Avatar    string    `gorm:"column:avatar;not null;comment:头像" json:"avatar"`                                     // 头像
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"` // 创建时间
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}