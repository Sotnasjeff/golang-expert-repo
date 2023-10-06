package database

import (
	"database/sql"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

var (
	INSERT              = "INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)"
	SELECT_ALL          = "SELECT id, name, description FROM categories"
	FIND_COURSE_BY_ID   = "SELECT c.id, c.name, c.description FROM categories c JOIN courses co ON c.id = co.category_id WHERE co.id = $1"
	FIND_CATEGORY_BY_ID = "SELECT id, name, description FROM categories WHERE id = $1"
)

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{
		db: db,
	}
}

func (c *Category) CreateCategory(name, description string) (Category, error) {
	id := uuid.New().String()
	_, err := c.db.Exec(INSERT, id, name, description)
	if err != nil {
		return Category{}, err
	}
	return Category{ID: id, Name: name, Description: description}, nil
}

func (c *Category) FindAll() ([]Category, error) {
	rows, err := c.db.Query(SELECT_ALL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	categories := []Category{}
	for rows.Next() {
		var id, name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}
		categories = append(categories, Category{ID: id, Name: name, Description: description})
	}

	return categories, nil
}

func (c *Category) FindByCourseID(courseId string) (Category, error) {
	var id, name, description string
	err := c.db.QueryRow(FIND_COURSE_BY_ID, courseId).Scan(&id, &name, &description)
	if err != nil {
		return Category{}, err
	}
	return Category{ID: id, Name: name, Description: description}, nil
}

func (c *Category) FindCategoryById(categoryId string) (Category, error) {
	var id, name, description string
	err := c.db.QueryRow(FIND_CATEGORY_BY_ID, categoryId).Scan(&id, &name, &description)
	if err != nil {
		return Category{}, err
	}
	return Category{ID: id, Name: name, Description: description}, nil
}
