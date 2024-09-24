package menu

import (
	"context"
	"database/sql"
	"github.com/oktapascal/go-simpro/model"
)

type Service struct {
	rpo model.MenuRepository
	db  *sql.DB
}

func (svc *Service) GetMenus(ctx context.Context) []model.MenuResponse {
	//TODO implement me
	panic("implement me")
}

func (svc *Service) GetMenuChildren(ctx context.Context, menuId string) []model.MenuChildResponse {
	//TODO implement me
	panic("implement me")
}
