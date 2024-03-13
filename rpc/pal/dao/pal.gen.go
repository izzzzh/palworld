// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"palworld/rpc/pal/model"
)

func newPal(db *gorm.DB, opts ...gen.DOOption) pal {
	_pal := pal{}

	_pal.palDo.UseDB(db, opts...)
	_pal.palDo.UseModel(&model.Pal{})

	tableName := _pal.palDo.TableName()
	_pal.ALL = field.NewAsterisk(tableName)
	_pal.ID = field.NewInt64(tableName, "id")
	_pal.Number = field.NewString(tableName, "number")
	_pal.Name = field.NewString(tableName, "name")
	_pal.EnName = field.NewString(tableName, "en_name")
	_pal.Icon = field.NewString(tableName, "icon")
	_pal.AttributeIds = field.NewString(tableName, "attribute_ids")
	_pal.Hp = field.NewInt32(tableName, "hp")
	_pal.Energy = field.NewInt32(tableName, "energy")
	_pal.Defensively = field.NewInt32(tableName, "defensively")
	_pal.Abilities = field.NewString(tableName, "abilities")
	_pal.Eat = field.NewInt32(tableName, "eat")
	_pal.PassiveSkill = field.NewInt32(tableName, "passive_skill")
	_pal.ActiveSkills = field.NewString(tableName, "active_skills")
	_pal.Level = field.NewInt32(tableName, "level")

	_pal.fillFieldMap()

	return _pal
}

// pal 帕鲁
type pal struct {
	palDo

	ALL          field.Asterisk
	ID           field.Int64
	Number       field.String // 编号
	Name         field.String // 名字
	EnName       field.String // 英文名字
	Icon         field.String // 头像
	AttributeIds field.String // 属性id
	Hp           field.Int32  // 血量
	Energy       field.Int32  // 攻击力
	Defensively  field.Int32  // 防御力
	Abilities    field.String // 工作能力
	Eat          field.Int32  // 消耗
	PassiveSkill field.Int32  // 被动技能
	ActiveSkills field.String // 主动技能
	Level        field.Int32  // 级别（孵蛋大小）

	fieldMap map[string]field.Expr
}

func (p pal) Table(newTableName string) *pal {
	p.palDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pal) As(alias string) *pal {
	p.palDo.DO = *(p.palDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pal) updateTableName(table string) *pal {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.Number = field.NewString(table, "number")
	p.Name = field.NewString(table, "name")
	p.EnName = field.NewString(table, "en_name")
	p.Icon = field.NewString(table, "icon")
	p.AttributeIds = field.NewString(table, "attribute_ids")
	p.Hp = field.NewInt32(table, "hp")
	p.Energy = field.NewInt32(table, "energy")
	p.Defensively = field.NewInt32(table, "defensively")
	p.Abilities = field.NewString(table, "abilities")
	p.Eat = field.NewInt32(table, "eat")
	p.PassiveSkill = field.NewInt32(table, "passive_skill")
	p.ActiveSkills = field.NewString(table, "active_skills")
	p.Level = field.NewInt32(table, "level")

	p.fillFieldMap()

	return p
}

func (p *pal) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pal) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 14)
	p.fieldMap["id"] = p.ID
	p.fieldMap["number"] = p.Number
	p.fieldMap["name"] = p.Name
	p.fieldMap["en_name"] = p.EnName
	p.fieldMap["icon"] = p.Icon
	p.fieldMap["attribute_ids"] = p.AttributeIds
	p.fieldMap["hp"] = p.Hp
	p.fieldMap["energy"] = p.Energy
	p.fieldMap["defensively"] = p.Defensively
	p.fieldMap["abilities"] = p.Abilities
	p.fieldMap["eat"] = p.Eat
	p.fieldMap["passive_skill"] = p.PassiveSkill
	p.fieldMap["active_skills"] = p.ActiveSkills
	p.fieldMap["level"] = p.Level
}

func (p pal) clone(db *gorm.DB) pal {
	p.palDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pal) replaceDB(db *gorm.DB) pal {
	p.palDo.ReplaceDB(db)
	return p
}

type palDo struct{ gen.DO }

type IPalDo interface {
	gen.SubQuery
	Debug() IPalDo
	WithContext(ctx context.Context) IPalDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPalDo
	WriteDB() IPalDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPalDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPalDo
	Not(conds ...gen.Condition) IPalDo
	Or(conds ...gen.Condition) IPalDo
	Select(conds ...field.Expr) IPalDo
	Where(conds ...gen.Condition) IPalDo
	Order(conds ...field.Expr) IPalDo
	Distinct(cols ...field.Expr) IPalDo
	Omit(cols ...field.Expr) IPalDo
	Join(table schema.Tabler, on ...field.Expr) IPalDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPalDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPalDo
	Group(cols ...field.Expr) IPalDo
	Having(conds ...gen.Condition) IPalDo
	Limit(limit int) IPalDo
	Offset(offset int) IPalDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPalDo
	Unscoped() IPalDo
	Create(values ...*model.Pal) error
	CreateInBatches(values []*model.Pal, batchSize int) error
	Save(values ...*model.Pal) error
	First() (*model.Pal, error)
	Take() (*model.Pal, error)
	Last() (*model.Pal, error)
	Find() ([]*model.Pal, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Pal, err error)
	FindInBatches(result *[]*model.Pal, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Pal) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPalDo
	Assign(attrs ...field.AssignExpr) IPalDo
	Joins(fields ...field.RelationField) IPalDo
	Preload(fields ...field.RelationField) IPalDo
	FirstOrInit() (*model.Pal, error)
	FirstOrCreate() (*model.Pal, error)
	FindByPage(offset int, limit int) (result []*model.Pal, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPalDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p palDo) Debug() IPalDo {
	return p.withDO(p.DO.Debug())
}

func (p palDo) WithContext(ctx context.Context) IPalDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p palDo) ReadDB() IPalDo {
	return p.Clauses(dbresolver.Read)
}

func (p palDo) WriteDB() IPalDo {
	return p.Clauses(dbresolver.Write)
}

func (p palDo) Session(config *gorm.Session) IPalDo {
	return p.withDO(p.DO.Session(config))
}

func (p palDo) Clauses(conds ...clause.Expression) IPalDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p palDo) Returning(value interface{}, columns ...string) IPalDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p palDo) Not(conds ...gen.Condition) IPalDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p palDo) Or(conds ...gen.Condition) IPalDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p palDo) Select(conds ...field.Expr) IPalDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p palDo) Where(conds ...gen.Condition) IPalDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p palDo) Order(conds ...field.Expr) IPalDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p palDo) Distinct(cols ...field.Expr) IPalDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p palDo) Omit(cols ...field.Expr) IPalDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p palDo) Join(table schema.Tabler, on ...field.Expr) IPalDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p palDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPalDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p palDo) RightJoin(table schema.Tabler, on ...field.Expr) IPalDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p palDo) Group(cols ...field.Expr) IPalDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p palDo) Having(conds ...gen.Condition) IPalDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p palDo) Limit(limit int) IPalDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p palDo) Offset(offset int) IPalDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p palDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPalDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p palDo) Unscoped() IPalDo {
	return p.withDO(p.DO.Unscoped())
}

func (p palDo) Create(values ...*model.Pal) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p palDo) CreateInBatches(values []*model.Pal, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p palDo) Save(values ...*model.Pal) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p palDo) First() (*model.Pal, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Pal), nil
	}
}

func (p palDo) Take() (*model.Pal, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Pal), nil
	}
}

func (p palDo) Last() (*model.Pal, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Pal), nil
	}
}

func (p palDo) Find() ([]*model.Pal, error) {
	result, err := p.DO.Find()
	return result.([]*model.Pal), err
}

func (p palDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Pal, err error) {
	buf := make([]*model.Pal, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p palDo) FindInBatches(result *[]*model.Pal, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p palDo) Attrs(attrs ...field.AssignExpr) IPalDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p palDo) Assign(attrs ...field.AssignExpr) IPalDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p palDo) Joins(fields ...field.RelationField) IPalDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p palDo) Preload(fields ...field.RelationField) IPalDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p palDo) FirstOrInit() (*model.Pal, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Pal), nil
	}
}

func (p palDo) FirstOrCreate() (*model.Pal, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Pal), nil
	}
}

func (p palDo) FindByPage(offset int, limit int) (result []*model.Pal, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p palDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p palDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p palDo) Delete(models ...*model.Pal) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *palDo) withDO(do gen.Dao) *palDo {
	p.DO = *do.(*gen.DO)
	return p
}
