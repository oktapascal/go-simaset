package menu

import (
	"encoding/json"
	"github.com/oktapascal/go-simpro/model"
	"io"
	"os"
)

type Repository struct{}

func (rpo *Repository) GetMenu(group string) *[]model.Menu {
	jsonFile, err := os.Open("storage/json/" + group + ".json")
	if err != nil {
		panic(err.Error())
	}

	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			panic(err.Error())
		}
	}(jsonFile)

	bytesValue, errBytes := io.ReadAll(jsonFile)
	if errBytes != nil {
		panic(errBytes.Error())
	}

	var menus []model.Menu
	err = json.Unmarshal(bytesValue, &menus)
	if err != nil {
		panic(err.Error())
	}

	return &menus
}
