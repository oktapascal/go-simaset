package menu

import (
	"context"
	"database/sql"
	"github.com/oktapascal/go-simpro/model"
)

type Repository struct{}

func (rpo *Repository) GetMenus(ctx context.Context, tx *sql.Tx) *[]model.Menu {
	query := "select id, name, icon_component, path_url, indeks from menus where deleted_at is null order by indeks"

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}()

	var menus []model.Menu
	for rows.Next() {
		var menu model.Menu
		err = rows.Scan(&menu.Id, &menu.Name, &menu.IconComponent, &menu.PathUrl, &menu.Indeks)
		if err != nil {
			panic(err)
		}

		menus = append(menus, menu)
	}

	return &menus
}

func (rpo *Repository) GetMenuChildren(ctx context.Context, tx *sql.Tx, menuId string) *[]model.MenuChild {
	query := "select id, name, path_url, indeks from menu_childs where menu_id = ? and deleted_at is null order by indeks"

	rows, err := tx.QueryContext(ctx, query, menuId)
	if err != nil {
		panic(err)
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}()

	var menus []model.MenuChild
	for rows.Next() {
		var menu model.MenuChild
		err = rows.Scan(&menu.Id, &menu.Name, &menu.PathUrl, &menu.Indeks)
		if err != nil {
			panic(err)
		}

		menus = append(menus, menu)
	}

	return &menus
}
