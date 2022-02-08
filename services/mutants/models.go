package mutants

import (
	"context"

	"mutant-ms/models/mutants"
)

//go:generate mockery --name Services --filename mutants_services.go --outpkg mocks --structname Services --disable-version-string
type Services interface {
	IsMutant(ctx context.Context, dna []string) error
	ValidateDna(ctx context.Context, dna []string) error
	GetStats(ctx context.Context) (mutants.MutantStats, error)
}
