package internal

import (
	"database/sql"
)

type DataManagerImpl struct {
	db *sql.DB
}

func NewDataManager(db *sql.DB) *DataManagerImpl {
	return &DataManagerImpl{
		db: db,
	}
}

func (d DataManagerImpl) GetEntry(id uint64) (string, error) {
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

func (d DataManagerImpl) PushData(entry Entry) error {
	_, err := d.db.Exec("INSERT INTO entries (id, base62_id, long_url, date_created) VALUES (?, ?, ?, ?)", entry.Id, entry.Base62_id, entry.LongUrl, entry.Date_Created)
	if err != nil {
		return err
	}
	return nil
}

func (d DataManagerImpl) Close() {
	d.db.Close()
}

func (d DataManagerImpl) Ping() error {
	return d.db.Ping()
}

func (d DataManagerImpl) Stats() sql.DBStats {
	return d.db.Stats()
}
