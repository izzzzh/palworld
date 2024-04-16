// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameSkill = "skill"

// Skill 技能
type Skill struct {
	ID          int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name        string         `gorm:"column:name;not null;comment:名称" json:"name"`                                         // 名称
	Description string         `gorm:"column:description;not null;comment:描述" json:"description"`                           // 描述
	Ct          int32          `gorm:"column:ct;not null;comment:冷却" json:"ct"`                                             // 冷却
	Power       int32          `gorm:"column:power;not null;comment:威力" json:"power"`                                       // 威力
	AttributeID int32          `gorm:"column:attribute_id;not null;comment:属性ID" json:"attribute_id"`                       // 属性ID
	CreatedAt   time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt   time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"` // 更新时间
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`                                    // 删除时间
}

// TableName Skill's table name
func (*Skill) TableName() string {
	return TableNameSkill
}
