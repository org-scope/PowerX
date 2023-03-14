// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	auth "PowerX/internal/handler/auth"
	clue "PowerX/internal/handler/clue"
	contact "PowerX/internal/handler/contact"
	customer "PowerX/internal/handler/customer"
	department "PowerX/internal/handler/department"
	employee "PowerX/internal/handler/employee"
	permission "PowerX/internal/handler/permission"
	public "PowerX/internal/handler/public"
	"PowerX/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/department-tree/:depId",
				Handler: department.GetDepartmentTreeHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/departments/:id",
				Handler: department.GetDepartmentHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/departments",
				Handler: department.CreateDepartmentHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/departments/:id",
				Handler: department.DeleteDepartmentHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/department/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/op/sync-employees",
				Handler: employee.SyncEmployeesHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/employees/:id",
				Handler: employee.GetEmployeeHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/employees",
				Handler: employee.ListEmployeesHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/employees",
				Handler: employee.CreateEmployeeHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/options",
				Handler: employee.GetEmployeeOptionsHandler(serverCtx),
			},
			{
				Method:  http.MethodPatch,
				Path:    "/employees/:id",
				Handler: employee.UpdateEmployeeHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/employees/:id",
				Handler: employee.DeleteEmployeeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/op/reset-password",
				Handler: employee.ResetPasswordHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/employee/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/roles",
				Handler: permission.ListRolesHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/roles",
				Handler: permission.CreateRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/roles/:roleCode",
				Handler: permission.GetRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/roles/:roleCode",
				Handler: permission.PutRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/role-employee-ids/:roleCode",
				Handler: permission.GetRoleEmployeeIdsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/recourses",
				Handler: permission.ListRecoursesHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/op/assign-auth",
				Handler: permission.AssignAuthHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/permission/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/op/login/basic",
				Handler: auth.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/op/login/exchange/:type",
				Handler: auth.ExchangeHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user-info",
				Handler: auth.GetUserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/menu-roles",
				Handler: auth.GetMenuRolesHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/auth/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/live-qr-codes",
				Handler: contact.CreateLiveQRCodeHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/live-qr-codes",
				Handler: contact.ListLiveQRCodeHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/contact/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/live-qr-code/:uid",
				Handler: public.AccessLiveQRCodeHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/public/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/customers/:id",
				Handler: customer.GetCustomerHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/op/sync-customers",
				Handler: customer.BatchSyncCustomersHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/customer/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/clues",
				Handler: clue.ListCluesHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/clues",
				Handler: clue.CreateCluesHandler(serverCtx),
			},
			{
				Method:  http.MethodPatch,
				Path:    "/clues/:id",
				Handler: clue.PatchClueHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/clues/:id",
				Handler: clue.DeleteClueHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/clue/v1"),
	)
}
