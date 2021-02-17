package _interfaces

import "pedro.prado.grpc.server.example/core/domain/entity"

type IndividualService interface {
	GetIndividual(id string) (*entity.Individual, error)
	CreateIndividual(individual entity.Individual) (*entity.Individual, error)
	ListIndividuals() ([]entity.Individual, error)
}
