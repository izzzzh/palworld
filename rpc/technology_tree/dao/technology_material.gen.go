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

	"palworld/rpc/technology_tree/model"
)

func newTechnologyMaterial(db *gorm.DB, opts ...gen.DOOption) technologyMaterial {
	_technologyMaterial := technologyMaterial{}

	_technologyMaterial.technologyMaterialDo.UseDB(db, opts...)
	_technologyMaterial.technologyMaterialDo.UseModel(&model.TechnologyMaterial{})

	tableName := _technologyMaterial.technologyMaterialDo.TableName()
	_technologyMaterial.ALL = field.NewAsterisk(tableName)
	_technologyMaterial.ID = field.NewInt64(tableName, "id")
	_technologyMaterial.TechnologyID = field.NewInt64(tableName, "technology_id")
	_technologyMaterial.MaterialID = field.NewInt64(tableName, "material_id")
	_technologyMaterial.Cnt = field.NewInt32(tableName, "cnt")

	_technologyMaterial.fillFieldMap()

	return _technologyMaterial
}

// technologyMaterial 科技材料表
type technologyMaterial struct {
	technologyMaterialDo

	ALL          field.Asterisk
	ID           field.Int64
	TechnologyID field.Int64 // 科技ID
	MaterialID   field.Int64 // 材料ID
	Cnt          field.Int32 // 数量

	fieldMap map[string]field.Expr
}

func (t technologyMaterial) Table(newTableName string) *technologyMaterial {
	t.technologyMaterialDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t technologyMaterial) As(alias string) *technologyMaterial {
	t.technologyMaterialDo.DO = *(t.technologyMaterialDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *technologyMaterial) updateTableName(table string) *technologyMaterial {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt64(table, "id")
	t.TechnologyID = field.NewInt64(table, "technology_id")
	t.MaterialID = field.NewInt64(table, "material_id")
	t.Cnt = field.NewInt32(table, "cnt")

	t.fillFieldMap()

	return t
}

func (t *technologyMaterial) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *technologyMaterial) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 4)
	t.fieldMap["id"] = t.ID
	t.fieldMap["technology_id"] = t.TechnologyID
	t.fieldMap["material_id"] = t.MaterialID
	t.fieldMap["cnt"] = t.Cnt
}

func (t technologyMaterial) clone(db *gorm.DB) technologyMaterial {
	t.technologyMaterialDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t technologyMaterial) replaceDB(db *gorm.DB) technologyMaterial {
	t.technologyMaterialDo.ReplaceDB(db)
	return t
}

type technologyMaterialDo struct{ gen.DO }

type ITechnologyMaterialDo interface {
	gen.SubQuery
	Debug() ITechnologyMaterialDo
	WithContext(ctx context.Context) ITechnologyMaterialDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITechnologyMaterialDo
	WriteDB() ITechnologyMaterialDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITechnologyMaterialDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITechnologyMaterialDo
	Not(conds ...gen.Condition) ITechnologyMaterialDo
	Or(conds ...gen.Condition) ITechnologyMaterialDo
	Select(conds ...field.Expr) ITechnologyMaterialDo
	Where(conds ...gen.Condition) ITechnologyMaterialDo
	Order(conds ...field.Expr) ITechnologyMaterialDo
	Distinct(cols ...field.Expr) ITechnologyMaterialDo
	Omit(cols ...field.Expr) ITechnologyMaterialDo
	Join(table schema.Tabler, on ...field.Expr) ITechnologyMaterialDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITechnologyMaterialDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITechnologyMaterialDo
	Group(cols ...field.Expr) ITechnologyMaterialDo
	Having(conds ...gen.Condition) ITechnologyMaterialDo
	Limit(limit int) ITechnologyMaterialDo
	Offset(offset int) ITechnologyMaterialDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITechnologyMaterialDo
	Unscoped() ITechnologyMaterialDo
	Create(values ...*model.TechnologyMaterial) error
	CreateInBatches(values []*model.TechnologyMaterial, batchSize int) error
	Save(values ...*model.TechnologyMaterial) error
	First() (*model.TechnologyMaterial, error)
	Take() (*model.TechnologyMaterial, error)
	Last() (*model.TechnologyMaterial, error)
	Find() ([]*model.TechnologyMaterial, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TechnologyMaterial, err error)
	FindInBatches(result *[]*model.TechnologyMaterial, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TechnologyMaterial) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITechnologyMaterialDo
	Assign(attrs ...field.AssignExpr) ITechnologyMaterialDo
	Joins(fields ...field.RelationField) ITechnologyMaterialDo
	Preload(fields ...field.RelationField) ITechnologyMaterialDo
	FirstOrInit() (*model.TechnologyMaterial, error)
	FirstOrCreate() (*model.TechnologyMaterial, error)
	FindByPage(offset int, limit int) (result []*model.TechnologyMaterial, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITechnologyMaterialDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t technologyMaterialDo) Debug() ITechnologyMaterialDo {
	return t.withDO(t.DO.Debug())
}

func (t technologyMaterialDo) WithContext(ctx context.Context) ITechnologyMaterialDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t technologyMaterialDo) ReadDB() ITechnologyMaterialDo {
	return t.Clauses(dbresolver.Read)
}

func (t technologyMaterialDo) WriteDB() ITechnologyMaterialDo {
	return t.Clauses(dbresolver.Write)
}

func (t technologyMaterialDo) Session(config *gorm.Session) ITechnologyMaterialDo {
	return t.withDO(t.DO.Session(config))
}

func (t technologyMaterialDo) Clauses(conds ...clause.Expression) ITechnologyMaterialDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t technologyMaterialDo) Returning(value interface{}, columns ...string) ITechnologyMaterialDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t technologyMaterialDo) Not(conds ...gen.Condition) ITechnologyMaterialDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t technologyMaterialDo) Or(conds ...gen.Condition) ITechnologyMaterialDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t technologyMaterialDo) Select(conds ...field.Expr) ITechnologyMaterialDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t technologyMaterialDo) Where(conds ...gen.Condition) ITechnologyMaterialDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t technologyMaterialDo) Order(conds ...field.Expr) ITechnologyMaterialDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t technologyMaterialDo) Distinct(cols ...field.Expr) ITechnologyMaterialDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t technologyMaterialDo) Omit(cols ...field.Expr) ITechnologyMaterialDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t technologyMaterialDo) Join(table schema.Tabler, on ...field.Expr) ITechnologyMaterialDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t technologyMaterialDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITechnologyMaterialDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t technologyMaterialDo) RightJoin(table schema.Tabler, on ...field.Expr) ITechnologyMaterialDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t technologyMaterialDo) Group(cols ...field.Expr) ITechnologyMaterialDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t technologyMaterialDo) Having(conds ...gen.Condition) ITechnologyMaterialDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t technologyMaterialDo) Limit(limit int) ITechnologyMaterialDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t technologyMaterialDo) Offset(offset int) ITechnologyMaterialDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t technologyMaterialDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITechnologyMaterialDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t technologyMaterialDo) Unscoped() ITechnologyMaterialDo {
	return t.withDO(t.DO.Unscoped())
}

func (t technologyMaterialDo) Create(values ...*model.TechnologyMaterial) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t technologyMaterialDo) CreateInBatches(values []*model.TechnologyMaterial, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t technologyMaterialDo) Save(values ...*model.TechnologyMaterial) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t technologyMaterialDo) First() (*model.TechnologyMaterial, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TechnologyMaterial), nil
	}
}

func (t technologyMaterialDo) Take() (*model.TechnologyMaterial, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TechnologyMaterial), nil
	}
}

func (t technologyMaterialDo) Last() (*model.TechnologyMaterial, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TechnologyMaterial), nil
	}
}

func (t technologyMaterialDo) Find() ([]*model.TechnologyMaterial, error) {
	result, err := t.DO.Find()
	return result.([]*model.TechnologyMaterial), err
}

func (t technologyMaterialDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TechnologyMaterial, err error) {
	buf := make([]*model.TechnologyMaterial, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t technologyMaterialDo) FindInBatches(result *[]*model.TechnologyMaterial, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t technologyMaterialDo) Attrs(attrs ...field.AssignExpr) ITechnologyMaterialDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t technologyMaterialDo) Assign(attrs ...field.AssignExpr) ITechnologyMaterialDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t technologyMaterialDo) Joins(fields ...field.RelationField) ITechnologyMaterialDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t technologyMaterialDo) Preload(fields ...field.RelationField) ITechnologyMaterialDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t technologyMaterialDo) FirstOrInit() (*model.TechnologyMaterial, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TechnologyMaterial), nil
	}
}

func (t technologyMaterialDo) FirstOrCreate() (*model.TechnologyMaterial, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TechnologyMaterial), nil
	}
}

func (t technologyMaterialDo) FindByPage(offset int, limit int) (result []*model.TechnologyMaterial, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t technologyMaterialDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t technologyMaterialDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t technologyMaterialDo) Delete(models ...*model.TechnologyMaterial) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *technologyMaterialDo) withDO(do gen.Dao) *technologyMaterialDo {
	t.DO = *do.(*gen.DO)
	return t
}
