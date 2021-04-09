package health

import (
	"github.com/blackhorseya/lobster/internal/app/lobster/biz/health/repo"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	er "github.com/blackhorseya/lobster/internal/pkg/entities/error"
	"go.uber.org/zap"
)

type impl struct {
	logger *zap.Logger
	repo   repo.IRepo
}

// NewImpl serve caller to create an IBiz
func NewImpl(logger *zap.Logger, repo repo.IRepo) IBiz {
	return &impl{logger: logger, repo: repo}
}

func (i *impl) Readiness(ctx contextx.Contextx) error {
	ok, err := i.repo.Ping(ctx)
	if err != nil {
		i.logger.Error(er.ErrDBConnect.Error())
		return er.ErrDBConnect
	}
	if !ok {
		i.logger.Error(er.ErrDBConnect.Error())
		return er.ErrDBConnect
	}

	return nil
}

func (i *impl) Liveness(ctx contextx.Contextx) error {
	ok, err := i.repo.Ping(ctx)
	if err != nil {
		i.logger.Error(er.ErrDBConnect.Error())
		return er.ErrDBConnect
	}
	if !ok {
		i.logger.Error(er.ErrDBConnect.Error())
		return er.ErrDBConnect
	}

	return nil
}
