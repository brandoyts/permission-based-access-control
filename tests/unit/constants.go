package unit

import (
	jwtProvider "github.com/brandoyts/permission-based-access-control/internal/adapters/jwt-provider"
	"github.com/brandoyts/permission-based-access-control/internal/adapters/persistence/memory"
)

var userRepo = memory.NewUserRepository()
var permissionRepo = memory.NewPermissionRepository()
var jwtAuth = jwtProvider.New("secret")
