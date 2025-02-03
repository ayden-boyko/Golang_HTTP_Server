package internal

import (
	"database/sql"
	"errors"
	"fmt"
)

// TODO: USE locking to prevent race conditions

type DataManagerImpl struct {
	db *sql.DB
}

func NewDataManager(db *sql.DB) (*DataManagerImpl, error) {
	if db == nil {
		return nil, errors.New("database connection is nil")
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	return &DataManagerImpl{db: db}, nil
}

func (d *DataManagerImpl) GetEntry(id uint64) (string, error) {

	if d.db == nil {
		return "database connection is not established", errors.New("database connection is not established")
	}

	rows, err := d.db.Query("SELECT * FROM entries WHERE id = ?", id)
	if err != nil {
		return "Rows not found", err
	}
	defer rows.Close()

	if rows.Next() {
		var entry Entry
		if err := rows.Scan(&entry.Id, &entry.Base62_id, &entry.LongUrl, &entry.Date_Created); err != nil {
			return "error scanning", err
		}
		return entry.LongUrl, nil
	}
	return "No entry found", nil
}

func (d *DataManagerImpl) PushData(entry Entry) error {

	if d.db == nil {
		return errors.New("database connection is not established")
	}

	if entry.Id == 0 || entry.Base62_id == "" || entry.LongUrl == "" {
		return errors.New("invalid entry")
	}

	// check if entry already exists
	_, err := d.db.Query("SELECT * FROM entries WHERE LongUrl = ?", entry.LongUrl)
	if err != nil {
		return err
	}

	_, err = d.db.Exec("INSERT INTO entries (id, base62_id, LongUrl, date_created) VALUES (?, ?, ?, ?)", entry.Id, entry.Base62_id, entry.LongUrl, entry.Date_Created)
	if err != nil {
		return err
	}

	return nil
}

func (d *DataManagerImpl) Close() {
	d.db.Close()
}

func (d *DataManagerImpl) Ping() error {
	return d.db.Ping()
}

func (d *DataManagerImpl) Stats() sql.DBStats {
	return d.db.Stats()
}
