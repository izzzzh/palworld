// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameGoods = "goods"

// Goods 物品
type Goods struct {
	ID          int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name        string `gorm:"column:name;not null;comment:名称" json:"name"`      // 名称
	Description string `gorm:"column:description;comment:描述" json:"description"` // 描述
	Image       string `gorm:"column:image;comment:图片" json:"image"`             // 图片
	Materials   string `gorm:"column:materials;comment:制造材料" json:"materials"`   // 制造材料
	Quality     int32  `gorm:"column:quality;comment:品质" json:"quality"`         // 品质
	Workload    int32  `gorm:"column:workload;comment:工作量" json:"workload"`      // 工作量
	Types       string `gorm:"column:types;comment:类型" json:"types"`             // 类型
}

// TableName Goods's table name
func (*Goods) TableName() string {
	return TableNameGoods
}
