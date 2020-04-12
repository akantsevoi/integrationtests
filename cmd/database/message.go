package database

import (
	"database/sql"
	"fmt"
	"log"

	// main package shouldn't care about it
	_ "github.com/go-sql-driver/mysql"
)

// Options _
type Options struct {
	Name, Password, Host, Port, DBname string
}

// NewStorage _
func NewStorage(o Options) (Storage, error) {
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", o.Name, o.Password, o.Host, o.Port, o.DBname)
	log.Println(connectionString)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return Storage{}, err
	}
	return Storage{
		db: db,
	}, nil
}

// Storage _
type Storage struct {
	db *sql.DB
}

// Message _
type Message struct {
	ID   int64  `json:"id"`
	Text string `json:"text"`
}

// AddMessage _
func (s *Storage) AddMessage(text string) (int64, error) {
	res, err := s.db.Exec("INSERT INTO Message (message) VALUES (?)", text)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// GetMessages _
func (s *Storage) GetMessages() ([]Message, error) {
	rows, err := s.db.Query("SELECT id, message FROM Message")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Message
	for rows.Next() {
		var id int64
		var message string
		if err := rows.Scan(&id, &message); err != nil {
			return nil, err
		}
		result = append(result, Message{
			ID:   id,
			Text: message,
		})
	}

	return result, nil
}
