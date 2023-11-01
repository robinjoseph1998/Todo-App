package usecase

import (
	"Todo/pkg/models"
	repo "Todo/pkg/repository/interfaces"
	use "Todo/pkg/usecase/interfaces"
	"errors"
)

type usecase struct {
	Repo repo.RepoInterfaces
}

func NewUsecase(Repo repo.RepoInterfaces) use.UseCaseInterface {
	return &usecase{Repo}
}

func (uu *usecase) ExecuteAddName(request models.Test) (string, error) {
	SavedName, err := uu.Repo.CreateName(request)
	if err != nil {
		return "", errors.New("Can't added name")
	}
	return SavedName, nil
}

func (uu *usecase) ExecuteShowName() ([]models.Test, error) {
	names, err := uu.Repo.GetName()
	if err != nil {
		return nil, errors.New("can't get name")
	}
	return names, nil
}
