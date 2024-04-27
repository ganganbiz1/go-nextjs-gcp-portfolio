package config

import (
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/logger"
)

func handleError(domainErr error, originErr error) error {
	if _, exist := domain.ErrMap[originErr]; exist {
		return originErr
	}
	logger.ErrorWithParams("", map[string]interface{}{
		"domainErr": domainErr.Error(),
		"originErr": originErr.Error(),
	})
	return domainErr
}
