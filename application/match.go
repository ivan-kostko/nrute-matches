package application

import "github.com/ivan-kostko/nrute-matches/domain"

type Match struct {
	Movements         []Movement
	ContractCondition *domain.ContractCondition
	IsApproved        bool
	Score             int
}
