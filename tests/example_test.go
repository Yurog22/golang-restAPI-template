package tests

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"template/internal/handlers"
	"template/internal/models"
	"template/internal/repository"
	"template/internal/responses"
	"testing"
)

func TestGet(t *testing.T) {
	t.Run("get successful return", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/example", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("e5b3b960-864a-47d4-b746-0dac54a1d810")

		mockRepo := &repository.MockRepository{}
		h := handlers.SetRepositoryHandler(mockRepo)

		if assert.NoError(t, h.GetExample(c)) {
			var response responses.Response
			assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &response))
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, http.StatusOK, response.Status)
			assert.Equal(t, "success", response.Message)

			responseDataBytes, _ := json.Marshal(response.Data)
			var exampleEntity models.Example
			assert.NoError(t, json.Unmarshal(responseDataBytes, &exampleEntity))
		}
	})

	t.Run("get bad request return", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/example", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("wrong uuid format")

		mockRepo := &repository.MockRepository{}
		h := handlers.SetRepositoryHandler(mockRepo)

		if assert.NoError(t, h.GetExample(c)) {
			var response responses.Response
			assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &response))
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, http.StatusBadRequest, response.Status)
			assert.Equal(t, "bad request", response.Message)
		}
	})
}

func TestList(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/example/list", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	mockRepo := &repository.MockRepository{}
	h := handlers.SetRepositoryHandler(mockRepo)

	if assert.NoError(t, h.ListExamples(c)) {
		var response responses.Response
		assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &response))
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, http.StatusOK, response.Status)
		assert.Equal(t, "success", response.Message)

		responseDataBytes, _ := json.Marshal(response.Data)
		var exampleEntities []models.Example
		assert.NoError(t, json.Unmarshal(responseDataBytes, &exampleEntities))
	}
}

func TestCreate(t *testing.T) {
	t.Run("create successful return", func(t *testing.T) {
		e := echo.New()

		postBody, _ := json.Marshal(echo.Map{
			"stringValue": "string example",
			"intValue":    1,
		})

		req := httptest.NewRequest(http.MethodPost, "/example", bytes.NewReader(postBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockRepo := &repository.MockRepository{}
		h := handlers.SetRepositoryHandler(mockRepo)

		if assert.NoError(t, h.CreateExample(c)) {
			var response responses.Response
			assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &response))
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, http.StatusOK, response.Status)
			assert.Equal(t, "success", response.Message)

			responseDataBytes, _ := json.Marshal(response.Data)
			var exampleEntity models.Example
			assert.NoError(t, json.Unmarshal(responseDataBytes, &exampleEntity))
		}
	})

	t.Run("create bad request return", func(t *testing.T) {
		e := echo.New()

		postBody, _ := json.Marshal(echo.Map{
			"stringValue": 42,
			"intValue":    "not int at all",
		})

		req := httptest.NewRequest(http.MethodPost, "/example", bytes.NewReader(postBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockRepo := &repository.MockRepository{}
		h := handlers.SetRepositoryHandler(mockRepo)

		if assert.NoError(t, h.CreateExample(c)) {
			var response responses.Response
			assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &response))
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, http.StatusBadRequest, response.Status)
			assert.Equal(t, "bad request", response.Message)
		}
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update successful return", func(t *testing.T) {
		e := echo.New()

		postBody, _ := json.Marshal(echo.Map{
			"id":          "a3953f48-6cb6-4dbe-ae67-19fb4d97060b",
			"stringValue": "string example",
			"intValue":    1,
		})

		req := httptest.NewRequest(http.MethodPut, "/example", bytes.NewReader(postBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockRepo := &repository.MockRepository{}
		h := handlers.SetRepositoryHandler(mockRepo)

		if assert.NoError(t, h.UpdateExample(c)) {
			var response responses.Response
			assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &response))
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, http.StatusOK, response.Status)
			assert.Equal(t, "success", response.Message)

			responseDataBytes, _ := json.Marshal(response.Data)
			var exampleEntity models.Example
			assert.NoError(t, json.Unmarshal(responseDataBytes, &exampleEntity))
		}
	})

	t.Run("update bad request return", func(t *testing.T) {
		e := echo.New()

		postBody, _ := json.Marshal(echo.Map{
			"id":          "not uuid at all",
			"stringValue": 42,
			"intValue":    "not int at all",
		})

		req := httptest.NewRequest(http.MethodPut, "/example", bytes.NewReader(postBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockRepo := &repository.MockRepository{}
		h := handlers.SetRepositoryHandler(mockRepo)

		if assert.NoError(t, h.UpdateExample(c)) {
			var response responses.Response
			assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &response))
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, http.StatusBadRequest, response.Status)
			assert.Equal(t, "bad request", response.Message)
		}
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete successful return", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/example", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("b74b663c-d880-4b78-bc02-7790b80e8848")

		mockRepo := &repository.MockRepository{}
		h := handlers.SetRepositoryHandler(mockRepo)

		if assert.NoError(t, h.DeleteExample(c)) {
			var response responses.Response
			assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &response))
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, http.StatusOK, response.Status)
			assert.Equal(t, "success", response.Message)
			assert.Nil(t, response.Data)
		}
	})

	t.Run("delete bad request return", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/example", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("wrong uuid format")

		mockRepo := &repository.MockRepository{}
		h := handlers.SetRepositoryHandler(mockRepo)

		if assert.NoError(t, h.DeleteExample(c)) {
			var response responses.Response
			assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &response))
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, http.StatusBadRequest, response.Status)
			assert.Equal(t, "bad request", response.Message)
		}
	})
}
