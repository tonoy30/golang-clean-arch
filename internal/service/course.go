package service

import (
	"github.com/tonoy30/clean-arch/internal/dto"
	"github.com/tonoy30/clean-arch/internal/repository"
)

type CourseService interface {
	GetCourse(courseID int64) (*dto.Course, error)
	CreateCourse(course *dto.Course) (int64, error)
}
type courseService struct {
	dao repository.DAO
}

func NewCourseService(dao repository.DAO) CourseService {
	return &courseService{dao: dao}
}

func (c courseService) GetCourse(courseID int64) (*dto.Course, error) {
	course, err := c.dao.NewCourseRepository().GetCourse(courseID)
	if err != nil {
		return nil, err
	}

	fullCourse := dto.Course{
		ID:          course.ID,
		Title:       course.Title,
		Description: course.Description,
		Price:       course.Price,
		UserID:      course.UserID,
	}

	return &fullCourse, nil
}
func (c courseService) CreateCourse(course *dto.Course) (int64, error) {
	return 0, nil
}
