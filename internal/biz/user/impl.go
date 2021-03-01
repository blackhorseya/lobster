package user

import (
	"github.com/blackhorseya/lobster/internal/biz/user/repo"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/pb"
)

type impl struct {
	repo repo.IRepo
}

// NewImpl serve caller to create an IBiz
func NewImpl(repo repo.IRepo) IBiz {
	return &impl{repo: repo}
}

func (i *impl) GetInfoByID(ctx contextx.Contextx, id string) (info *pb.Profile, err error) {
	// todo: 2021-02-28|17:31|doggy|implement me
	panic("implement me")
}

func (i *impl) GetInfoByEmail(ctx contextx.Contextx, email string) (info *pb.Profile, err error) {
	logger := ctx.WithField("email", email)

	ret, err := i.repo.QueryInfoByEmail(ctx, email)
	if err != nil {
		logger.WithError(err)
		return nil, err
	}

	return ret, nil
}

func (i *impl) GetInfoByAccessToken(ctx contextx.Contextx, token string) (info *pb.Profile, err error) {
	// todo: 2021-02-28|17:31|doggy|implement me
	panic("implement me")
}

func (i *impl) Signup(ctx contextx.Contextx, email string) (info *pb.Profile, err error) {
	// todo: 2021-02-28|17:31|doggy|implement me
	panic("implement me")
}

func (i *impl) Login(ctx contextx.Contextx, email, token string) (info *pb.Profile, err error) {
	// todo: 2021-02-28|17:31|doggy|implement me
	panic("implement me")
}

func (i *impl) Logout(ctx contextx.Contextx, user *pb.Profile) (err error) {
	// todo: 2021-02-28|17:31|doggy|implement me
	panic("implement me")
}
