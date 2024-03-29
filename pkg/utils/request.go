package utils

import (
	"strconv"

	"github.com/thangyt01/go-echo/pkg/constants"
	"github.com/thangyt01/go-echo/pkg/logger"
	"github.com/labstack/echo/v4"
)

func GetLoggerFromContext(c echo.Context) logger.Logger {
	logg := c.Get(constants.ContextKeyLogger)
	if logg != nil {
		return logg.(logger.Logger)
	}
	return logger.Log()
}

func GetResourceIdFromParam(c echo.Context, paramName string) int64 {
	id, _ := strconv.Atoi(c.Param(paramName))
	return int64(id)
}
