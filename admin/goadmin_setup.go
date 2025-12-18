package admin

import (
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/plugins/admin"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	_ "github.com/GoAdminGroup/themes/adminlte"
)

var Adm *admin.Admin

// SetupGoAdmin initializes and adds the GoAdmin plugin to the engine
func SetupGoAdmin(eng *engine.Engine) error {
	// Initialize admin plugin with table generators
	Adm = admin.NewAdmin(&table.Generator{
		Generators: map[string]table.Generator{
			"users":  GetUsersTable,
			"stores": GetStoresTable,
			"brands": GetBrandsTable,
		},
	})

	// Add admin plugin to engine
	eng.AddPlugins(Adm)

	return nil
}

