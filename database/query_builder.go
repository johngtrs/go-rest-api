package database

import (
	"strconv"

	"github.com/jmoiron/sqlx"
)

type QueryBuilder struct {
	entity      interface{}
	db          *sqlx.DB
	stmtSelect  string
	stmtFrom    string
	stmtWhere   string
	binds       []interface{}
	stmtGroupBy string
	stmtHaving  string
	stmtOrderBy string
	stmtLimit   string
}

func NewQueryBuilder(db *sqlx.DB, entity interface{}) *QueryBuilder {
	return &QueryBuilder{db: db, entity: entity}
}

func (b *QueryBuilder) Select(cols string) *QueryBuilder {
	b.stmtSelect = "SELECT " + cols + " "
	return b
}

func (b *QueryBuilder) From(table string, alias string) *QueryBuilder {
	b.stmtFrom = "FROM " + table + " " + alias + " "
	return b
}

func (b *QueryBuilder) Join(table string, alias string, condition string) *QueryBuilder {
	b.stmtFrom += "JOIN " + table + " " + alias + " ON (" + condition + ") "
	return b
}

func (b *QueryBuilder) InnerJoin(table string, alias string, condition string) *QueryBuilder {
	b.stmtFrom += "INNER JOIN " + table + " " + alias + " ON (" + condition + ") "
	return b
}

func (b *QueryBuilder) LeftJoin(table string, alias string, condition string) *QueryBuilder {
	b.stmtFrom += "LEFT JOIN " + table + " " + alias + " ON (" + condition + ") "
	return b
}

func (b *QueryBuilder) RightJoin(table string, alias string, condition string) *QueryBuilder {
	b.stmtFrom += "RIGHT JOIN " + table + " " + alias + " ON (" + condition + ") "
	return b
}

func (b *QueryBuilder) Where(condition string, args ...interface{}) *QueryBuilder {
	b.stmtWhere = "WHERE " + condition + " "
	b.binds = args
	return b
}

func (b *QueryBuilder) GroupBy(cols string) *QueryBuilder {
	b.stmtGroupBy = "GROUP BY " + cols + " "
	return b
}

func (b *QueryBuilder) Having(condition string) *QueryBuilder {
	b.stmtHaving = "HAVING " + condition + " "
	return b
}

func (b *QueryBuilder) OrderBy(order string) *QueryBuilder {
	b.stmtOrderBy = "ORDER BY " + order + " "
	return b
}

func (b *QueryBuilder) Limit(limit int) *QueryBuilder {
	b.stmtLimit = "LIMIT " + strconv.Itoa(limit) + " "
	return b
}

func (b *QueryBuilder) Exec() error {
	return b.db.Select(b.entity, b.stmtSelect+b.stmtFrom+b.stmtWhere+b.stmtGroupBy+b.stmtHaving+b.stmtOrderBy+b.stmtLimit, b.binds...)
}

func (b *QueryBuilder) ExecOne() error {
	return b.db.Get(b.entity, b.stmtSelect+b.stmtFrom+b.stmtWhere+b.stmtGroupBy+b.stmtHaving+b.stmtOrderBy+b.stmtLimit, b.binds...)
}
