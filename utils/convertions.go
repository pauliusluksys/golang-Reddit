package utils

import (
	"github.com/hashicorp/go-hclog"
	"strconv"
)

func StringToInt(logger hclog.Logger, s string) (*int, error) {
	result, err := strconv.Atoi(s)
	logger.Debug("converting: ", s)
	if err != nil {
		logger.Debug("error during convertion from string to integer: ", err.Error())
		return nil, err
	}
	return &result, nil
}
