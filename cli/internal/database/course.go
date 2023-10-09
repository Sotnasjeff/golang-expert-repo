package database

import (
	"database/sql"

	"github.com/google/uuid"
)

var (
	//INSERT_COURSES     = "INSERT INTO courses (id, name, description, category_id) VALUES ($1, $2, $3) RETURNING id"
	CREATE_COURSES     = "INSERT INTO courses (id, name, description, category_id) VALUES ($1, $2, $3, $4)"
	SELECT_ALL_COURSES = "SELECT id, name, description, category_id FROM courses"
	FIND_BY_ID         = "SELECT id, name, description, category_id FROM courses WHERE category_id = $1"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryID  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

//func (c *Course) CreateCourse(name, description, categoryID string) (*Course, error) {
//	id := uuid.New().String()
//	err := c.db.QueryRow(INSERT_COURSES, id, name, description, categoryID).Scan(&c.ID)
//	if err != nil {
//		return nil, err
//	}
//	return c, nil
//}

func (c *Course) CreateCourse(name, description, category_id string) (Course, error) {
	id := uuid.New().String()
	_, err := c.db.Exec(CREATE_COURSES, id, name, description, category_id)
	if err != nil {
		return Course{}, err
	}
	return Course{
		ID:          id,
		Name:        name,
		Description: description,
		CategoryID:  category_id,
	}, nil
}

func (c *Course) FindAll() ([]Course, error) {
	rows, err := c.db.Query(SELECT_ALL_COURSES)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	courses := []Course{}
	for rows.Next() {
		var id, name, description, category_id string
		if err := rows.Scan(&id, &name, &description, &category_id); err != nil {
			return nil, err
		}
		courses = append(courses, Course{
			ID:          id,
			Name:        name,
			Description: description,
			CategoryID:  category_id,
		})
	}

	return courses, nil

}

func (c *Course) FindByCategoryId(id string) ([]Course, error) {
	rows, err := c.db.Query(FIND_BY_ID, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	courses := []Course{}
	for rows.Next() {
		var id, name, description, category_id string
		if err := rows.Scan(&id, &name, &description, &category_id); err != nil {
			return nil, err
		}
		courses = append(courses, Course{
			ID:          id,
			Name:        name,
			Description: description,
			CategoryID:  category_id,
		})
	}
	return courses, nil
}
