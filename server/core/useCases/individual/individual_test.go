package individualService

import (
	"fmt"
	"reflect"
	"testing"

	"pedro.prado.grpc.server.example/core/domain/entity"
)

func TestGet(t *testing.T) {
	individuals := map[string]entity.Individual{
		"1": entity.Individual{ID: "1", Name: "John", DateOfBirth: "11/01/1991", Nationalty: "USA"},
		"2": entity.Individual{ID: "2", Name: "Doe", DateOfBirth: "12/09/1986", Nationalty: "UK"},
		"3": entity.Individual{ID: "3", Name: "Jorge", DateOfBirth: "27/02/1980", Nationalty: "BR"},
		"4": entity.Individual{ID: "4", Name: "Takashi", DateOfBirth: "30/06/2000", Nationalty: "JP"},
		"5": entity.Individual{ID: "5", Name: "Joseph", DateOfBirth: "17/04/1997", Nationalty: "FR"},
	}

	service := New(individuals)

	individual, _ := service.GetIndividual("2")
	fmt.Println(individual)

	if !reflect.DeepEqual(individuals["2"], individual) {
		fmt.Errorf("Not equals")
	}
}

func TestList(t *testing.T) {
	individuals := map[string]entity.Individual{
		"1": entity.Individual{ID: "1", Name: "John", DateOfBirth: "11/01/1991", Nationalty: "USA"},
		"2": entity.Individual{ID: "2", Name: "Doe", DateOfBirth: "12/09/1986", Nationalty: "UK"},
		"3": entity.Individual{ID: "3", Name: "Jorge", DateOfBirth: "27/02/1980", Nationalty: "BR"},
		"4": entity.Individual{ID: "4", Name: "Takashi", DateOfBirth: "30/06/2000", Nationalty: "JP"},
		"5": entity.Individual{ID: "5", Name: "Joseph", DateOfBirth: "17/04/1997", Nationalty: "FR"},
	}

	service := New(individuals)

	indivs, _ := service.ListIndividuals()

	fmt.Println(indivs)
}
