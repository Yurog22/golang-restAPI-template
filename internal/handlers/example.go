package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"template/internal/models"
	"template/internal/repository"
	"template/internal/responses"
)

type RepositoryHandler struct {
	Repo repository.Repository
}

func SetRepositoryHandler(repo repository.Repository) *RepositoryHandler {
	return &RepositoryHandler{
		Repo: repo,
	}
}

var validate *validator.Validate

// GetExample godoc
//
//	@Summary		Get example
//	@Description	Get example database records by id
//	@Tags			Example
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"id"
//	@Success		200	{object}	responses.Response{data=entities.Example}
//	@Failure		400	{object}	responses.Response
//	@Failure		500	{object}	responses.Response
//	@Router			/example/{id} [get]
func (h *RepositoryHandler) GetExample(c echo.Context) error {
	parsedID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return responses.BadRequestResponse(c, err)
	}

	example := models.Example{ID: parsedID}

	if err := h.Repo.GetExample(&example); err != nil {
		return responses.InternalServerErrorResponse(c, err)
	}

	return responses.SuccessResponse(c, example)
}

// ListExamples godoc
//
//	@Summary		List examples
//	@Description	List all example database records
//	@Tags			Example
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	responses.Response{data=entities.Example}
//	@Failure		400	{object}	responses.Response
//	@Failure		500	{object}	responses.Response
//	@Router			/example/list [get]
func (h *RepositoryHandler) ListExamples(c echo.Context) error {
	var examples []models.Example

	if err := h.Repo.ListExamples(&examples); err != nil {
		return responses.InternalServerErrorResponse(c, err)
	}

	return responses.SuccessResponse(c, examples)
}

// CreateExample godoc
//
//	@Summary		Create example
//	@Description	Create example database record
//	@Tags			Example
//	@Accept			json
//	@Produce		json
//	@Param			request	body		handlers.CreateExample.createExampleInput	true	"Create example"
//	@Success		200		{object}	responses.Response{data=entities.Example}
//	@Failure		400		{object}	responses.Response
//	@Failure		500		{object}	responses.Response
//	@Router			/example [post]
func (h *RepositoryHandler) CreateExample(c echo.Context) error {
	type createExampleInput struct {
		StringValue string `json:"stringValue" validate:"required" example:"string value"`
		IntValue    int    `json:"intValue" validate:"required" example:"4"`
	}

	var input createExampleInput

	if err := c.Bind(&input); err != nil {
		return responses.BadRequestResponse(c, err)
	}

	validate = validator.New()
	if err := validate.Struct(input); err != nil {
		return responses.BadRequestResponse(c, err)
	}

	example := models.Example{
		StringValue: input.StringValue,
		IntValue:    input.IntValue,
	}

	if err := h.Repo.CreateExample(&example); err != nil {
		return responses.InternalServerErrorResponse(c, err)
	}

	return responses.SuccessResponse(c, example)
}

// UpdateExample godoc
//
//	@Summary		Update example
//	@Description	Update example database record
//	@Tags			Example
//	@Accept			json
//	@Produce		json
//	@Param			request	body		handlers.UpdateExample.updateExampleInput	true	"Update example"
//	@Success		200		{object}	responses.Response{data=entities.Example}
//	@Failure		400		{object}	responses.Response
//	@Failure		500		{object}	responses.Response
//	@Router			/example [put]
func (h *RepositoryHandler) UpdateExample(c echo.Context) error {
	type updateExampleInput struct {
		ID          uuid.UUID `json:"id" validate:"required" example:"2badcae8-3b24-45f1-bdda-d8ed87c3d002"`
		StringValue string    `json:"stringValue" validate:"required" example:"string value"`
		IntValue    int       `json:"intValue" validate:"required" example:"4"`
	}

	var input updateExampleInput

	if err := c.Bind(&input); err != nil {
		return responses.BadRequestResponse(c, err)
	}

	validate = validator.New()
	if err := validate.Struct(input); err != nil {
		return responses.BadRequestResponse(c, err)
	}

	example := models.Example{
		ID:          input.ID,
		StringValue: input.StringValue,
		IntValue:    input.IntValue,
	}

	if err := h.Repo.UpdateExample(&example); err != nil {
		return responses.InternalServerErrorResponse(c, err)
	}

	return responses.SuccessResponse(c, example)
}

// DeleteExample godoc
//
//	@Summary		Delete example
//	@Description	Delete example database record by id
//	@Tags			Example
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"id"
//	@Success		200	{object}	responses.Response
//	@Failure		400	{object}	responses.Response
//	@Failure		500	{object}	responses.Response
//	@Router			/example/{id} [delete]
func (h *RepositoryHandler) DeleteExample(c echo.Context) error {
	parsedID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return responses.BadRequestResponse(c, err)
	}

	if err := h.Repo.DeleteExample(&models.Example{ID: parsedID}); err != nil {
		return responses.InternalServerErrorResponse(c, err)
	}

	return responses.SuccessResponse(c, nil)
}
