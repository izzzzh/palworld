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

func newPalSkillMap(db *gorm.DB, opts ...gen.DOOption) palSkillMap {
	_palSkillMap := palSkillMap{}

	_palSkillMap.palSkillMapDo.UseDB(db, opts...)
	_palSkillMap.palSkillMapDo.UseModel(&model.PalSkillMap{})

	tableName := _palSkillMap.palSkillMapDo.TableName()
	_palSkillMap.ALL = field.NewAsterisk(tableName)
	_palSkillMap.ID = field.NewInt64(tableName, "id")
	_palSkillMap.PalID = field.NewInt64(tableName, "pal_id")
	_palSkillMap.SkillID = field.NewInt64(tableName, "skill_id")

	_palSkillMap.fillFieldMap()

	return _palSkillMap
}

// palSkillMap 帕鲁技能表
type palSkillMap struct {
	palSkillMapDo

	ALL     field.Asterisk
	ID      field.Int64
	PalID   field.Int64 // 帕鲁ID
	SkillID field.Int64 // 技能ID

	fieldMap map[string]field.Expr
}

func (p palSkillMap) Table(newTableName string) *palSkillMap {
	p.palSkillMapDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p palSkillMap) As(alias string) *palSkillMap {
	p.palSkillMapDo.DO = *(p.palSkillMapDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *palSkillMap) updateTableName(table string) *palSkillMap {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.PalID = field.NewInt64(table, "pal_id")
	p.SkillID = field.NewInt64(table, "skill_id")

	p.fillFieldMap()

	return p
}

func (p *palSkillMap) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *palSkillMap) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 3)
	p.fieldMap["id"] = p.ID
	p.fieldMap["pal_id"] = p.PalID
	p.fieldMap["skill_id"] = p.SkillID
}

func (p palSkillMap) clone(db *gorm.DB) palSkillMap {
	p.palSkillMapDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p palSkillMap) replaceDB(db *gorm.DB) palSkillMap {
	p.palSkillMapDo.ReplaceDB(db)
	return p
}

type palSkillMapDo struct{ gen.DO }

type IPalSkillMapDo interface {
	gen.SubQuery
	Debug() IPalSkillMapDo
	WithContext(ctx context.Context) IPalSkillMapDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPalSkillMapDo
	WriteDB() IPalSkillMapDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPalSkillMapDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPalSkillMapDo
	Not(conds ...gen.Condition) IPalSkillMapDo
	Or(conds ...gen.Condition) IPalSkillMapDo
	Select(conds ...field.Expr) IPalSkillMapDo
	Where(conds ...gen.Condition) IPalSkillMapDo
	Order(conds ...field.Expr) IPalSkillMapDo
	Distinct(cols ...field.Expr) IPalSkillMapDo
	Omit(cols ...field.Expr) IPalSkillMapDo
	Join(table schema.Tabler, on ...field.Expr) IPalSkillMapDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPalSkillMapDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPalSkillMapDo
	Group(cols ...field.Expr) IPalSkillMapDo
	Having(conds ...gen.Condition) IPalSkillMapDo
	Limit(limit int) IPalSkillMapDo
	Offset(offset int) IPalSkillMapDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPalSkillMapDo
	Unscoped() IPalSkillMapDo
	Create(values ...*model.PalSkillMap) error
	CreateInBatches(values []*model.PalSkillMap, batchSize int) error
	Save(values ...*model.PalSkillMap) error
	First() (*model.PalSkillMap, error)
	Take() (*model.PalSkillMap, error)
	Last() (*model.PalSkillMap, error)
	Find() ([]*model.PalSkillMap, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PalSkillMap, err error)
	FindInBatches(result *[]*model.PalSkillMap, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PalSkillMap) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPalSkillMapDo
	Assign(attrs ...field.AssignExpr) IPalSkillMapDo
	Joins(fields ...field.RelationField) IPalSkillMapDo
	Preload(fields ...field.RelationField) IPalSkillMapDo
	FirstOrInit() (*model.PalSkillMap, error)
	FirstOrCreate() (*model.PalSkillMap, error)
	FindByPage(offset int, limit int) (result []*model.PalSkillMap, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPalSkillMapDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p palSkillMapDo) Debug() IPalSkillMapDo {
	return p.withDO(p.DO.Debug())
}

func (p palSkillMapDo) WithContext(ctx context.Context) IPalSkillMapDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p palSkillMapDo) ReadDB() IPalSkillMapDo {
	return p.Clauses(dbresolver.Read)
}

func (p palSkillMapDo) WriteDB() IPalSkillMapDo {
	return p.Clauses(dbresolver.Write)
}

func (p palSkillMapDo) Session(config *gorm.Session) IPalSkillMapDo {
	return p.withDO(p.DO.Session(config))
}

func (p palSkillMapDo) Clauses(conds ...clause.Expression) IPalSkillMapDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p palSkillMapDo) Returning(value interface{}, columns ...string) IPalSkillMapDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p palSkillMapDo) Not(conds ...gen.Condition) IPalSkillMapDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p palSkillMapDo) Or(conds ...gen.Condition) IPalSkillMapDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p palSkillMapDo) Select(conds ...field.Expr) IPalSkillMapDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p palSkillMapDo) Where(conds ...gen.Condition) IPalSkillMapDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p palSkillMapDo) Order(conds ...field.Expr) IPalSkillMapDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p palSkillMapDo) Distinct(cols ...field.Expr) IPalSkillMapDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p palSkillMapDo) Omit(cols ...field.Expr) IPalSkillMapDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p palSkillMapDo) Join(table schema.Tabler, on ...field.Expr) IPalSkillMapDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p palSkillMapDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPalSkillMapDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p palSkillMapDo) RightJoin(table schema.Tabler, on ...field.Expr) IPalSkillMapDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p palSkillMapDo) Group(cols ...field.Expr) IPalSkillMapDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p palSkillMapDo) Having(conds ...gen.Condition) IPalSkillMapDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p palSkillMapDo) Limit(limit int) IPalSkillMapDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p palSkillMapDo) Offset(offset int) IPalSkillMapDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p palSkillMapDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPalSkillMapDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p palSkillMapDo) Unscoped() IPalSkillMapDo {
	return p.withDO(p.DO.Unscoped())
}

func (p palSkillMapDo) Create(values ...*model.PalSkillMap) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p palSkillMapDo) CreateInBatches(values []*model.PalSkillMap, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p palSkillMapDo) Save(values ...*model.PalSkillMap) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p palSkillMapDo) First() (*model.PalSkillMap, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PalSkillMap), nil
	}
}

func (p palSkillMapDo) Take() (*model.PalSkillMap, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PalSkillMap), nil
	}
}

func (p palSkillMapDo) Last() (*model.PalSkillMap, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PalSkillMap), nil
	}
}

func (p palSkillMapDo) Find() ([]*model.PalSkillMap, error) {
	result, err := p.DO.Find()
	return result.([]*model.PalSkillMap), err
}

func (p palSkillMapDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PalSkillMap, err error) {
	buf := make([]*model.PalSkillMap, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p palSkillMapDo) FindInBatches(result *[]*model.PalSkillMap, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p palSkillMapDo) Attrs(attrs ...field.AssignExpr) IPalSkillMapDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p palSkillMapDo) Assign(attrs ...field.AssignExpr) IPalSkillMapDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p palSkillMapDo) Joins(fields ...field.RelationField) IPalSkillMapDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p palSkillMapDo) Preload(fields ...field.RelationField) IPalSkillMapDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p palSkillMapDo) FirstOrInit() (*model.PalSkillMap, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PalSkillMap), nil
	}
}

func (p palSkillMapDo) FirstOrCreate() (*model.PalSkillMap, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PalSkillMap), nil
	}
}

func (p palSkillMapDo) FindByPage(offset int, limit int) (result []*model.PalSkillMap, count int64, err error) {
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

func (p palSkillMapDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p palSkillMapDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p palSkillMapDo) Delete(models ...*model.PalSkillMap) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *palSkillMapDo) withDO(do gen.Dao) *palSkillMapDo {
	p.DO = *do.(*gen.DO)
	return p
}