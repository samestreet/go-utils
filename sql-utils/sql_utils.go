package sqlutils

import (
	"fmt"
	"strings"
)

func CreateCreateSqlStatement(table string, fields []string, returning []string) string {
	sqlcmd := "INSERT INTO " + table + "(" + strings.Join(fields, ",") + ") VALUES("
	var values []string
	var c int
	for c = 1; c <= len(fields); c++ {
		values = append(values, fmt.Sprintf("$%d", c))
	}
	sqlcmd += strings.Join(values, ",") + ") RETURNING " + strings.Join(returning, ",")
	return sqlcmd
}

func CreateUpdateSqlStatement(table string, fields []string, whereClause string, returning []string) string {
	sqlcmd := "UPDATE " + table + " SET "
	c := 1
	var values []string
	for _, f := range fields {
		values = append(values, fmt.Sprintf("%s=$%d", f, c))
		c++
	}
	sqlcmd += strings.Join(values, ",")
	sqlcmd += " WHERE " + whereClause
	sqlcmd += " RETURNING " + strings.Join(returning, ",")
	return sqlcmd
}

func CreateSelectCountSqlStatement(table string) string {
	return fmt.Sprintf("SELECT COUNT(*) AS total FROM %s", table)
}

func CreateSelectCountWithClauseSqlStatement(table, where string) string {
	return fmt.Sprintf("SELECT COUNT(*) AS total FROM %s WHERE %s", table, where)
}

func CreateSelectSqlStatementWithOffset(table string, fields []string, page, size int, orderByColumn, orderByDirection string) string {
	offset := (page - 1) * size
	sqlcmd := fmt.Sprintf("SELECT %s FROM %s ORDER BY %s %s LIMIT %d OFFSET %d", strings.Join(fields, ","), table, orderByColumn, orderByDirection, size, offset)
	return sqlcmd
}

func CreateSelectSqlStatementWithOffsetAndClause(table string, fields []string, where string, page, size int, orderByColumn, orderByDirection string) string {
	offset := (page - 1) * size
	sqlcmd := fmt.Sprintf("SELECT %s FROM %s WHERE %s ORDER BY %s %s LIMIT %d OFFSET %d", strings.Join(fields, ","), table, where, orderByColumn, orderByDirection, size, offset)
	return sqlcmd
}
