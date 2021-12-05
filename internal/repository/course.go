package repository

import (
	"github.com/Masterminds/squirrel"
	"github.com/tonoy30/go-commerce/internal/domain"
)

type CourseRepository interface {
	GetCourse(id int64) (*domain.Course, error)
}

type courseRepository struct{}

func (c *courseRepository) GetCourse(id int64) (*domain.Course, error) {
	qb := pgQb().
		Select("id", "user_id", "title", "description", "price").
		From(domain.CourseTableName).
		Where(squirrel.Eq{"id": id})

	var course domain.Course
	err := qb.QueryRow().Scan(&course.ID, &course.UserID, &course.Title, &course.Description, &course.Price)
	if err != nil {
		return &domain.Course{}, err
	}

	return &course, nil
}
