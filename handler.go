package main

import (
	"context"
	"errors"
	"github.com/Stafford1986/test_manticore/usecase/entity"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type RequestsHandler struct {
	resumeUseCase  ResumeUseCase
	vacancyUseCase VacancyUseCase
}

type ResumeUseCase interface {
	Save(ctx context.Context, req *entity.ResumeEntity) error
	Update(ctx context.Context, req *entity.ResumeEntity) error
	Search(ctx context.Context, req *entity.SearchRequest) (entity.Resumes, error)
	Suggests(ctx context.Context, req string) ([]string, error)
}

type VacancyUseCase interface {
	Save(ctx context.Context, req *entity.VacancyEntity) error
	Update(ctx context.Context, req *entity.VacancyEntity) error
	Search(ctx context.Context, req *entity.SearchRequest) (entity.Vacancies, error)
	Suggests(ctx context.Context, req string) ([]string, error)
}

func NewHandler(resumeUseCase ResumeUseCase, vacancyUseCase VacancyUseCase) *RequestsHandler {
	return &RequestsHandler{
		resumeUseCase:  resumeUseCase,
		vacancyUseCase: vacancyUseCase,
	}
}

func (h *RequestsHandler) createResumeIndexHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resume := &entity.ResumeEntity{}
		body, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())

			return
		}

		err = resume.UnmarshalJSON(body)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "err bind resume")

			return
		}

		err = h.resumeUseCase.Save(ctx, resume)
		if err != nil {
			var validationErr *entity.ValidationErr
			if errors.As(err, &validationErr) {
				ctx.JSON(http.StatusBadRequest, err.Error())

				return
			}

			if errors.Is(err, entity.ErrForbiddenCharacter) {
				ctx.JSON(http.StatusBadRequest, err.Error())

				return
			}

			if errors.Is(err, entity.ErrDuplicateId) {
				ctx.JSON(http.StatusBadRequest, err.Error())

				return
			}

			ctx.JSON(http.StatusInternalServerError, err.Error())

			return
		}
	}
}

func (h *RequestsHandler) updateResumeIndexHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resume := &entity.ResumeEntity{}
		body, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())

			return
		}

		err = resume.UnmarshalJSON(body)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "err bind resume")

			return
		}

		err = h.resumeUseCase.Update(ctx, resume)
		if err != nil {
			var validationErr *entity.ValidationErr
			if errors.As(err, &validationErr) {
				ctx.JSON(http.StatusBadRequest, err.Error())

				return
			}

			if errors.Is(err, entity.ErrForbiddenCharacter) {
				ctx.JSON(http.StatusBadRequest, err.Error())

				return
			}

			ctx.JSON(http.StatusInternalServerError, err.Error())

			return
		}
	}
}

func (h *RequestsHandler) createSearchResumeHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())

			return
		}
		searchReq := &entity.SearchRequest{}

		err = searchReq.UnmarshalJSON(body)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "err bind search request")

			return
		}
		res, err := h.resumeUseCase.Search(ctx, searchReq)
		if err != nil {
			switch {
			case errors.Is(err, entity.ErrValidateSearchReq):
				ctx.JSON(http.StatusBadRequest, err.Error())

				return
			case errors.Is(err, entity.ErrForbiddenCharacter):
				ctx.JSON(http.StatusBadRequest, err.Error())

				return
			}
			ctx.JSON(http.StatusInternalServerError, err.Error())

			return
		}

		if len(res) == 0 {
			ctx.JSON(http.StatusNotFound, "resumes not found")

			return
		}

		resBytes, err := res.MarshalJSON()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())

			return
		}

		ctx.Data(http.StatusOK, "application/json", resBytes)
	}
}

func (h *RequestsHandler) createResumeSuggestHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		q := ctx.Query("req")
		if q == "" {
			ctx.JSON(http.StatusBadRequest, "err. empty req")

			return
		}

		res, err := h.resumeUseCase.Suggests(ctx, q)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())

			return
		}

		if len(res) == 0 {
			ctx.JSON(http.StatusNotFound, "not found")

			return
		}

		ctx.JSON(http.StatusOK, res)
	}
}

func (h *RequestsHandler) createVacancyIndexHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		vacancy := &entity.VacancyEntity{}
		body, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())

			return
		}

		err = vacancy.UnmarshalJSON(body)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "err bind vacancy")

			return
		}

		err = h.vacancyUseCase.Save(ctx, vacancy)
		if err != nil {
			var validationErr *entity.ValidationErr
			if errors.As(err, &validationErr) {
				ctx.JSON(http.StatusBadRequest, err.Error())

				return
			}

			if errors.Is(err, entity.ErrForbiddenCharacter) {
				ctx.JSON(http.StatusBadRequest, err.Error())

				return
			}

			if errors.Is(err, entity.ErrDuplicateId) {
				ctx.JSON(http.StatusBadRequest, err.Error())

				return
			}

			ctx.JSON(http.StatusInternalServerError, err.Error())

			return
		}
	}
}

func (h *RequestsHandler) updateVacanciesIndexHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		vacancy := &entity.VacancyEntity{}
		body, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())

			return
		}

		err = vacancy.UnmarshalJSON(body)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "err bind vacancy")

			return
		}

		err = h.vacancyUseCase.Update(ctx, vacancy)
		if err != nil {
			var validationErr *entity.ValidationErr
			if errors.As(err, &validationErr) {
				ctx.JSON(http.StatusBadRequest, err.Error())

				return
			}

			if errors.Is(err, entity.ErrForbiddenCharacter) {
				ctx.JSON(http.StatusBadRequest, err.Error())

				return
			}

			ctx.JSON(http.StatusInternalServerError, err.Error())

			return
		}
	}
}

func (h *RequestsHandler) createSearchVacancyHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())

			return
		}
		searchReq := &entity.SearchRequest{}

		err = searchReq.UnmarshalJSON(body)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "err bind search request")

			return
		}
		res, err := h.vacancyUseCase.Search(ctx, searchReq)
		if err != nil {
			switch {
			case errors.Is(err, entity.ErrValidateSearchReq):
				ctx.JSON(http.StatusBadRequest, err.Error())

				return
			case errors.Is(err, entity.ErrForbiddenCharacter):
				ctx.JSON(http.StatusBadRequest, err.Error())

				return
			}
			ctx.JSON(http.StatusInternalServerError, err.Error())

			return
		}

		if len(res) == 0 {
			ctx.JSON(http.StatusNotFound, "vacancies not found")

			return
		}

		resBytes, err := res.MarshalJSON()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())

			return
		}

		ctx.Data(http.StatusOK, "application/json", resBytes)
	}
}

func (h *RequestsHandler) createVacancySuggestHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		q := ctx.Query("req")
		if q == "" {
			ctx.JSON(http.StatusBadRequest, "err. empty req")

			return
		}

		res, err := h.resumeUseCase.Suggests(ctx, q)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())

			return
		}

		if len(res) == 0 {
			ctx.JSON(http.StatusNotFound, "not found")

			return
		}

		ctx.JSON(http.StatusOK, res)
	}
}
