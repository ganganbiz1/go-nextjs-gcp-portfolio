package service_test

import (
	"context"
	"testing"

	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/entity"
	mock_gcp "github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/externals/gcp/mock"
	mock_repository "github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/repository/mock"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/service"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_User_Create(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock_repository.NewMockIfUserRepository(ctrl)
	firebase := mock_gcp.NewMockIfFirebaseClient(ctrl)

	service := service.NewUserService(userRepo, firebase)

	type args struct {
		e        *entity.User
		password string
	}

	user := &entity.User{
		ID:                 0,
		Email:              "test@test.com",
		Name:               "テスト太郎",
		FirebaseUID:        "",
		FirebaseProviderID: "",
	}

	tests := []struct {
		testName string
		args     args
		setMock  func()
		wantErr  error
	}{
		{
			testName: "正常系",
			args:     args{e: user, password: "test"},
			setMock: func() {
				userRepo.EXPECT().CountByEmail(gomock.Any(), gomock.Any()).Return(0, nil)
				userRepo.EXPECT().CountByName(gomock.Any(), gomock.Any()).Return(0, nil)
				firebase.EXPECT().Signup(gomock.Any(), gomock.Any(), gomock.Any()).Return(user, nil)
				userRepo.EXPECT().Create(gomock.Any(), user).Return(nil)
			},
			wantErr: nil,
		},
		{
			testName: "異常系",
			args:     args{e: user, password: "test"},
			setMock: func() {
				userRepo.EXPECT().CountByEmail(gomock.Any(), gomock.Any()).Return(0, nil)
				userRepo.EXPECT().CountByName(gomock.Any(), gomock.Any()).Return(0, nil)
				firebase.EXPECT().Signup(gomock.Any(), gomock.Any(), gomock.Any()).Return(user, nil)
				userRepo.EXPECT().Create(gomock.Any(), user).Return(domain.ErrInternal)
			},
			wantErr: domain.ErrInternal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			tt.setMock()
			err := service.Signup(context.Background(), tt.args.e, tt.args.password)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_User_Search(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock_repository.NewMockIfUserRepository(ctrl)
	firebase := mock_gcp.NewMockIfFirebaseClient(ctrl)

	service := service.NewUserService(userRepo, firebase)

	type args struct {
		id int
	}

	expectedUser := &entity.User{
		ID:                 1,
		Email:              "test@test.com",
		Name:               "テスト太郎",
		FirebaseUID:        "abcdefg",
		FirebaseProviderID: "test",
		Articles:           nil,
	}

	tests := []struct {
		testName string
		args     args
		setMock  func()
		want     *entity.User
		wantErr  error
	}{
		{
			testName: "正常系",
			args:     args{id: 1},
			setMock: func() {
				userRepo.EXPECT().GetWithArticles(gomock.Any(), gomock.Any()).Return(expectedUser, nil)
			},
			want:    expectedUser,
			wantErr: nil,
		},
		{
			testName: "異常系: 取得エラー",
			args:     args{id: 1},
			setMock: func() {
				userRepo.EXPECT().GetWithArticles(gomock.Any(), gomock.Any()).Return(nil, domain.ErrInternal)
			},
			want:    nil,
			wantErr: domain.ErrInternal,
		},
		{
			testName: "異常系: NotFoundエラー",
			args:     args{id: 1},
			setMock: func() {
				userRepo.EXPECT().GetWithArticles(gomock.Any(), gomock.Any()).Return(nil, domain.ErrNotFound)
			},
			want:    nil,
			wantErr: domain.ErrNotFound,
		},
		{
			testName: "異常系: NotFound",
			args:     args{id: 1},
			setMock: func() {
				userRepo.EXPECT().GetWithArticles(gomock.Any(), gomock.Any()).Return(nil, nil)
			},
			want:    nil,
			wantErr: domain.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			tt.setMock()
			got, err := service.Search(context.Background(), tt.args.id)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
