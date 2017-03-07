package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jakelong95/TownCenter/models"

	"github.com/pborman/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gopkg.in/gin-gonic/gin.v1"
)

func TestRoasterViewSuccess(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	id := uuid.NewUUID()
	tc, roasterMock := mockRoaster()
	roasterMock.On("GetByID", id.String()).Return(&models.Roaster{}, nil)

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/roaster/" + id.String(), nil)
	tc.router.ServeHTTP(recorder, request)

	assert.Equal(200, recorder.Code)
}

func TestRoasterViewFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	id := uuid.NewUUID()
	tc, roasterMock := mockRoaster()
	roasterMock.On("GetByID", id.String()).Return(nil, fmt.Errorf("This is an error"))

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/roaster/" + id.String(), nil)
	tc.router.ServeHTTP(recorder, request)

	assert.Equal(500, recorder.Code)
}

func TestRoasterViewAllSuccess(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	tc, roasterMock := mockRoaster()
	roasterMock.On("GetAll", 0, 20).Return(make([]*models.Roaster, 0), nil)

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/roaster", nil)
	tc.router.ServeHTTP(recorder, request)

	assert.Equal(200, recorder.Code)
}

func TestRoasterViewAllFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	tc, roasterMock := mockRoaster()
	roasterMock.On("GetAll", 0, 20).Return(make([]*models.Roaster, 0), fmt.Errorf("This is an error"))

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/roaster", nil)
	tc.router.ServeHTTP(recorder, request)

	assert.Equal(500, recorder.Code)
}

func TestRoasterViewAllParams(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	tc, roasterMock := mockRoaster()
	roasterMock.On("GetAll", 20, 40).Return(make([]*models.Roaster, 0), nil)

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/roaster?offset=20&limit=40", nil)
	tc.router.ServeHTTP(recorder, request)

	assert.Equal(200, recorder.Code)
}

func TestRoasterNewSuccess(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	tc, roasterMock := mockRoaster()
	roasterMock.On("Insert", mock.AnythingOfType("*models.Roaster")).Return(nil)

	roaster := getRoasterString(models.NewRoaster("", "", "", "", "", "", "", "", ""))
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/roaster", roaster)
	tc.router.ServeHTTP(recorder, request)

	assert.Equal(200, recorder.Code)
}

func TestRoasterNewFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	tc, roasterMock := mockRoaster()
	roasterMock.On("Insert", mock.AnythingOfType("*models.Roaster")).Return(fmt.Errorf("This is an error"))

	roaster := getRoasterString(models.NewRoaster("", "", "", "", "", "", "", "", ""))
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/roaster", roaster)
	tc.router.ServeHTTP(recorder, request)

	assert.Equal(500, recorder.Code)
}

/*func TestRoasterNewInvalid(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	tc, roasterMock := mockRoaster()
	roasterMock.On("Insert", mock.AnythingOfType("*models.Roaster")).Return(nil)

	roaster := bytes.NewReader([]byte("{\"id\": \"INVALID\"}"))
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/roaster", roaster)
	tc.router.ServeHTTP(recorder, request)

	assert.Equal(400, recorder.Code)
}*/

func TestRoasterUpdateSuccess(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	roaster := models.NewRoaster("", "", "", "", "", "", "", "", "")

	tc, roasterMock := mockRoaster()
	roasterMock.On("Update", roaster, roaster.ID.String()).Return(nil)

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("PUT", "/api/roaster/"+roaster.ID.String(), getRoasterString(roaster))
	tc.router.ServeHTTP(recorder, request)

	assert.Equal(200, recorder.Code)
}

func TestRoasterUpdateFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	roaster := models.NewRoaster("", "", "", "", "", "", "", "", "")

	tc, roasterMock := mockRoaster()
	roasterMock.On("Update", roaster, roaster.ID.String()).Return(fmt.Errorf("This is an error"))

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("PUT", "/api/roaster/"+roaster.ID.String(), getRoasterString(roaster))
	tc.router.ServeHTTP(recorder, request)

	assert.Equal(500, recorder.Code)
}

func TestRoasterUpdateInvalid(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	tc, roasterMock := mockRoaster()
	roasterMock.On("Update", mock.AnythingOfType("*models.Roaster"), "").Return(fmt.Errorf("some error"))

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("PUT", "/api/roaster/INVALID", bytes.NewReader([]byte("{\"id\": \"INVALID\"}")))
	tc.router.ServeHTTP(recorder, request)

	assert.Equal(400, recorder.Code)
}

func TestRoasterDeleteSuccess(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	id := uuid.NewUUID()
	tc, roasterMock := mockRoaster()
	roasterMock.On("Delete", id.String()).Return(nil)

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("DELETE", "/api/roaster/"+id.String(), nil)
	tc.router.ServeHTTP(recorder, request)

	assert.Equal(200, recorder.Code)
}

func TestRoasterDeleteFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	id := uuid.NewUUID()
	tc, roasterMock := mockRoaster()
	roasterMock.On("Delete", id.String()).Return(fmt.Errorf("This is an error"))

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("DELETE", "/api/roaster/"+id.String(), nil)
	tc.router.ServeHTTP(recorder, request)

	assert.Equal(500, recorder.Code)
}

func getRoasterString(m *models.Roaster) io.Reader {
	s, _ := json.Marshal(m)
	return bytes.NewReader(s)
}