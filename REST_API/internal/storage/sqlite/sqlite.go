package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/Abdullah-Builds/Go_Projects/internal/config"
	"github.com/Abdullah-Builds/Go_Projects/internal/types"
	_ "modernc.org/sqlite"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite", cfg.Storage_Path)

	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`create table if not exists students(
	   id integer primary key autoincrement,
	   name varchar(50),
	   email varchar(50),
	   age integer
	)`)

	if err != nil {
		return nil, err
	}
	return &Sqlite{
		Db: db,
	}, nil

}

func (s *Sqlite) CreateStudent(name string, age int64, email string) (int64, error) {
	stmt, err := s.Db.Prepare("Insert into students(name, email, age) values(?,?,?)")

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(name, email, age)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Sqlite) GetStudentByID(id int) (*types.Student, error) {
	var student types.Student

	err := s.Db.QueryRow(
		"SELECT  name, email, age FROM students WHERE id = ?",
		id,
	).Scan(
		&student.Name,
		&student.Email,
		&student.Age,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("student not found")
		}
		return nil, err
	}

	return &student, nil
}
func (s *Sqlite) GetStudents(limit int, cursor int) ([]types.Student, *int, error) {
	query := `
		SELECT id, name, email, age
		FROM students
		WHERE id > ?
		ORDER BY id ASC
		LIMIT ?
	`

	rows, err := s.Db.Query(query, cursor, limit)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var students []types.Student
	var nextCursor *int

	for rows.Next() {
		var student types.Student
		var id int

		err := rows.Scan(
			&id,
			&student.Name,
			&student.Email,
			&student.Age,
		)
		if err != nil {
			return nil, nil, err
		}

		students = append(students, student)

		// Keep track of the last ID for the next cursor
		lastID := id
		nextCursor = &lastID

		fmt.Println(nextCursor)
	}

	if err := rows.Err(); err != nil {
		return nil, nil, err
	}

	return students, nextCursor, nil
}