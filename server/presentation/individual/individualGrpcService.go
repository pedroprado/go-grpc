package individualGrpcService

import (
	"context"

	"pedro.prado.grpc.server.example/core/domain/entity"

	interfaces "pedro.prado.grpc.server.example/core/_interfaces"
	"pedro.prado.grpc.server.example/presentation/individual/proto"
)

type individualGRPCServer struct {
	proto.UnimplementedIndividualServiceServer
	individualService interfaces.IndividualService
}

func New(individualService interfaces.IndividualService) proto.IndividualServiceServer {
	return &individualGRPCServer{individualService: individualService}
}

func (ref *individualGRPCServer) GetIndividual(ctx context.Context, query *proto.GetQuery) (*proto.Individual, error) {
	individual, err := ref.individualService.GetIndividual(query.Id)
	if err != nil {
		return nil, err
	}

	if individual == nil {
		return &proto.Individual{}, nil
	}

	return &proto.Individual{
		Id:          individual.ID,
		Name:        individual.Name,
		DateOfBirth: individual.DateOfBirth,
		Nationality: individual.Nationalty,
	}, nil
}

func (ref *individualGRPCServer) CreateIndividual(ctx context.Context, request *proto.Individual) (*proto.Individual, error) {
	individual := entity.Individual{ID: request.Id, Name: request.Name, DateOfBirth: request.DateOfBirth, Nationalty: request.Nationality}

	created, err := ref.individualService.CreateIndividual(individual)
	if err != nil {
		return nil, err
	}

	return &proto.Individual{
		Id:          created.ID,
		Name:        created.Name,
		DateOfBirth: created.DateOfBirth,
		Nationality: created.Nationalty,
	}, nil
}

func (ref *individualGRPCServer) ListIndividuals(context.Context, *proto.ListQuery) (*proto.Individuals, error) {
	individuals, err := ref.individualService.ListIndividuals()
	if err != nil {
		return nil, err
	}

	responses := make([]*proto.Individual, len(individuals))

	for i, individual := range individuals {
		response := &proto.Individual{
			Id:          individual.ID,
			Name:        individual.Name,
			DateOfBirth: individual.DateOfBirth,
			Nationality: individual.Nationalty,
		}
		responses[i] = response
	}

	return &proto.Individuals{Individuals: responses}, nil
}
