package db

import (
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"os"
	"slices"
)

type DB[T interface{}] interface {
	Save(t T) (*Item[T], error)
	Delete(id uuid.UUID)
	GetAll() []T
}

type Item[T interface{}] struct {
	Data T         `json:"data"`
	Id   uuid.UUID `json:"id"`
}

type db[T interface{}] struct {
	Data     []Item[T]
	fileName string
}

func (d *db[T]) GetAll() []T {
	//TODO implement me
	panic("implement me")
}

func New[T interface{}](fileName string) (DB[T], error) {
	d := db[T]{
		Data:     make([]Item[T], 0),
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

	err = json.Unmarshal(byteValue, &d.Data)
	if err != nil {
		d.Data = make([]Item[T], 0)
	}

	return &d, nil
}

func (d *db[T]) Save(t T) (*Item[T], error) {
	data := Item[T]{
		Id:   uuid.New(),
		Data: t,
	}
	d.Data = append(d.Data, data)
	return &data, d.persist()
}

func (d *db[T]) Delete(id uuid.UUID) {
	slices.DeleteFunc(d.Data, func(x Item[T]) bool { return x.Id == id })
}

func (d *db[T]) persist() (err error) {
	fi, err := os.OpenFile(d.fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0660)
	if err != nil {
		return err
	}
	defer func() { err = fi.Close() }()

	bytes, err := json.MarshalIndent(d.Data, "", "  ")
	if err != nil {
		return err
	}

	_, err = fi.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}
