package service

import (
	"testing"

	"github.com/JesusG2000/hexsatisfaction/internal/model"
	"github.com/JesusG2000/hexsatisfaction/internal/model/dto"
	m "github.com/JesusG2000/hexsatisfaction/internal/service/mock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestUserRoleService_FindAllUser(t *testing.T) {
	a := assert.New(t)
	type test struct {
		name   string
		fn     func(userRole *m.UserRole, data test)
		expRes []model.User
		expErr error
	}
	tt := []test{
		{
			name: "FindAllUser errors",
			fn: func(userRole *m.UserRole, data test) {
				userRole.On("FindAllUser").
					Return(data.expRes, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find users"),
		},
		{
			name: "All ok",
			fn: func(userRole *m.UserRole, data test) {
				userRole.On("FindAllUser").
					Return(data.expRes, nil)
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
				tc.fn(userRole, tc)
			}
			user, err := service.FindAllUser()
			if err != nil {
				a.Equal(tc.expErr.Error(), err.Error())
			}
			a.Equal(tc.expRes, user)
		})
	}
}
