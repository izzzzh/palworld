// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNamePalMateMap = "pal_mate_map"

// PalMateMap 帕鲁配对表
type PalMateMap struct {
	ID        int64 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ParentOne int32 `gorm:"column:parent_one;not null;comment:父母一号" json:"parent_one"` // 父母一号
	ParentTwo int32 `gorm:"column:parent_two;not null;comment:父母二号" json:"parent_two"` // 父母二号
	Result    int32 `gorm:"column:result;not null;comment:后代" json:"result"`           // 后代
}

// TableName PalMateMap's table name
func (*PalMateMap) TableName() string {
	return TableNamePalMateMap
}
