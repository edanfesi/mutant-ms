package mutants

import (
	"context"
	mutantModel "mutant-ms/models/mutants"
)

//go:generate mockery --name Repositories --filename mutants_repositories.go --outpkg mocks --structname Repositories --disable-version-string
type Repositories interface {
	Save(ctx context.Context, mutantDna mutantModel.MutantDna) error
	GetAll(ctx context.Context) ([]mutantModel.MutantDna, error)
}
