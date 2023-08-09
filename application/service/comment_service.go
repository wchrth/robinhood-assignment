package service

import (
	"robinhood-assignment/application/dto"
	"robinhood-assignment/domain/entity"
	"robinhood-assignment/domain/repository"
)

type CommentService interface {
	GetByID(id int) (*dto.CommentResponse, error)
	GetAll(offset, limit int) ([]dto.CommentResponse, error)
	Create(request *dto.CreateCommentRequest) error
}

type commentService struct {
	CommentRepository repository.CommentRepository
}

func NewCommentService(ar repository.CommentRepository) CommentService {
	return commentService{CommentRepository: ar}
}

func (as commentService) GetByID(id int) (*dto.CommentResponse, error) {
	comment, err := as.CommentRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	commentResponse := dto.NewCommentResponse(*comment)

	return &commentResponse, nil
}

func (as commentService) GetAll(offset, limit int) ([]dto.CommentResponse, error) {
	comments, err := as.CommentRepository.GetAll(offset, limit)
	if err != nil {
		return nil, err
	}

	commentResponses := make([]dto.CommentResponse, 0)
	for _, comment := range comments {
		commentResponses = append(commentResponses, dto.NewCommentResponse(comment))
	}

	return commentResponses, nil
}

func (as commentService) Create(request *dto.CreateCommentRequest) error {

	comment := &entity.Comment{
		Comment:       request.Comment,
		AppointmentID: request.AppointmentID,
		UserID:        request.UserID,
	}

	return as.CommentRepository.Save(comment)
}
