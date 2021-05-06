package service

import (
	"testing"

	"github.com/JesusG2000/hexsatisfaction/internal/model"
	"github.com/JesusG2000/hexsatisfaction/internal/model/dto"
	m "github.com/JesusG2000/hexsatisfaction/internal/service/mock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestUserRoleService_FindAllUser(t *testing.T) {
	tt := []struct {
		name   string
		fn     func(userRole *m.UserRole)
		expRes []model.User
		expErr error
	}{
		{
			name: "FindAllUser errors",
			fn: func(userRole *m.UserRole) {
				userRole.On("FindAllUser").
					Return(nil, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find users"),
		},
		{
			name: "All ok",
			fn: func(userRole *m.UserRole) {
				userRole.On("FindAllUser").
					Return([]model.User{
						{
							ID:       1,
							Login:    "test",
							Password: "test",
							RoleID:   dto.USER,
						},
						{
							ID:       2,
							Login:    "test1",
							Password: "test1",
							RoleID:   dto.USER,
						},
					}, nil)
			},
			expRes: []model.User{
				{
					ID:       1,
					Login:    "test",
					Password: "test",
					RoleID:   dto.USER,
				},
				{
					ID:       2,
					Login:    "test1",
					Password: "test1",
					RoleID:   dto.USER,
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			userRole := new(m.UserRole)
			service := NewUserRoleService(userRole)
			if tc.fn != nil {
				tc.fn(userRole)
			}
			user, err := service.FindAllUser()
			if err != nil {
				require.EqualError(t, tc.expErr, err.Error())
			} else {
				require.Nil(t, err)
			}
			require.Equal(t, tc.expRes, user)
		})
	}
}
