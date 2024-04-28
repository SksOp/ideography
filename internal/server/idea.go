package server

import (
	"fmt"
	"ideography/internal/idea"
	"net/http"

	"github.com/labstack/echo/v4"
)

/* Schema of expexted request */

type request struct {
	Idea string `json:"idea"`
}

/* Schema of response */
type response struct {
	Message string `json:"message"`
	Idea    string `json:"idea"`
}

func sendError(message string) response {
	return response{message, ""}
}
func sendSuccess(idea string) response {
	return response{"Idea processed successfully", idea}
}

/* Route handler for hadling idea */
func (s *Server) Idea(c echo.Context) error {
	var req request

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, sendError(fmt.Sprintf("Error: %s", err)))
	}

	processedIdea, err := idea.HandleNewIdea(req.Idea)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, sendError("Currently we are unable to process your idea"))
	}

	return c.JSON(http.StatusOK, sendSuccess(processedIdea))

}
