package menu

import (
	"context"
	"database/sql"
	"errors"
	"github.com/oktapascal/go-simpro/model"
)

type Repository struct{}

func (rpo *Repository) GetMenus(ctx context.Context, tx *sql.Tx) *[]model.Menu {
	query := `select t1.id, t1.name, t1.icon_component, t1.path_url, t1.indeks, ifnull(t2.menu_count, 0) as menu_count
	from menus t1
	left join (
	    select menu_id, count(id) menu_count
	    from menu_childs
	    where deleted_at is null
	    group by menu_id
	) t2 on t1.id=t2.menu_id
	where t1.deleted_at is null 
	order by t1.indeks`

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
		err = rows.Scan(&menu.Id, &menu.Name, &menu.IconComponent, &menu.PathUrl, &menu.Indeks, &menu.Children)
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

func (rpo *Repository) FindMenuById(ctx context.Context, tx *sql.Tx, menuId string) (*model.Menu, error) {
	query := "select id, name, icon_component, path_url, indeks from menus where id = ? and deleted_at is null"

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

	menu := new(model.Menu)
	if rows.Next() {
		err = rows.Scan(&menu.Id, &menu.Name, &menu.IconComponent, &menu.PathUrl, &menu.Indeks)
		if err != nil {
			panic(err)
		}

		return menu, nil
	} else {
		return nil, errors.New("menu not found")
	}
}
