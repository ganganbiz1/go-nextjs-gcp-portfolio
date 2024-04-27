package service_test

import (
	"context"
	"testing"

	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/entity"
	mock_repository "github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/repository/mock"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/service"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_Artucle_Create(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock_repository.NewMockIfUserRepository(ctrl)
	articleRepo := mock_repository.NewMockIfArticleRepository(ctrl)

	service := service.NewArticleService(userRepo, articleRepo)

	type args struct {
		e *entity.Article
	}

	article := &entity.Article{
		ID:      0,
		Title:   "testタイトル",
		Content: "テスト記事本文",
		ImageID: 1,
		UserID:  1,
	}

	tests := []struct {
		testName string
		args     args
		setMock  func()
		wantErr  error
	}{
		{
			testName: "正常系",
			args:     args{e: article},
			setMock: func() {
				userRepo.EXPECT().Get(gomock.Any(), gomock.Any()).Return(&entity.User{ID: 1}, nil)
				articleRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: nil,
		},
		{
			testName: "異常系: 登録エラー",
			args:     args{e: article},
			setMock: func() {
				userRepo.EXPECT().Get(gomock.Any(), gomock.Any()).Return(&entity.User{ID: 1}, nil)
				articleRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(domain.ErrInternal)
			},
			wantErr: domain.ErrInternal,
		},
		{
			testName: "異常系: ユーザなし",
			args:     args{e: article},
			setMock: func() {
				userRepo.EXPECT().Get(gomock.Any(), gomock.Any()).Return(nil, nil)
			},
			wantErr: domain.ErrConflict,
		},
		{
			testName: "異常系: ユーザ取得エラー",
			args:     args{e: article},
			setMock: func() {
				userRepo.EXPECT().Get(gomock.Any(), gomock.Any()).Return(nil, domain.ErrInternal)
			},
			wantErr: domain.ErrInternal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			tt.setMock()
			err := service.Create(context.Background(), tt.args.e)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

// TODO: Create以外のメソッドもテスト
