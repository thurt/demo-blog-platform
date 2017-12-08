package mysqlprovider_internal

import (
	"database/sql"

	"golang.org/x/net/context"

	helper "github.com/thurt/demo-blog-platform/cms/mysqlprovider_helper"
	pb "github.com/thurt/demo-blog-platform/cms/proto"
)

type provider_internal struct {
	db *sql.DB
	q  sqlQueryI
}

type sqlQueryI interface {
	GetUserPassword() string
}

type sqlQuery struct{}

func New(db *sql.DB) pb.CmsInternalServer {
	return &provider_internal{db, &sqlQuery{}}
}

func NewSqlQuery() sqlQueryI {
	return &sqlQuery{}
}

func (q *sqlQuery) GetUserPassword() string {
	return "SELECT password FROM users WHERE id=?"
}

func (p *provider_internal) GetUserPassword(ctx context.Context, r *pb.UserRequest) (*pb.UserPassword, error) {
	u := &pb.UserPassword{}
	err := p.db.QueryRow(p.q.GetUserPassword(), r.GetId()).Scan(&u.Password)

	if err != nil {
		return nil, helper.SqlErrorToGrpcError(err)
	}

	return u, nil
}
