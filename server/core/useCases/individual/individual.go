package individualService

import (
	"errors"

	interfaces "pedro.prado.grpc.server.example/core/_interfaces"
	"pedro.prado.grpc.server.example/core/domain/entity"
)

type individualService struct {
	individuals map[string]entity.Individual
}

func New(individuals map[string]entity.Individual) interfaces.IndividualService {

	if individuals == nil {
		individuals = map[string]entity.Individual{
			"1": entity.Individual{ID: "1", Name: "John", DateOfBirth: "11/01/1991", Nationalty: "USA"},
			"2": entity.Individual{ID: "2", Name: "Doe", DateOfBirth: "12/09/1986", Nationalty: "UK"},
			"3": entity.Individual{ID: "3", Name: "Jorge", DateOfBirth: "27/02/1980", Nationalty: "BR"},
			"4": entity.Individual{ID: "4", Name: "Takashi", DateOfBirth: "30/06/2000", Nationalty: "JP"},
			"5": entity.Individual{ID: "5", Name: "Joseph", DateOfBirth: "17/04/1997", Nationalty: "FR"},
		}
	}

	return &individualService{individuals: individuals}
}

func (ref *individualService) GetIndividual(id string) (*entity.Individual, error) {
	individual, exists := ref.individuals[id]
	if exists {
		return &individual, nil
	}

	return nil, nil
}

func (ref *individualService) CreateIndividual(individual entity.Individual) (*entity.Individual, error) {
	_, exists := ref.individuals[individual.ID]
	if exists {
		return nil, errors.New("individual already exists")
	}

	ref.individuals[individual.ID] = individual

	return &individual, nil
}

func (ref *individualService) ListIndividuals() ([]entity.Individual, error) {
	individuals := make([]entity.Individual, len(ref.individuals))

	i := 0
	for _, individual := range ref.individuals {
		individuals[i] = individual
		i++
	}

	return individuals, nil
}
