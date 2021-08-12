package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/TechBowl-japan/go-stations/model"
)

// A TODOService implements CRUD of TODO entities.
type TODOService struct {
	db *sql.DB
}

// NewTODOService returns new TODOService.
func NewTODOService(db *sql.DB) *TODOService {
	return &TODOService{
		db: db,
	}
}

// CreateTODO creates a TODO on DB.
func (s *TODOService) CreateTODO(ctx context.Context, subject, description string) (*model.TODO, error) {
	const (
		insert  = `INSERT INTO todos(subject, description) VALUES(?, ?)`
		confirm = `SELECT subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)
	args := []string{subject, description}
	query := make([]interface{}, 0, len(args))
	for _, a := range args {
		query = append(query, a)
	}
	// fmt.Println("-------------")
	// fmt.Println(subject, description, query)
	// fmt.Println("-------------")
	savetodo, err := s.db.ExecContext(ctx, insert, query...)
	if err != nil {
		return nil, err
	}
	todoid, err := savetodo.LastInsertId()
	if err != nil {
		return nil, err
	}

	var Subject string
	var Description string
	var CreatedAt time.Time
	var UpdatedAt time.Time

	result := s.db.QueryRowContext(ctx, confirm, todoid)
	if err = result.Scan(&Subject, &Description, &CreatedAt, &UpdatedAt); err != nil {
		return nil, err
	}

	todo := &model.TODO{
		Subject:     Subject,
		Description: Description,
		CreatedAt:   CreatedAt,
		UpdatedAt:   UpdatedAt,
	}

	return todo, err
}

// ReadTODO reads TODOs on DB.
func (s *TODOService) ReadTODO(ctx context.Context, prevID, size int64) ([]*model.TODO, error) {
	const (
		read       = `SELECT id, subject, description, created_at, updated_at FROM todos ORDER BY id DESC LIMIT ?`
		readWithID = `SELECT id, subject, description, created_at, updated_at FROM todos WHERE id < ? ORDER BY id DESC LIMIT ?`
	)

	return nil, nil
}

// UpdateTODO updates the TODO on DB.
func (s *TODOService) UpdateTODO(ctx context.Context, id int64, subject, description string) (*model.TODO, error) {
	const (
		update  = `UPDATE todos SET subject = ?, description = ? WHERE id = ?`
		confirm = `SELECT subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)

	return nil, nil
}

// DeleteTODO deletes TODOs on DB by ids.
func (s *TODOService) DeleteTODO(ctx context.Context, ids []int64) error {
	const deleteFmt = `DELETE FROM todos WHERE id IN (?%s)`

	return nil
}
