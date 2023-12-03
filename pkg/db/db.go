package db

import (
	"encoding/json"
	"io"
	"os"
)

type DB[T interface{}] interface {
	Save(t *T) error
	GetAll() []*T
}

type db[T interface{}] struct {
	data     []*T
	fileName string
}

func New[T interface{}](fileName string) (DB[T], error) {
	d := db[T]{
		data:     make([]*T, 0),
		fileName: fileName,
	}

	fi, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		return &d, err
	}
	defer func() { err = fi.Close() }()

	byteValue, err := io.ReadAll(fi)
	if err != nil {
		return &d, err
	}

	err = json.Unmarshal(byteValue, &d.data)
	if err != nil {
		d.data = make([]*T, 0)
	}

	return &d, nil
}
