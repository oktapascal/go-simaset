package menu

import (
	"context"
	"database/sql"
	"github.com/oktapascal/go-simpro/helper"
	"github.com/oktapascal/go-simpro/model"
)

type Service struct {
	rpo model.MenuRepository
	db  *sql.DB
}

func (svc *Service) GetMenus(ctx context.Context) []model.MenuResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	menus := svc.rpo.GetMenus(ctx, tx)

	var result []model.MenuResponse
	if len(*menus) > 0 {
		for _, value := range *menus {
			menu := model.MenuResponse{
				Id:            value.Id,
				Name:          value.Name,
				IconComponent: value.IconComponent,
				PathUrl:       value.PathUrl,
				Indeks:        value.Indeks,
			}

			result = append(result, menu)
		}
	}

	return result
}

func (svc *Service) GetMenuChildren(ctx context.Context, menuId string) []model.MenuChildResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	menus := svc.rpo.GetMenuChildren(ctx, tx, menuId)

	var result []model.MenuChildResponse
	if len(*menus) > 0 {
		for _, value := range *menus {
			menu := model.MenuChildResponse{
				Id:      value.Id,
				Name:    value.Name,
				PathUrl: value.PathUrl,
				Indeks:  value.Indeks,
			}

			result = append(result, menu)
		}
	}

	return result
}
