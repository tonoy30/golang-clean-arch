package app

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (a app) HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "healthy")
}
func (a app) GetCourseById(c echo.Context) error {
	courseId, _ := strconv.Atoi(c.Param("id"))
	course, err := a.courseService.GetCourse(int64(courseId))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, course)
}
