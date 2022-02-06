package mutant

import "context"

type Services interface {
	IsMutant(ctx context.Context, dna []string) (bool, error)
	ValidateDna(ctx context.Context, dna []string) error
}
