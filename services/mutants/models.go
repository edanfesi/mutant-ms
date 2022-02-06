package mutants

import "context"

//go:generate mockery --name Services --filename mutants_services.go --outpkg mocks --structname Services --disable-version-string
type Services interface {
	IsMutant(ctx context.Context, dna []string) bool
	ValidateDna(ctx context.Context, dna []string) error
}
