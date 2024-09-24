package menu

import (
	"context"
	"database/sql"
	"github.com/oktapascal/go-simpro/model"
)

type Repository struct{}

func (rpo *Repository) GetMenus(ctx context.Context, tx *sql.Tx) *[]model.Menu {
	//TODO implement me
	panic("implement me")
}

func (rpo *Repository) GetMenuChildren(ctx context.Context, tx *sql.Tx, menuId string) *[]model.MenuChild {
	//TODO implement me
	panic("implement me")
}
