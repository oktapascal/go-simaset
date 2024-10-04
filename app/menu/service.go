package menu

import (
	"github.com/oktapascal/go-simpro/model"
)

type Service struct {
	rpo model.MenuRepository
}

func (svc *Service) GetMenu(group string) *[]model.Menu {
	menus := svc.rpo.GetMenu(group)

	return menus
}
