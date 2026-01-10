package service

import (
	"errors"
	"testing"

	"github.com/ahmadzakyarifin/gin-jwt-auth/internal/dto"
	"github.com/ahmadzakyarifin/gin-jwt-auth/internal/entity"
	"github.com/ahmadzakyarifin/gin-jwt-auth/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func Test_authService_Login(t *testing.T) {
	hashPass,_ := bcrypt.GenerateFromPassword([]byte("kaylacantik"),bcrypt.DefaultCost)

	dummyData := &entity.User{
		ID: 1,
		Name: "kayla",
		Email: "kayla456@test.com",
		Password: string(hashPass),
		Role: "user",
	}

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		req     *dto.LoginRequest
		want    bool
		wantErr bool
		mockSetup func(mocks *mocks.AuthRepository)
	}{
		{
			name: "Login Berhasil",
			req: &dto.LoginRequest{
				Email: "kayla456@test.com",
				Password: "kaylacantik",
			},
			want: true,
			wantErr: false,
			mockSetup: func(mocks *mocks.AuthRepository) {
				mocks.On("FindByEmail","kayla456@test.com").Return(dummyData,nil)
			},
		},
		{
			name: "Login Gagal (email salah)",
			req: &dto.LoginRequest{
				Email: "salah@test.com",
				Password: "kaylacantik",
			},
			want: false,
			wantErr: true,
			mockSetup: func(mocks *mocks.AuthRepository) {
				mocks.On("FindByEmail","salah@test.com").Return(nil,errors.New("User NOt Found"))
			},
		},
		{
			name: "Login Gagal (password salah)",
			req: &dto.LoginRequest{
				Email: "kayla456@test.com",
				Password: "passwordsalah",
			},
			want: false,
			wantErr: true,
			mockSetup: func(mocks *mocks.AuthRepository) {
				mocks.On("FindByEmail","kayla456@test.com").Return(dummyData,nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(mocks.AuthRepository)

			if tt.mockSetup != nil {
				tt.mockSetup(mockRepo)
			}

			s := &authService{
				repo: mockRepo,
			}

			got,err := s.Login(tt.req)
			if tt.wantErr{
				assert.Error(t,err,"Seharusnya error")
			}else {
				assert.NoError(t,err,"Seharusnya tidak error")
			}

			if tt.want {
				assert.NotEmpty(t,got,"Seharusnya terisi")
			}else {
				assert.Empty(t,got,"Seharusnya tidak terisi atau kosong")
			}
			mockRepo.AssertExpectations(t)

		})
	}
}
