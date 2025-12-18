package admin

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

// GetUsersTable returns the users table configuration
func GetUsersTable(ctx *table.Context) table.Table {
	usersTable := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := usersTable.GetInfo()
	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField("Name", "name", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("Email", "email", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("Phone", "phone", db.Varchar)
	info.AddField("Status", "status", db.Varchar).FieldDisplay(func(value types.FieldModel) interface{} {
		if value.Value == "active" {
			return types.Label().SetContent("Active").SetColor("success").GetContent()
		}
		return types.Label().SetContent("Inactive").SetColor("danger").GetContent()
	})
	info.AddField("Created At", "created_at", db.Datetime).FieldFilterable()
	info.AddField("Updated At", "updated_at", db.Datetime)

	info.SetTable("users").SetTitle("Users").SetDescription("Manage Users")

	formList := usersTable.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).FieldDisplayButCanNotEditWhenAdd()
	formList.AddField("Name", "name", db.Varchar, form.Text).FieldMust()
	formList.AddField("Email", "email", db.Varchar, form.Email).FieldMust()
	formList.AddField("Phone", "phone", db.Varchar, form.Text)
	formList.AddField("Status", "status", db.Varchar, form.Select).
		FieldOptions(types.FieldOptions{
			{Text: "Active", Value: "active"},
			{Text: "Inactive", Value: "inactive"},
		}).
		FieldDefault("active")
	formList.AddField("Created At", "created_at", db.Datetime, form.Datetime).FieldDisplayButCanNotEditWhenAdd()
	formList.AddField("Updated At", "updated_at", db.Datetime, form.Datetime).FieldDisplayButCanNotEditWhenAdd()

	formList.SetTable("users").SetTitle("Users").SetDescription("Manage Users")

	return usersTable
}

// GetStoresTable returns the stores table configuration
func GetStoresTable(ctx *table.Context) table.Table {
	storesTable := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := storesTable.GetInfo()
	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField("Name", "name", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("Address", "address", db.Varchar)
	info.AddField("Phone", "phone", db.Varchar)
	info.AddField("Email", "email", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("Status", "status", db.Varchar).FieldDisplay(func(value types.FieldModel) interface{} {
		if value.Value == "active" {
			return types.Label().SetContent("Active").SetColor("success").GetContent()
		}
		return types.Label().SetContent("Inactive").SetColor("danger").GetContent()
	})
	info.AddField("Created At", "created_at", db.Datetime).FieldFilterable()
	info.AddField("Updated At", "updated_at", db.Datetime)

	info.SetTable("stores").SetTitle("Stores").SetDescription("Manage Stores")

	formList := storesTable.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).FieldDisplayButCanNotEditWhenAdd()
	formList.AddField("Name", "name", db.Varchar, form.Text).FieldMust()
	formList.AddField("Address", "address", db.Varchar, form.Text)
	formList.AddField("Phone", "phone", db.Varchar, form.Text)
	formList.AddField("Email", "email", db.Varchar, form.Email)
	formList.AddField("Status", "status", db.Varchar, form.Select).
		FieldOptions(types.FieldOptions{
			{Text: "Active", Value: "active"},
			{Text: "Inactive", Value: "inactive"},
		}).
		FieldDefault("active")
	formList.AddField("Created At", "created_at", db.Datetime, form.Datetime).FieldDisplayButCanNotEditWhenAdd()
	formList.AddField("Updated At", "updated_at", db.Datetime, form.Datetime).FieldDisplayButCanNotEditWhenAdd()

	formList.SetTable("stores").SetTitle("Stores").SetDescription("Manage Stores")

	return storesTable
}

// GetBrandsTable returns the brands table configuration
func GetBrandsTable(ctx *table.Context) table.Table {
	brandsTable := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := brandsTable.GetInfo()
	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField("Name", "name", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("Description", "description", db.Text)
	info.AddField("Logo", "logo", db.Varchar).FieldDisplay(func(value types.FieldModel) interface{} {
		if value.Value != "" {
			return types.Img().SetSrc(value.Value).SetWidth("50").SetHeight("50").GetContent()
		}
		return "-"
	})
	info.AddField("Status", "status", db.Varchar).FieldDisplay(func(value types.FieldModel) interface{} {
		if value.Value == "active" {
			return types.Label().SetContent("Active").SetColor("success").GetContent()
		}
		return types.Label().SetContent("Inactive").SetColor("danger").GetContent()
	})
	info.AddField("Created At", "created_at", db.Datetime).FieldFilterable()
	info.AddField("Updated At", "updated_at", db.Datetime)

	info.SetTable("brands").SetTitle("Brands").SetDescription("Manage Brands")

	formList := brandsTable.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).FieldDisplayButCanNotEditWhenAdd()
	formList.AddField("Name", "name", db.Varchar, form.Text).FieldMust()
	formList.AddField("Description", "description", db.Text, form.TextArea)
	formList.AddField("Logo", "logo", db.Varchar, form.Text).FieldPlaceholder("Logo URL")
	formList.AddField("Status", "status", db.Varchar, form.Select).
		FieldOptions(types.FieldOptions{
			{Text: "Active", Value: "active"},
			{Text: "Inactive", Value: "inactive"},
		}).
		FieldDefault("active")
	formList.AddField("Created At", "created_at", db.Datetime, form.Datetime).FieldDisplayButCanNotEditWhenAdd()
	formList.AddField("Updated At", "updated_at", db.Datetime, form.Datetime).FieldDisplayButCanNotEditWhenAdd()

	formList.SetTable("brands").SetTitle("Brands").SetDescription("Manage Brands")

	return brandsTable
}

