package health

import (
	"github.com/blackhorseya/lobster/internal/app/lobster/biz/health/repo"
	"github.com/blackhorseya/lobster/internal/pkg/entities/errors"
	"github.com/blackhorseya/lobster/internal/pkg/utils/contextx"
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
		i.logger.Error(errors.ErrDBConnect.Error())
		return errors.ErrDBConnect
	}
	if !ok {
		i.logger.Error(errors.ErrDBConnect.Error())
		return errors.ErrDBConnect
	}

	return nil
}

func (i *impl) Liveness(ctx contextx.Contextx) error {
	ok, err := i.repo.Ping(ctx)
	if err != nil {
		i.logger.Error(errors.ErrDBConnect.Error())
		return errors.ErrDBConnect
	}
	if !ok {
		i.logger.Error(errors.ErrDBConnect.Error())
		return errors.ErrDBConnect
	}

	return nil
}
