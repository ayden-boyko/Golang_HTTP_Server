package tests

import (
	models "Golang_HTTP_Server/internal/models"
	"database/sql"
	"log"
	"testing"
	"time"

	_ "modernc.org/sqlite"
)

var testDataManager *models.DataManagerImpl
var validateDataManager *sql.DB

func init() {
	db, err := sql.Open("sqlite", "test.db")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	test_time := time.Date(2025, 1, 26, 16, 11, 35, 0, time.FixedZone("EST", -5*60*60))
	//drop table
	db.Exec("DROP TABLE IF EXISTS entries")

	//create table
	db.Exec("CREATE TABLE IF NOT EXISTS entries (id INTEGER PRIMARY KEY, base62_id TEXT, LongUrl TEXT, date_created DATE, UNIQUE(id, base62_id, LongUrl))")
	//add fake data
	_, err = db.Exec("INSERT OR IGNORE INTO entries (id, base62_id, LongUrl, date_created) VALUES (?, ?, ?, ?)", 1, "123", "https://test.com", test_time)
	db.Exec("INSERT OR IGNORE INTO entries (id, base62_id, LongUrl, date_created) VALUES (2, '456', 'https://youtube.com', ?)", test_time)

	if err != nil {
		log.Printf("Error inserting into database: %v", err)
	}

	testDataManager, _ = models.NewDataManager(db)
	validateDataManager = db
}

func TestDataManager(t *testing.T) {
	if testDataManager.Ping() != nil {
		t.Error("DataManager not initialized correctly:", testDataManager.Ping())
	}
}

func TestPushData(t *testing.T) {
	test_time := time.Date(2025, 1, 26, 16, 11, 35, 0, time.FixedZone("EST", -5*60*60))
	_, err := testDataManager.PushData(models.Entry{Id: 3, Base62_id: "789", LongUrl: "https://google.com", Date_Created: test_time}) // not pushing fake data

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	var longUrl string
	rows, _ := validateDataManager.Query("SELECT LongUrl FROM entries WHERE id = 3")

	if rows.Next() {
		rows.Scan(&longUrl)
	}

	defer rows.Close()

	if longUrl != "https://google.com" {
		t.Errorf("Expected long URL '%s', but got '%s'", "https://google.com", longUrl)
	}
}

func TestGetEntry(t *testing.T) {

	data, err := testDataManager.GetEntry(uint64(1))
	if err != nil {
		t.Errorf("Expected no error, but got: %v, long URL: %s", err, data)
	}

	if data != "https://test.com" {
		t.Errorf("Expected long URL '%s', but got '%s'", "https://test.com", data)
	}

}

func TestDataManagerClose(t *testing.T) {
	testDataManager.Close()

	if testDataManager.Ping() == nil {
		t.Error("DataManager not closed correctly")
	}
}
