package service

import (
	"database/sql"
	"errors"
	"robinhood-assignment/internal/api"
	"robinhood-assignment/internal/dto"
	"robinhood-assignment/internal/model"
	"robinhood-assignment/internal/repository"
	"time"
)

type CommentService interface {
	Create(commentDTO *dto.CommentDTO, appointmentID, userID int64) error
	Update(commentDTO *dto.CommentDTO, id, userID int64) error
	Delete(id, userID int64) error
	GetByAppointmentID(appointmentID int64) ([]model.Comment, error)
}

type commentServiceImpl struct {
	commentRepo repository.CommentRepository
	userRepo    repository.UserRepository
}

func NewCommentServiceImpl(commentRepo repository.CommentRepository, userRepo repository.UserRepository) CommentService {
	return &commentServiceImpl{commentRepo, userRepo}
}

func (service *commentServiceImpl) Create(commentDTO *dto.CommentDTO, appointmentID, userID int64) error {

	user, err := service.userRepo.GetByID(userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return api.ErrUserNotFound
		}
		return err
	}

	comment := dto.ConvertCommentDTOToCommentModel(commentDTO)

	comment.UserID = userID
	comment.AppointmentID = appointmentID
	comment.CreatedDate = time.Now().UTC()
	comment.CreatedBy = user.DisplayName
	comment.UpdatedDate = time.Now().UTC()
	comment.UpdatedBy = user.DisplayName

	err = service.commentRepo.Create(comment)
	if err != nil {
		return err
	}

	return nil
}

func (service *commentServiceImpl) Update(commentDTO *dto.CommentDTO, id, userID int64) error {

	comment, err := service.commentRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return api.ErrDataNotFound
		}
		return err
	}

	if comment.UserID != userID {
		return api.ErrAccessDenied
	}

	user, err := service.userRepo.GetByID(userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return api.ErrUserNotFound
		}
		return err
	}

	comment.Description = commentDTO.Description
	comment.UpdatedDate = time.Now().UTC()
	comment.UpdatedBy = user.DisplayName

	err = service.commentRepo.Update(comment)
	if err != nil {
		return err
	}

	return nil
}

func (service *commentServiceImpl) Delete(id, userID int64) error {
	comment, err := service.commentRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return api.ErrDataNotFound
		}
		return err
	}

	if comment.UserID != userID {
		return api.ErrAccessDenied
	}

	err = service.commentRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (service *commentServiceImpl) GetByAppointmentID(appointmentID int64) ([]model.Comment, error) {
	return service.commentRepo.GetByAppointmentID(appointmentID)
}
