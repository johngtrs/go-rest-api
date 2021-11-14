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

func (b *QueryBuilder) From(table string) *QueryBuilder {
	b.stmtFrom = "FROM " + table + " "
	return b
}

func (b *QueryBuilder) Join(table string, condition string) *QueryBuilder {
	b.stmtFrom += "JOIN " + table + " ON (" + condition + ") "
	return b
}

func (b *QueryBuilder) InnerJoin(table string, condition string) *QueryBuilder {
	b.stmtFrom += "INNER JOIN " + table + " ON (" + condition + ") "
	return b
}

func (b *QueryBuilder) LeftJoin(table string, condition string) *QueryBuilder {
	b.stmtFrom += "LEFT JOIN " + table + " ON (" + condition + ") "
	return b
}

func (b *QueryBuilder) RightJoin(table string, condition string) *QueryBuilder {
	b.stmtFrom += "RIGHT JOIN " + table + " ON (" + condition + ") "
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
	// log.Println(b.stmtSelect+b.stmtFrom+b.stmtWhere+b.stmtOrderBy, b.binds)
	return b.db.Select(b.entity, b.stmtSelect+b.stmtFrom+b.stmtWhere+b.stmtOrderBy, b.binds...)
}

func (b *QueryBuilder) ExecOne() error {
	// log.Println(b.stmtSelect+b.stmtFrom+b.stmtWhere+b.stmtOrderBy, b.binds)
	return b.db.Get(b.entity, b.stmtSelect+b.stmtFrom+b.stmtWhere+b.stmtOrderBy, b.binds...)
}
