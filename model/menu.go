package model

import (
	"context"
	"database/sql"
	"net/http"
)

type (
	Menu struct {
		Id            string
		Name          string
		IconComponent string
		PathUrl       string
		Indeks        int8
		Children      int8
	}

	MenuChild struct {
		Id      string
		MenuId  string
		Name    string
		PathUrl string
		Indeks  int8
	}

	MenuResponse struct {
		Id            string `json:"id"`
		Name          string `json:"name"`
		IconComponent string `json:"icon_component"`
		PathUrl       string `json:"path_url"`
		Indeks        int8   `json:"indeks"`
		Children      int8   `json:"children"`
	}

	MenuChildResponse struct {
		Id      string `json:"id"`
		Name    string `json:"name"`
		PathUrl string `json:"path_url"`
		Indeks  int8   `json:"indeks"`
	}

	MenuRepository interface {
		GetMenus(ctx context.Context, tx *sql.Tx) *[]Menu
		GetMenuChildren(ctx context.Context, tx *sql.Tx, menuId string) *[]MenuChild
		FindMenuById(ctx context.Context, tx *sql.Tx, menuId string) (*Menu, error)
	}

	MenuService interface {
		GetMenus(ctx context.Context) []MenuResponse
		GetMenuChildren(ctx context.Context, menuId string) []MenuChildResponse
	}

	MenuHandler interface {
		GetMenus() http.HandlerFunc
		GetMenuChildren() http.HandlerFunc
	}
)
