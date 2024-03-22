// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q           = new(Query)
	Pal         *pal
	PalSkillMap *palSkillMap
	Skill       *skill
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	Pal = &Q.Pal
	PalSkillMap = &Q.PalSkillMap
	Skill = &Q.Skill
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:          db,
		Pal:         newPal(db, opts...),
		PalSkillMap: newPalSkillMap(db, opts...),
		Skill:       newSkill(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Pal         pal
	PalSkillMap palSkillMap
	Skill       skill
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:          db,
		Pal:         q.Pal.clone(db),
		PalSkillMap: q.PalSkillMap.clone(db),
		Skill:       q.Skill.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:          db,
		Pal:         q.Pal.replaceDB(db),
		PalSkillMap: q.PalSkillMap.replaceDB(db),
		Skill:       q.Skill.replaceDB(db),
	}
}

type queryCtx struct {
	Pal         IPalDo
	PalSkillMap IPalSkillMapDo
	Skill       ISkillDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Pal:         q.Pal.WithContext(ctx),
		PalSkillMap: q.PalSkillMap.WithContext(ctx),
		Skill:       q.Skill.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
