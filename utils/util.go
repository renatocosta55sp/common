package utils

import (
	"github.com/google/uuid"
	"github.com/renatocosta55sp/common/errors"
)

type Util struct{}

const (
	DEFAULT_TIMEZONE = "America/Sao_Paulo"
)

func (Util) StringToUUID(value string) (result uuid.UUID, err error) {
	result, err = uuid.Parse(value)
	if err != nil {
		err = errors.New("error.util.string.to.uuid")
		return
	}

	return
}
