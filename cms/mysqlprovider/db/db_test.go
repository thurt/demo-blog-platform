// +build integration

package dbtest

import (
	"io/ioutil"
	"testing"

	"github.com/xwb1989/sqlparser"
)

func checkSyntax(sql string, t *testing.T, strict bool) {
	var err error
	if strict {
		_, err = sqlparser.ParseStrictDDL(sql)
	} else {
		_, err = sqlparser.Parse(sql)
	}

	if err != nil {
		t.Error(t.Name(), err, sql)
	}
}

func TestCreateTablesSQLSyntax(t *testing.T) {
	d, err := ioutil.ReadFile("./docker-entrypoint-initdb.d/1_createTables.sql")
	if err != nil {
		panic(err)
	}

	sql := string(d)

	var stmt string
	for sql != "\n" {
		stmt, sql, err = sqlparser.SplitStatement(sql)
		if err != nil {
			panic(err)
		}

		// NOTE: using strict false here b/c the parser isnt able to parse FOREIGN KEY https://github.com/xwb1989/sqlparser/issues/26
		checkSyntax(stmt, t, false)
	}
}
