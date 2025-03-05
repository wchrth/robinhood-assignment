package dto

import "robinhood-assignment/internal/model"

type CommentDTO struct {
	Description string `json:"description" binding:"required"`
}

func ConvertCommentDTOToCommentModel(dto *CommentDTO) *model.Comment {
	return &model.Comment{
		Description: dto.Description,
	}
}
