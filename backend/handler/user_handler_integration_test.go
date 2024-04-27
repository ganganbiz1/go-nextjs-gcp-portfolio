package handler_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/entity"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/externals/gcp"
	mock_gcp "github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/externals/gcp/mock"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/service"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/handler"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/infra/repository/model"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/usecase"
	test_wire "github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/wire/it"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_User_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	firebase := mock_gcp.NewMockIfFirebaseClient(ctrl)
	defer ctrl.Finish()

	di, cleanup, err := di()
	defer cleanup()

	if err != nil {
		t.Fatal("di error occured")
	}

	db := getDB(di)

	// データの事前準備
	err = truncate(db)
	if err != nil {
		t.Fatal("truncate error occured")
	}

	user := &entity.User{
		ID:                 0,
		Email:              "test@test.com",
		Name:               "テスト太郎",
		FirebaseUID:        "123abc",
		FirebaseProviderID: "firebase",
	}

	firebase.EXPECT().Signup(gomock.Any(), gomock.Any(), gomock.Any()).Return(user, nil)

	reqDatas := []struct {
		email    string
		name     string
		password string
	}{
		{
			email:    "test@test.com",
			name:     "test太郎",
			password: "12345abcde@",
		},
		// 必要に応じてここにデータバリエーション追加
	}

	for _, data := range reqDatas {
		json := fmt.Sprintf(
			`
				{
					"email": "%s",
					"name": "%s",
					"password": "%s@"
				}
			`,
			data.email,
			data.name,
			data.password,
		)

		e := di.DIC.TestServer.Echo
		req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(json))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.Set("example", "aa") // コンテキストが必要な場合は設定

		h := newUserHandler(di, firebase)
		apErr := h.Signup(c)

		var m model.User
		err = db.
			Where("email = ?", reqDatas[0].email).
			First(&m).
			Error

		if err != nil {
			t.Fatal("db error occured")
		}

		assert.Equal(t, nil, apErr)
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, reqDatas[0].email, m.Email)
		assert.Equal(t, reqDatas[0].name, m.Name)
		assert.Equal(t, user.FirebaseUID, m.FirebaseUID)
		assert.Equal(t, user.FirebaseProviderID, m.FirebaseProviderID)
	}
}

func Test_User_Search(t *testing.T) {
	ctrl := gomock.NewController(t)
	firebase := mock_gcp.NewMockIfFirebaseClient(ctrl)
	defer ctrl.Finish()

	di, cleanup, err := di()
	defer cleanup()

	if err != nil {
		t.Fatal("di error occured")
	}

	db := getDB(di)

	// データの事前準備
	err = truncate(db)
	if err != nil {
		t.Fatal("truncate error occured")
	}

	err = createUsers(db)
	if err != nil {
		t.Fatal("createUsers error occured")
	}

	expected := `
								{
									"message": "success",
									"data": {
										"id": 1,
										"email": "test1@test.com",
										"name": "テスト太郎",
										"firebaseUId": "abc",
										"firebaseProviderId": "firebase"
									}
								}
							`

	e := di.DIC.TestServer.Echo
	req := httptest.NewRequest(http.MethodGet, "/users/:userId", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:userId")
	c.SetParamNames("userId")
	c.SetParamValues("1")
	c.Set("example", "aa") // コンテキストが必要な場合は設定

	h := newUserHandler(di, firebase)
	apErr := h.Search(c)

	assert.Equal(t, nil, apErr)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, expected, rec.Body.String())
}

func newUserHandler(di *test_wire.DIManager, firebaseClient gcp.IfFirebaseClient) handler.IfUserHandler {
	s := service.NewUserService(
		di.DIC.UserRepository,
		firebaseClient,
	)

	u := usecase.NewUserUsecase(s)

	return handler.NewUserHandler(u)
}

func createUsers(db *gorm.DB) error {
	t := time.Now()
	users := []*model.User{
		{
			ID:                 1,
			Email:              "test1@test.com",
			Name:               "テスト太郎",
			FirebaseUID:        "abc",
			FirebaseProviderID: "firebase",
			CreatedAt:          t,
			UpdatedAt:          t,
		},
		{
			ID:                 2,
			Email:              "test2@test.com",
			Name:               "テスト次郎",
			FirebaseUID:        "def",
			FirebaseProviderID: "google.com",
			CreatedAt:          t,
			UpdatedAt:          t,
		},
	}

	if err := db.
		Create(users).
		Error; err != nil {
		return err
	}
	return nil
}
