package repo

import (
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/pb"
)

// IRepo declare user repo service function
type IRepo interface {
	// QueryInfoByEmail
	QueryInfoByEmail(ctx contextx.Contextx, email string) (info *pb.Profile, err error)

	// UserRegister
	UserRegister(ctx contextx.Contextx, newUser pb.Profile) (info *pb.Profile, err error)
}
