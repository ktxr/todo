package database

import (
	"todo/models"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	Conn *sqlx.DB
}

var DB Database

func ConnectDB(dsn string) error {
	conn, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return err
	}
	DB.Conn = conn
	return nil
}

func InitDB() error {
	_, err := DB.Conn.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id			integer GENERATED ALWAYS AS IDENTITY,
			description	text,
			note		text
		)
	`)
	return err
}

func GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := DB.Conn.Select(&tasks, `
		SELECT id, description, note
		FROM tasks
	`)
	return tasks, err
}

func GetTaskByID(id int) (models.Task, error) {
	var task models.Task
	err := DB.Conn.Get(&task, `
		SELECT id, description, note
		FROM tasks
		WHERE id = $1
	`, id)
	return task, err
}

func Create(task models.Task) error {
	_, err := DB.Conn.NamedExec(`
		INSERT INTO tasks
		(note, description)
		VALUES (:note, :description)
	`, task)
	return err
}

func Delete(id int) error {
	_, err := DB.Conn.Exec(`
		DELETE FROM tasks
		WHERE id = $1
	`, id)
	return err
}
