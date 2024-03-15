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

	"palworld/rpc/skill/model"
)

func newSkill(db *gorm.DB, opts ...gen.DOOption) skill {
	_skill := skill{}

	_skill.skillDo.UseDB(db, opts...)
	_skill.skillDo.UseModel(&model.Skill{})

	tableName := _skill.skillDo.TableName()
	_skill.ALL = field.NewAsterisk(tableName)
	_skill.ID = field.NewInt64(tableName, "id")
	_skill.Name = field.NewString(tableName, "name")
	_skill.Description = field.NewString(tableName, "description")
	_skill.Ct = field.NewInt32(tableName, "ct")
	_skill.Power = field.NewInt32(tableName, "power")
	_skill.AttributeID = field.NewInt32(tableName, "attribute_id")

	_skill.fillFieldMap()

	return _skill
}

// skill 技能
type skill struct {
	skillDo

	ALL         field.Asterisk
	ID          field.Int64
	Name        field.String // 名称
	Description field.String // 描述
	Ct          field.Int32  // 冷却
	Power       field.Int32  // 威力
	AttributeID field.Int32  // 属性ID

	fieldMap map[string]field.Expr
}

func (s skill) Table(newTableName string) *skill {
	s.skillDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s skill) As(alias string) *skill {
	s.skillDo.DO = *(s.skillDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *skill) updateTableName(table string) *skill {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.Name = field.NewString(table, "name")
	s.Description = field.NewString(table, "description")
	s.Ct = field.NewInt32(table, "ct")
	s.Power = field.NewInt32(table, "power")
	s.AttributeID = field.NewInt32(table, "attribute_id")

	s.fillFieldMap()

	return s
}

func (s *skill) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *skill) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 6)
	s.fieldMap["id"] = s.ID
	s.fieldMap["name"] = s.Name
	s.fieldMap["description"] = s.Description
	s.fieldMap["ct"] = s.Ct
	s.fieldMap["power"] = s.Power
	s.fieldMap["attribute_id"] = s.AttributeID
}

func (s skill) clone(db *gorm.DB) skill {
	s.skillDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s skill) replaceDB(db *gorm.DB) skill {
	s.skillDo.ReplaceDB(db)
	return s
}

type skillDo struct{ gen.DO }

type ISkillDo interface {
	gen.SubQuery
	Debug() ISkillDo
	WithContext(ctx context.Context) ISkillDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISkillDo
	WriteDB() ISkillDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISkillDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISkillDo
	Not(conds ...gen.Condition) ISkillDo
	Or(conds ...gen.Condition) ISkillDo
	Select(conds ...field.Expr) ISkillDo
	Where(conds ...gen.Condition) ISkillDo
	Order(conds ...field.Expr) ISkillDo
	Distinct(cols ...field.Expr) ISkillDo
	Omit(cols ...field.Expr) ISkillDo
	Join(table schema.Tabler, on ...field.Expr) ISkillDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISkillDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISkillDo
	Group(cols ...field.Expr) ISkillDo
	Having(conds ...gen.Condition) ISkillDo
	Limit(limit int) ISkillDo
	Offset(offset int) ISkillDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISkillDo
	Unscoped() ISkillDo
	Create(values ...*model.Skill) error
	CreateInBatches(values []*model.Skill, batchSize int) error
	Save(values ...*model.Skill) error
	First() (*model.Skill, error)
	Take() (*model.Skill, error)
	Last() (*model.Skill, error)
	Find() ([]*model.Skill, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Skill, err error)
	FindInBatches(result *[]*model.Skill, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Skill) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISkillDo
	Assign(attrs ...field.AssignExpr) ISkillDo
	Joins(fields ...field.RelationField) ISkillDo
	Preload(fields ...field.RelationField) ISkillDo
	FirstOrInit() (*model.Skill, error)
	FirstOrCreate() (*model.Skill, error)
	FindByPage(offset int, limit int) (result []*model.Skill, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISkillDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s skillDo) Debug() ISkillDo {
	return s.withDO(s.DO.Debug())
}

func (s skillDo) WithContext(ctx context.Context) ISkillDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s skillDo) ReadDB() ISkillDo {
	return s.Clauses(dbresolver.Read)
}

func (s skillDo) WriteDB() ISkillDo {
	return s.Clauses(dbresolver.Write)
}

func (s skillDo) Session(config *gorm.Session) ISkillDo {
	return s.withDO(s.DO.Session(config))
}

func (s skillDo) Clauses(conds ...clause.Expression) ISkillDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s skillDo) Returning(value interface{}, columns ...string) ISkillDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s skillDo) Not(conds ...gen.Condition) ISkillDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s skillDo) Or(conds ...gen.Condition) ISkillDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s skillDo) Select(conds ...field.Expr) ISkillDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s skillDo) Where(conds ...gen.Condition) ISkillDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s skillDo) Order(conds ...field.Expr) ISkillDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s skillDo) Distinct(cols ...field.Expr) ISkillDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s skillDo) Omit(cols ...field.Expr) ISkillDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s skillDo) Join(table schema.Tabler, on ...field.Expr) ISkillDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s skillDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISkillDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s skillDo) RightJoin(table schema.Tabler, on ...field.Expr) ISkillDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s skillDo) Group(cols ...field.Expr) ISkillDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s skillDo) Having(conds ...gen.Condition) ISkillDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s skillDo) Limit(limit int) ISkillDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s skillDo) Offset(offset int) ISkillDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s skillDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISkillDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s skillDo) Unscoped() ISkillDo {
	return s.withDO(s.DO.Unscoped())
}

func (s skillDo) Create(values ...*model.Skill) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s skillDo) CreateInBatches(values []*model.Skill, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s skillDo) Save(values ...*model.Skill) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s skillDo) First() (*model.Skill, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Skill), nil
	}
}

func (s skillDo) Take() (*model.Skill, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Skill), nil
	}
}

func (s skillDo) Last() (*model.Skill, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Skill), nil
	}
}

func (s skillDo) Find() ([]*model.Skill, error) {
	result, err := s.DO.Find()
	return result.([]*model.Skill), err
}

func (s skillDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Skill, err error) {
	buf := make([]*model.Skill, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s skillDo) FindInBatches(result *[]*model.Skill, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s skillDo) Attrs(attrs ...field.AssignExpr) ISkillDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s skillDo) Assign(attrs ...field.AssignExpr) ISkillDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s skillDo) Joins(fields ...field.RelationField) ISkillDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s skillDo) Preload(fields ...field.RelationField) ISkillDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s skillDo) FirstOrInit() (*model.Skill, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Skill), nil
	}
}

func (s skillDo) FirstOrCreate() (*model.Skill, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Skill), nil
	}
}

func (s skillDo) FindByPage(offset int, limit int) (result []*model.Skill, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s skillDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s skillDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s skillDo) Delete(models ...*model.Skill) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *skillDo) withDO(do gen.Dao) *skillDo {
	s.DO = *do.(*gen.DO)
	return s
}
