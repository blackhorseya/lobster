package health

import (
	"github.com/blackhorseya/lobster/internal/app/biz/health/repo"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	er "github.com/blackhorseya/lobster/internal/pkg/entities/error"
)

type impl struct {
	repo repo.IRepo
}

// NewImpl serve caller to create an IBiz
func NewImpl(repo repo.IRepo) IBiz {
	return &impl{repo: repo}
}

func (i *impl) Readiness(ctx contextx.Contextx) error {
	ok, err := i.repo.Ping(ctx)
	if err != nil {
		ctx.WithField("err", err).Error(er.ErrDBConnect)
		return er.ErrDBConnect
	}
	if !ok {
		ctx.WithField("ok", ok).Error(er.ErrDBConnect)
		return er.ErrDBConnect
	}

	return nil
}

func (i *impl) Liveness(ctx contextx.Contextx) error {
	ok, err := i.repo.Ping(ctx)
	if err != nil {
		ctx.WithField("err", err).Error(er.ErrDBConnect)
		return er.ErrDBConnect
	}
	if !ok {
		ctx.WithField("ok", ok).Error(er.ErrDBConnect)
		return er.ErrDBConnect
	}

	return nil
}
