// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package openapi

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// Article 記事
type Article struct {
	// Content 本文
	Content *string `json:"content,omitempty"`

	// Id 記事ID
	Id *int `json:"id,omitempty"`

	// ImageId 画像ID
	ImageId *int `json:"imageId,omitempty"`

	// Title タイトル
	Title *string `json:"title,omitempty"`
}

// User ユーザ
type User struct {
	// CognitoUserId cognitoUserId
	CognitoUserId *string `json:"cognitoUserId,omitempty"`

	// Email メールアドレス
	Email *string `json:"email,omitempty"`

	// Id ユーザID
	Id *int `json:"id,omitempty"`

	// Name ユーザ名前
	Name *string `json:"name,omitempty"`
}

// PostArticlesJSONBody defines parameters for PostArticles.
type PostArticlesJSONBody struct {
	// Content 本文
	Content string `json:"content"`

	// ImageId 画像ID
	ImageId int `json:"imageId"`

	// Title タイトル
	Title string `json:"title"`
}

// PutArticlesArticleIdJSONBody defines parameters for PutArticlesArticleId.
type PutArticlesArticleIdJSONBody struct {
	// Content 本文
	Content string `json:"content"`

	// ImageId 画像ID
	ImageId int `json:"imageId"`

	// Title タイトル
	Title string `json:"title"`
}

// PostUsersJSONBody defines parameters for PostUsers.
type PostUsersJSONBody struct {
	// Email メールアドレス
	Email string `json:"email"`

	// Name ユーザ名
	Name string `json:"name"`
}

// PutUsersUserIdJSONBody defines parameters for PutUsersUserId.
type PutUsersUserIdJSONBody struct {
	// CognitoUserId cognitoUserId
	CognitoUserId *string `json:"cognitoUserId,omitempty"`

	// Email メールアドレス
	Email *string `json:"email,omitempty"`

	// Name ユーザ名前
	Name *string `json:"name,omitempty"`
}

// PostArticlesJSONRequestBody defines body for PostArticles for application/json ContentType.
type PostArticlesJSONRequestBody PostArticlesJSONBody

// PutArticlesArticleIdJSONRequestBody defines body for PutArticlesArticleId for application/json ContentType.
type PutArticlesArticleIdJSONRequestBody PutArticlesArticleIdJSONBody

// PostUsersJSONRequestBody defines body for PostUsers for application/json ContentType.
type PostUsersJSONRequestBody PostUsersJSONBody

// PutUsersUserIdJSONRequestBody defines body for PutUsersUserId for application/json ContentType.
type PutUsersUserIdJSONRequestBody PutUsersUserIdJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// 記事一覧
	// (GET /articles)
	GetArticles(ctx echo.Context) error
	// 記事を作成
	// (POST /articles)
	PostArticles(ctx echo.Context) error
	// ユーザ詳細
	// (GET /articles/{articleId})
	GetArticlesArticleId(ctx echo.Context, articleId int) error
	// 記事編集
	// (PUT /articles/{articleId})
	PutArticlesArticleId(ctx echo.Context, articleId int) error
	// ヘルスチェック
	// (GET /healthcheck)
	GetHealthcheck(ctx echo.Context) error
	// ユーザ一覧
	// (GET /users)
	GetUsers(ctx echo.Context) error
	// ユーザーを作成
	// (POST /users)
	PostUsers(ctx echo.Context) error
	// ユーザ詳細
	// (GET /users/{userId})
	GetUsersUserId(ctx echo.Context, userId int) error
	// ユーザ編集
	// (PUT /users/{userId})
	PutUsersUserId(ctx echo.Context, userId int) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetArticles converts echo context to params.
func (w *ServerInterfaceWrapper) GetArticles(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetArticles(ctx)
	return err
}

// PostArticles converts echo context to params.
func (w *ServerInterfaceWrapper) PostArticles(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostArticles(ctx)
	return err
}

// GetArticlesArticleId converts echo context to params.
func (w *ServerInterfaceWrapper) GetArticlesArticleId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "articleId" -------------
	var articleId int

	err = runtime.BindStyledParameterWithOptions("simple", "articleId", ctx.Param("articleId"), &articleId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter articleId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetArticlesArticleId(ctx, articleId)
	return err
}

// PutArticlesArticleId converts echo context to params.
func (w *ServerInterfaceWrapper) PutArticlesArticleId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "articleId" -------------
	var articleId int

	err = runtime.BindStyledParameterWithOptions("simple", "articleId", ctx.Param("articleId"), &articleId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter articleId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PutArticlesArticleId(ctx, articleId)
	return err
}

// GetHealthcheck converts echo context to params.
func (w *ServerInterfaceWrapper) GetHealthcheck(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetHealthcheck(ctx)
	return err
}

// GetUsers converts echo context to params.
func (w *ServerInterfaceWrapper) GetUsers(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUsers(ctx)
	return err
}

// PostUsers converts echo context to params.
func (w *ServerInterfaceWrapper) PostUsers(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostUsers(ctx)
	return err
}

// GetUsersUserId converts echo context to params.
func (w *ServerInterfaceWrapper) GetUsersUserId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "userId" -------------
	var userId int

	err = runtime.BindStyledParameterWithOptions("simple", "userId", ctx.Param("userId"), &userId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUsersUserId(ctx, userId)
	return err
}

// PutUsersUserId converts echo context to params.
func (w *ServerInterfaceWrapper) PutUsersUserId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "userId" -------------
	var userId int

	err = runtime.BindStyledParameterWithOptions("simple", "userId", ctx.Param("userId"), &userId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PutUsersUserId(ctx, userId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/articles", wrapper.GetArticles)
	router.POST(baseURL+"/articles", wrapper.PostArticles)
	router.GET(baseURL+"/articles/:articleId", wrapper.GetArticlesArticleId)
	router.PUT(baseURL+"/articles/:articleId", wrapper.PutArticlesArticleId)
	router.GET(baseURL+"/healthcheck", wrapper.GetHealthcheck)
	router.GET(baseURL+"/users", wrapper.GetUsers)
	router.POST(baseURL+"/users", wrapper.PostUsers)
	router.GET(baseURL+"/users/:userId", wrapper.GetUsersUserId)
	router.PUT(baseURL+"/users/:userId", wrapper.PutUsersUserId)

}
