package admin

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

// GetUsersTable returns the users table configuration with full CRUD operations
func GetUsersTable(ctx *context.Context) table.Table {
	usersTable := table.NewDefaultTable(ctx)

	info := usersTable.GetInfo()
	
	// Configure table display fields with sorting and filtering
	info.AddField("ID", "id", db.Int).
		FieldSortable().
		FieldWidth(80)
	
	info.AddField("Name", "name", db.Varchar).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).
		FieldSortable().
		FieldWidth(150)
	
	info.AddField("Email", "email", db.Varchar).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).
		FieldSortable().
		FieldWidth(200)
	
	info.AddField("Phone", "phone", db.Varchar).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).
		FieldWidth(120)
	
	info.AddField("Status", "status", db.Varchar).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).
		FieldDisplay(func(value types.FieldModel) interface{} {
			if value.Value == "active" {
				return `<span class="label label-success">Active</span>`
			}
			return `<span class="label label-danger">Inactive</span>`
		}).
		FieldWidth(100)
	
	info.AddField("Created At", "created_at", db.Datetime).
		FieldFilterable(types.FilterType{FormType: form.DatetimeRange}).
		FieldSortable().
		FieldWidth(150)
	
	info.AddField("Updated At", "updated_at", db.Datetime).
		FieldSortable().
		FieldWidth(150)

	// Configure table settings
	info.SetTable("users").
		SetTitle("Users Management").
		SetDescription("Create, Read, Update, and Delete Users").
		SetDefaultPageSize(10).
		SetFilterFormLayout(form.LayoutTwoCol)

	// Configure form fields for Create/Update operations
	formList := usersTable.GetForm()
	
	formList.AddField("ID", "id", db.Int, form.Default).
		FieldNotAllowAdd().
		FieldNotAllowEdit()
	
	formList.AddField("Name", "name", db.Varchar, form.Text).
		FieldMust().
		FieldPlaceholder("Enter user name").
		FieldHelpMsg("Full name of the user")
	
	formList.AddField("Email", "email", db.Varchar, form.Email).
		FieldMust().
		FieldPlaceholder("user@example.com").
		FieldHelpMsg("Valid email address")
	
	formList.AddField("Phone", "phone", db.Varchar, form.Text).
		FieldPlaceholder("+1234567890").
		FieldHelpMsg("Contact phone number")
	
	formList.AddField("Status", "status", db.Varchar, form.Select).
		FieldOptions(types.FieldOptions{
			{Text: "Active", Value: "active"},
			{Text: "Inactive", Value: "inactive"},
		}).
		FieldDefault("active").
		FieldMust()
	
	formList.AddField("Created At", "created_at", db.Datetime, form.Datetime).
		FieldNotAllowAdd().
		FieldNotAllowEdit().
		FieldNowWhenInsert()
	
	formList.AddField("Updated At", "updated_at", db.Datetime, form.Datetime).
		FieldNotAllowAdd().
		FieldNotAllowEdit().
		FieldNowWhenUpdate()

	// Configure form settings
	formList.SetTable("users").
		SetTitle("User Form").
		SetDescription("Add or Edit User Information")

	return usersTable
}

// GetStoresTable returns the stores table configuration with full CRUD operations
func GetStoresTable(ctx *context.Context) table.Table {
	storesTable := table.NewDefaultTable(ctx)

	info := storesTable.GetInfo()
	
	// Configure table display fields
	info.AddField("ID", "id", db.Int).
		FieldSortable().
		FieldWidth(80)
	
	info.AddField("Name", "name", db.Varchar).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).
		FieldSortable().
		FieldWidth(150)
	
	info.AddField("Address", "address", db.Varchar).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).
		FieldWidth(200)
	
	info.AddField("Phone", "phone", db.Varchar).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).
		FieldWidth(120)
	
	info.AddField("Email", "email", db.Varchar).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).
		FieldSortable().
		FieldWidth(180)
	
	info.AddField("Status", "status", db.Varchar).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).
		FieldDisplay(func(value types.FieldModel) interface{} {
			if value.Value == "active" {
				return `<span class="label label-success">Active</span>`
			}
			return `<span class="label label-danger">Inactive</span>`
		}).
		FieldWidth(100)
	
	info.AddField("Created At", "created_at", db.Datetime).
		FieldFilterable(types.FilterType{FormType: form.DatetimeRange}).
		FieldSortable().
		FieldWidth(150)
	
	info.AddField("Updated At", "updated_at", db.Datetime).
		FieldSortable().
		FieldWidth(150)

	// Configure table settings
	info.SetTable("stores").
		SetTitle("Stores Management").
		SetDescription("Create, Read, Update, and Delete Stores").
		SetDefaultPageSize(10).
		SetFilterFormLayout(form.LayoutTwoCol)

	// Configure form fields
	formList := storesTable.GetForm()
	
	formList.AddField("ID", "id", db.Int, form.Default).
		FieldNotAllowAdd().
		FieldNotAllowEdit()
	
	formList.AddField("Name", "name", db.Varchar, form.Text).
		FieldMust().
		FieldPlaceholder("Enter store name").
		FieldHelpMsg("Name of the store")
	
	formList.AddField("Address", "address", db.Varchar, form.TextArea).
		FieldPlaceholder("Enter store address").
		FieldHelpMsg("Physical address of the store")
	
	formList.AddField("Phone", "phone", db.Varchar, form.Text).
		FieldPlaceholder("+1234567890").
		FieldHelpMsg("Store contact phone number")
	
	formList.AddField("Email", "email", db.Varchar, form.Email).
		FieldPlaceholder("store@example.com").
		FieldHelpMsg("Store email address")
	
	formList.AddField("Status", "status", db.Varchar, form.Select).
		FieldOptions(types.FieldOptions{
			{Text: "Active", Value: "active"},
			{Text: "Inactive", Value: "inactive"},
		}).
		FieldDefault("active").
		FieldMust()
	
	formList.AddField("Created At", "created_at", db.Datetime, form.Datetime).
		FieldNotAllowAdd().
		FieldNotAllowEdit().
		FieldNowWhenInsert()
	
	formList.AddField("Updated At", "updated_at", db.Datetime, form.Datetime).
		FieldNotAllowAdd().
		FieldNotAllowEdit().
		FieldNowWhenUpdate()

	// Configure form settings
	formList.SetTable("stores").
		SetTitle("Store Form").
		SetDescription("Add or Edit Store Information")

	return storesTable
}

// GetBrandsTable returns the brands table configuration with full CRUD operations
func GetBrandsTable(ctx *context.Context) table.Table {
	brandsTable := table.NewDefaultTable(ctx)

	info := brandsTable.GetInfo()
	
	// Configure table display fields
	info.AddField("ID", "id", db.Int).
		FieldSortable().
		FieldWidth(80)
	
	info.AddField("Logo", "logo", db.Varchar).
		FieldDisplay(func(value types.FieldModel) interface{} {
			if value.Value != "" {
				return `<img src="` + value.Value + `" style="max-width: 60px; max-height: 60px; border-radius: 4px;" onerror="this.src='data:image/svg+xml,%3Csvg xmlns=\'http://www.w3.org/2000/svg\' width=\'60\' height=\'60\'%3E%3Crect width=\'60\' height=\'60\' fill=\'%23ddd\'/%3E%3Ctext x=\'50%25\' y=\'50%25\' text-anchor=\'middle\' dy=\'.3em\' fill=\'%23999\' font-size=\'12\'%3ENo Image%3C/text%3E%3C/svg%3E';" />`
			}
			return `<span class="text-muted">No Logo</span>`
		}).
		FieldWidth(100)
	
	info.AddField("Name", "name", db.Varchar).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).
		FieldSortable().
		FieldWidth(180)
	
	info.AddField("Description", "description", db.Text).
		FieldDisplay(func(value types.FieldModel) interface{} {
			desc := value.Value
			if len(desc) > 100 {
				return desc[:100] + "..."
			}
			return desc
		}).
		FieldWidth(250)
	
	info.AddField("Status", "status", db.Varchar).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).
		FieldDisplay(func(value types.FieldModel) interface{} {
			if value.Value == "active" {
				return `<span class="label label-success">Active</span>`
			}
			return `<span class="label label-danger">Inactive</span>`
		}).
		FieldWidth(100)
	
	info.AddField("Created At", "created_at", db.Datetime).
		FieldFilterable(types.FilterType{FormType: form.DatetimeRange}).
		FieldSortable().
		FieldWidth(150)
	
	info.AddField("Updated At", "updated_at", db.Datetime).
		FieldSortable().
		FieldWidth(150)

	// Configure table settings
	info.SetTable("brands").
		SetTitle("Brands Management").
		SetDescription("Create, Read, Update, and Delete Brands").
		SetDefaultPageSize(10).
		SetFilterFormLayout(form.LayoutTwoCol)

	// Configure form fields
	formList := brandsTable.GetForm()
	
	formList.AddField("ID", "id", db.Int, form.Default).
		FieldNotAllowAdd().
		FieldNotAllowEdit()
	
	formList.AddField("Name", "name", db.Varchar, form.Text).
		FieldMust().
		FieldPlaceholder("Enter brand name").
		FieldHelpMsg("Name of the brand")
	
	formList.AddField("Description", "description", db.Text, form.TextArea).
		FieldPlaceholder("Enter brand description").
		FieldHelpMsg("Detailed description of the brand")
	
	formList.AddField("Logo", "logo", db.Varchar, form.Url).
		FieldPlaceholder("https://example.com/logo.png").
		FieldHelpMsg("URL of the brand logo image").
		FieldDisplay(func(value types.FieldModel) interface{} {
			if value.Value != "" {
				return `<img src="` + value.Value + `" style="max-width: 100px; max-height: 100px; border-radius: 4px; margin-top: 10px;" onerror="this.style.display='none';" />`
			}
			return ""
		})
	
	formList.AddField("Status", "status", db.Varchar, form.Select).
		FieldOptions(types.FieldOptions{
			{Text: "Active", Value: "active"},
			{Text: "Inactive", Value: "inactive"},
		}).
		FieldDefault("active").
		FieldMust()
	
	formList.AddField("Created At", "created_at", db.Datetime, form.Datetime).
		FieldNotAllowAdd().
		FieldNotAllowEdit().
		FieldNowWhenInsert()
	
	formList.AddField("Updated At", "updated_at", db.Datetime, form.Datetime).
		FieldNotAllowAdd().
		FieldNotAllowEdit().
		FieldNowWhenUpdate()

	// Configure form settings
	formList.SetTable("brands").
		SetTitle("Brand Form").
		SetDescription("Add or Edit Brand Information")

	return brandsTable
}

