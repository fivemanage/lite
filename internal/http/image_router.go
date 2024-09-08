package http

import (
	"github.com/labstack/echo/v4"
)

func (r *Server) imageRouterGroup(group *echo.Group) {
	imageGroup := group.Group("/image")

	imageGroup.POST("/", func(c echo.Context) error {
		var err error

		// get multipart form
		err = c.Request().ParseMultipartForm(10 << 20)
		if err != nil {
			return err
		}

		_, err = c.FormFile("image")

		return nil
	})

	imageGroup.GET("/:key", func(c echo.Context) error {
		return nil
	})

	imageGroup.DELETE("/:key", func(c echo.Context) error {
		return nil
	})
}
