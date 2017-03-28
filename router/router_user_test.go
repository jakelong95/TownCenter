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

func TestUserViewSuccess(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	id := uuid.NewUUID()
	tc, userMock := mockUser()
	userMock.On("GetByID", id.String()).Return(&models.User{}, nil)

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/user/"+id.String(), nil)
	tc.router.ServeHTTP(recorder, request)

	assert.Equal(200, recorder.Code)
}

func TestUserViewFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	id := uuid.NewUUID()
	tc, userMock := mockUser()
	userMock.On("GetByID", id.String()).Return(nil, fmt.Errorf("This is an error"))

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/user/"+id.String(), nil)
	tc.router.ServeHTTP(recorder, request)

	assert.Equal(500, recorder.Code)
}

// func TestUserViewAllSuccess(t *testing.T) {
// 	assert := assert.New(t)

// 	gin.SetMode(gin.TestMode)

// 	tc, userMock := mockUser()
// 	userMock.On("GetAll", 0, 20).Return(make([]*models.User, 0), nil)

// 	recorder := httptest.NewRecorder()
// 	request, _ := http.NewRequest("GET", "/api/user", nil)
// 	tc.router.ServeHTTP(recorder, request)

// 	assert.Equal(200, recorder.Code)
// }

// func TestUserViewAllFail(t *testing.T) {
// 	assert := assert.New(t)

// 	gin.SetMode(gin.TestMode)

// 	tc, userMock := mockUser()
// 	userMock.On("GetAll", 0, 20).Return(make([]*models.User, 0), fmt.Errorf("This is an error"))

// 	recorder := httptest.NewRecorder()
// 	request, _ := http.NewRequest("GET", "/api/user/list", nil)
// 	tc.router.ServeHTTP(recorder, request)

// 	assert.Equal(500, recorder.Code)
// }

// func TestUserViewAllParams(t *testing.T) {
// 	assert := assert.New(t)

// 	gin.SetMode(gin.TestMode)

// 	tc, userMock := mockUser()
// 	userMock.On("GetAll", 20, 40).Return(make([]*models.User, 0), nil)

// 	recorder := httptest.NewRecorder()
// 	request, _ := http.NewRequest("GET", "/api/user/list?offset=20&limit=40", nil)
// 	tc.router.ServeHTTP(recorder, request)

// 	assert.Equal(200, recorder.Code)
// }

func TestUserNewSuccess(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	tc, userMock := mockUser()
	userMock.On("Insert", mock.AnythingOfType("*models.User")).Return(nil)
	userMock.On("GetByEmail", "").Return(nil, nil)

	user := getUserString(models.NewUser("", "", "", "", "", "", "", "", "", "", ""))
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/user", user)
	tc.router.ServeHTTP(recorder, request)

	assert.NotNil(recorder.Header().Get("X-Auth"))
	assert.Equal(200, recorder.Code)
}

func TestUserNewFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	tc, userMock := mockUser()
	userMock.On("Insert", mock.AnythingOfType("*models.User")).Return(fmt.Errorf("This is an error"))
	userMock.On("GetByEmail", "").Return(nil, nil)

	user := getUserString(models.NewUser("", "", "", "", "", "", "", "", "", "", ""))
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/user", user)
	tc.router.ServeHTTP(recorder, request)

	assert.Equal(500, recorder.Code)
}

func TestUserAlreadyExists(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	tc, userMock := mockUser()
	userMock.On("GetByEmail", "").Return(models.NewUser("", "", "", "", "", "", "", "", "", "", ""), nil)

	user := getUserString(models.NewUser("", "", "", "", "", "", "", "", "", "", ""))
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/user", user)
	tc.router.ServeHTTP(recorder, request)

	assert.Equal(400, recorder.Code)
}

func TestUserNewInvalid(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	tc, userMock := mockUser()
	userMock.On("Insert", mock.AnythingOfType("*models.User")).Return(nil)
	userMock.On("GetByEmail", "").Return(nil, nil)

	user := bytes.NewReader([]byte("{\"id\": \"INVALID\"}"))
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/user", user)
	tc.router.ServeHTTP(recorder, request)

	assert.Equal(400, recorder.Code)
}

func TestUserUpdateSuccess(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	user := models.NewUser("", "", "", "", "", "", "", "", "", "", "")

	tc, userMock := mockUser()
	userMock.On("Update", user, user.ID.String()).Return(nil)

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("PUT", "/api/user/"+user.ID.String(), getUserString(user))
	tc.router.ServeHTTP(recorder, request)

	assert.Equal(200, recorder.Code)
}

func TestUserUpdateFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	user := models.NewUser("", "", "", "", "", "", "", "", "", "", "")

	tc, userMock := mockUser()
	userMock.On("Update", user, user.ID.String()).Return(fmt.Errorf("This is an error"))

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("PUT", "/api/user/"+user.ID.String(), getUserString(user))
	tc.router.ServeHTTP(recorder, request)

	assert.Equal(500, recorder.Code)
}

func TestUserUpdateInvalid(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	tc, userMock := mockUser()
	userMock.On("Update", mock.AnythingOfType("*models.User"), "").Return(fmt.Errorf("some error"))

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("PUT", "/api/user/INVALID", bytes.NewReader([]byte("{\"id\": \"INVALID\"}")))
	tc.router.ServeHTTP(recorder, request)

	assert.Equal(400, recorder.Code)
}

func TestUserDeleteSuccess(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	id := uuid.NewUUID()
	tc, userMock := mockUser()
	userMock.On("Delete", id.String()).Return(nil)

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("DELETE", "/api/user/"+id.String(), nil)
	tc.router.ServeHTTP(recorder, request)

	assert.Equal(200, recorder.Code)
}

func TestUserDeleteFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	id := uuid.NewUUID()
	tc, userMock := mockUser()
	userMock.On("Delete", id.String()).Return(fmt.Errorf("This is an error"))

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("DELETE", "/api/user/"+id.String(), nil)
	tc.router.ServeHTTP(recorder, request)

	assert.Equal(500, recorder.Code)
}

func getUserString(m *models.User) io.Reader {
	s, _ := json.Marshal(m)
	return bytes.NewReader(s)
}
