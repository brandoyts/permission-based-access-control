package unit

import "github.com/brandoyts/permission-based-access-control/internal/adapters/persistence/memory"

var userRepo = memory.NewUserRepository()
var permissionRepo = memory.NewPermissionRepository()
