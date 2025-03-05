package handler

import (
	"net/http"
	"robinhood-assignment/internal/api"
	"robinhood-assignment/internal/dto"
	"robinhood-assignment/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	commentService service.CommentService
}

func NewCommentHandler(commentService service.CommentService) *CommentHandler {
	return &CommentHandler{commentService}
}

func (h *CommentHandler) CreateComment(c *gin.Context) {
	var commentDTO dto.CommentDTO
	if err := c.ShouldBindJSON(&commentDTO); err != nil {
		apiErr := api.ErrInvalidInput
		api.RespondError(c, apiErr)
		return
	}

	appointmentID, err := strconv.ParseInt(c.Param("appointment_id"), 10, 64)
	if err != nil {
		apiErr := api.ErrInvalidInput
		api.RespondError(c, apiErr)
		return
	}

	userID := c.GetInt64("userID")

	err = h.commentService.Create(&commentDTO, appointmentID, userID)
	if err != nil {
		api.RespondError(c, err)
		return
	}

	api.RespondSuccessNoData(c, http.StatusCreated, "Comment created successfully")
}

func (h *CommentHandler) UpdateComment(c *gin.Context) {
	var commentDTO dto.CommentDTO
	if err := c.ShouldBindJSON(&commentDTO); err != nil {
		apiErr := api.ErrInvalidInput
		api.RespondError(c, apiErr)
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		apiErr := api.ErrInvalidInput
		api.RespondError(c, apiErr)
		return
	}

	userID := c.GetInt64("userID")

	err = h.commentService.Update(&commentDTO, id, userID)
	if err != nil {
		api.RespondError(c, err)
		return
	}

	api.RespondSuccessNoData(c, http.StatusOK, "Comment updated successfully")
}

func (h *CommentHandler) DeleteComment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		apiErr := api.ErrInvalidInput
		api.RespondError(c, apiErr)
		return
	}

	userID := c.GetInt64("userID")

	err = h.commentService.Delete(id, userID)
	if err != nil {
		api.RespondError(c, err)
		return
	}

	api.RespondSuccessNoData(c, http.StatusOK, "Comment deleted successfully")
}

func (h *CommentHandler) GetCommentsByAppointment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		apiErr := api.ErrInvalidInput
		api.RespondError(c, apiErr)
		return
	}

	comments, err := h.commentService.GetByAppointmentID(id)
	if err != nil {
		api.RespondError(c, err)
		return
	}

	api.RespondSuccess(c, http.StatusOK, "Comments retrieved successfully", comments)
}
