package internal

import (
	models "Golang_HTTP_Server/internal/models"
)

type DataManager interface {
	GetEntry(uint64) (string, error)
	PushData(models.Entry) error
}
