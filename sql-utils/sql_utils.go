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
