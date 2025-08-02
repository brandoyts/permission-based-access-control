package unit

import (
	"testing"

	"github.com/brandoyts/permission-based-access-control/internal/core/services/permissionService"
	"github.com/brandoyts/permission-based-access-control/tests/mock"
	"github.com/golang/mock/gomock"
)

func TestCreatePermission(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockPermissionRepository(ctrl)

	repo.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(nil)

	service := permissionService.New(repo)

	err := service.CreatePermission("can_create")
	if err != nil {
		t.Error(err)
	}
}
