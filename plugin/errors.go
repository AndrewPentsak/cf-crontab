package plugin

import (
	"errors"
)

var (
	errMissingOrInvalidToken = errors.New("missing or invalid token")
	errNoDeployedCFCrontabFound = errors.New("no deployed cf-crontab server found in current space")
)
