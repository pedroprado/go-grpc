package individualAdapter

import (
	"context"

	protoIndividual "pedro.prado.grpc.client.example/adapter/individual/proto"
	"pedro.prado.grpc.client.example/core/domain/entity"
)

type IndividualAdapter interface {
	GetIndividual(id string) (*entity.Individual, error)
	CreateIndividual(individual *entity.Individual) (*entity.Individual, error)
	ListIndividuals() ([]entity.Individual, error)
}

type individualAdapter struct {
	ctx    context.Context
	client protoIndividual.IndividualServiceClient
}

func New(ctx context.Context, client protoIndividual.IndividualServiceClient) IndividualAdapter {
	return &individualAdapter{
		ctx:    ctx,
		client: client,
	}
}

func (ref *individualAdapter) GetIndividual(id string) (*entity.Individual, error) {

	response, err := ref.client.GetIndividual(ref.ctx, &protoIndividual.GetQuery{Id: id})
	if err != nil {
		return nil, err
	}

	return &entity.Individual{
		ID:          response.Id,
		Name:        response.Name,
		DateOfBirth: response.DateOfBirth,
		Nationality: response.Nationality,
	}, nil
}

func (ref *individualAdapter) CreateIndividual(individual *entity.Individual) (*entity.Individual, error) {
	return nil, nil
}

func (ref *individualAdapter) ListIndividuals() ([]entity.Individual, error) {

	responses, err := ref.client.ListIndividuals(ref.ctx, &protoIndividual.ListQuery{Id: "1"})
	if err != nil {
		return nil, err
	}

	individuals := make([]entity.Individual, len(responses.Individuals))

	for i, response := range responses.Individuals {
		individual := entity.Individual{
			ID:          response.Id,
			Name:        response.Name,
			DateOfBirth: response.DateOfBirth,
			Nationality: response.Nationality,
		}
		individuals[i] = individual
	}

	return individuals, nil
}
