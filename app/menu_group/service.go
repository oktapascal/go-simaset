package menu_group

import (
	"context"
	"database/sql"
	"github.com/oktapascal/go-simpro/exception"
	"github.com/oktapascal/go-simpro/helper"
	"github.com/oktapascal/go-simpro/model"
)

type Service struct {
	rpo model.MenuGroupRepository
	db  *sql.DB
}

func (svc *Service) StoreMenuGroup(ctx context.Context, request *model.MenuGroupRequest) model.MenuGroupResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	menuGroup := new(model.MenuGroup)
	menuGroup.Name = request.Name

	menuGroup = svc.rpo.CreateMenuGroup(ctx, tx, menuGroup)

	return model.MenuGroupResponse{
		Id:   menuGroup.Id,
		Name: menuGroup.Name,
	}
}

func (svc *Service) UpdateMenuGroup(ctx context.Context, request *model.MenuGroupRequest) model.MenuGroupResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	menuGroup, errMenuGroup := svc.rpo.GetMenuGroup(ctx, tx, request.Id)
	if errMenuGroup != nil {
		panic(exception.NewNotFoundError(errMenuGroup.Error()))
	}

	menuGroup.Name = request.Name

	menuGroup = svc.rpo.UpdateMenuGroup(ctx, tx, menuGroup)

	return model.MenuGroupResponse{
		Id:   menuGroup.Id,
		Name: menuGroup.Name,
	}
}

func (svc *Service) GetAllMenuGroups(ctx context.Context) []model.MenuGroupResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	menuGroups := svc.rpo.GetAllMenuGroups(ctx, tx)

	var result []model.MenuGroupResponse
	if len(*menuGroups) > 0 {
		for _, value := range *menuGroups {
			menuGroup := model.MenuGroupResponse{
				Id:   value.Id,
				Name: value.Name,
			}

			result = append(result, menuGroup)
		}
	}

	return result
}

func (svc *Service) GetOneMenuGroup(ctx context.Context, id string) model.MenuGroupResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	groupMenu, errMenuGroup := svc.rpo.GetMenuGroup(ctx, tx, id)
	if errMenuGroup != nil {
		panic(exception.NewNotFoundError(errMenuGroup.Error()))
	}

	return model.MenuGroupResponse{
		Id:   groupMenu.Id,
		Name: groupMenu.Name,
	}
}

func (svc *Service) DeleteMenuGroup(ctx context.Context, id string) {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	_, errMenuGroup := svc.rpo.GetMenuGroup(ctx, tx, id)
	if errMenuGroup != nil {
		panic(exception.NewNotFoundError(errMenuGroup.Error()))
	}

	svc.rpo.DeleteMenuGroup(ctx, tx, id)
}
