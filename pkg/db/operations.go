package db

import (
	"encoding/json"
	"os"
)

func (d *db[T]) Save(t *T) error {
	d.data = append(d.data, t)
	return d.persist()
}

func (d *db[T]) GetAll() []*T {
	return d.data
}

func (d *db[T]) persist() (err error) {
	fi, err := os.OpenFile(d.fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0660)
	if err != nil {
		return err
	}
	defer func() { err = fi.Close() }()

	bytes, err := json.MarshalIndent(d.data, "", "  ")
	if err != nil {
		return err
	}

	_, err = fi.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}
