package api

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		Name:        "Login",
		Method:      "POST",
		Pattern:     "/login",
		HandlerFunc: Login,
	},
	Route{
		Name:        "Logout",
		Method:      "POST",
		Pattern:     "/logout",
		HandlerFunc: Logout,
	},
	Route{
		Name:        "Change password",
		Method:      "PUT",
		Pattern:     "/change-password",
		HandlerFunc: ChangePassword,
	},
	Route{
		Name:        "List",
		Method:      "GET",
		Pattern:     "/library",
		HandlerFunc: List,
	},
	Route{
		Name:        "Add",
		Method:      "POST",
		Pattern:     "/library",
		HandlerFunc: Add,
	},
	Route{
		Name:        "Remove",
		Method:      "DELETE",
		Pattern:     "/library/{uuid}",
		HandlerFunc: Remove,
	},
	Route{
		Name:        "Update",
		Method:      "PUT",
		Pattern:     "/library/{uuid}",
		HandlerFunc: Update,
	},
	Route{
		Name:        "Generate OTP",
		Method:      "GET",
		Pattern:     "/library/{uuid}/otp",
		HandlerFunc: GeneratePassCode,
	},
	Route{
		Name:        "Change master password",
		Method:      "PUT",
		Pattern:     "/change-master-password",
		HandlerFunc: ChangeMasterPassword,
	},
	Route{
		Name:        "Generate password",
		Method:      "GET",
		Pattern:     "/generate-password",
		HandlerFunc: GeneratePassword,
	},
}
