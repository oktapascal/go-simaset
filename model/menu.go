package model

import (
	"net/http"
)

type (
	Menu struct {
		Id       string      `json:"id"`
		Name     string      `json:"name"`
		Icon     string      `json:"icon"`
		PathUrl  string      `json:"path_url"`
		Children []MenuChild `json:"children"`
	}

	MenuChild struct {
		Id       string      `json:"id"`
		Name     string      `json:"name"`
		PathUrl  string      `json:"path_url"`
		Children []MenuChild `json:"children"`
	}

	MenuRepository interface {
		GetMenu(group string) *[]Menu
	}

	MenuService interface {
		GetMenu(group string) *[]Menu
	}

	MenuHandler interface {
		GetMenu() http.HandlerFunc
	}
)
