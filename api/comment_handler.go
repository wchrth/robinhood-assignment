package api

import (
	"net/http"
	"robinhood-assignment/application/dto"
	"robinhood-assignment/application/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	CommentService service.CommentService
}

func NewCommentHandler(as service.CommentService) CommentHandler {
	return CommentHandler{CommentService: as}
}

func (ch CommentHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid request")
		return
	}

	appointment, err := ch.CommentService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	response := NewResponse(http.StatusOK, appointment)
	c.JSON(response.StatusCode, response)
}

func (ch CommentHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	if page <= 0 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	appointments, err := ch.CommentService.GetAll(offset, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	response := NewResponse(http.StatusOK, appointments)
	c.JSON(response.StatusCode, response)
}

func (ch CommentHandler) Create(c *gin.Context) {
	request := dto.CreateCommentRequest{}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := ch.CommentService.Create(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	response := NewResponse(http.StatusCreated, request)
	c.JSON(response.StatusCode, response)

}
