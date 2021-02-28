package repo

import (
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/pb"
	"github.com/jmoiron/sqlx"
)

type impl struct {
	rw sqlx.DB
}

// NewImpl serve caller to create an IRepo
func NewImpl(rw sqlx.DB) IRepo {
	return &impl{rw: rw}
}

func (i *impl) QueryInfoByEmail(ctx contextx.Contextx, email string) (info *pb.Profile, err error) {
	panic("implement me")
}

func (i *impl) UserRegister(ctx contextx.Contextx, newUser pb.Profile) (info *pb.Profile, err error) {
	panic("implement me")
}
