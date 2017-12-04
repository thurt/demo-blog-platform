package mysqlprovider_helper

import (
	"database/sql"

	"github.com/VividCortex/mysqlerr"
	"github.com/go-sql-driver/mysql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func SqlErrorToGrpcError(err error) error {
	var e error

	if deviceErr, ok := err.(*mysql.MySQLError); ok {
		// these are just a few mysql errors that i found while trying to insert bad values
		// i'm not familiar with mysql errors but i suspect there will be more errors that come up commonly

		switch deviceErr.Number {
		case
			mysqlerr.ER_DUP_ENTRY,                       // tried to insert primary key value that already exists
			mysqlerr.ER_TRUNCATED_WRONG_VALUE_FOR_FIELD, // tried to insert/update a value with an incorrect type
			mysqlerr.ER_DATA_TOO_LONG,                   // tried to insert string that is too long
			mysqlerr.ER_NO_DEFAULT_FOR_FIELD,            // tried to insert row without passing a required value
			mysqlerr.ER_ROW_IS_REFERENCED_2,             // tried to update/delete a row key that is still referenced as a foreign key in another tabele
			mysqlerr.ER_NO_REFERENCED_ROW_2:             // tried to supply a foreign key value that is not found in parent table
			e = status.Error(codes.InvalidArgument, err.Error())
		case mysqlerr.ER_PARSE_ERROR: // tried to execute a sql statement that has syntax error(s)
			e = status.Error(codes.Internal, err.Error())
		default: // unknown
			e = status.Error(codes.Unknown, err.Error())
		}
	} else if err == sql.ErrNoRows { // this error is specific only to QueryRow invocations
		e = status.Error(codes.NotFound, err.Error())
	} else {
		e = status.Error(codes.Unknown, err.Error())
	}

	return e
}
