package mysqlprovider_internal

import (
	"database/sql"

	"golang.org/x/net/context"

	pb "github.com/thurt/demo-blog-platform/cms/proto"
)

type provider_internal struct {
	db *sql.DB
	q  sqlQueryI
}

type sqlQueryI interface {
	GetUserPassword() string
}

type SqlQuery struct{}

func New(db *sql.DB) pb.CmsInternalServer {
	return &provider_internal{db, &SqlQuery{}}
}

func (q *SqlQuery) GetUserPassword() string {
	return "SELECT password FROM users WHERE id=?"
}

func (p *provider_internal) GetUserPassword(ctx context.Context, r *pb.UserRequest) (*pb.UserPassword, error) {
	u := &pb.UserPassword{}
	err := p.db.QueryRow(p.q.GetUserPassword(), r.GetId()).Scan(&u.Password)

	if err != nil {
		return nil, err
	}

	return u, nil
}
