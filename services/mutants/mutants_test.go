package mutants

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	mutantContext "mutant-ms/utils/context"
	"mutant-ms/utils/logger"
)

func TestIsMutant_IsMutant(t *testing.T) {
	ctx := mutantContext.SetLogger(context.Background(), logger.New("-"))

	dna := []string{
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTG",
	}

	service := &MutantsServices{}

	err := service.IsMutant(ctx, dna)

	assert.NoError(t, err)
}

func TestIsMutant_IsNotMutant(t *testing.T) {
	ctx := mutantContext.SetLogger(context.Background(), logger.New("-"))

	dna := []string{
		"ATGCCA",
		"CAGTGC",
		"TTATGT",
		"AGACCG",
		"CCCCTA",
		"TCACTG",
	}

	service := &MutantsServices{}

	err := service.IsMutant(ctx, dna)

	assert.Error(t, err)
}

func TestValidateDna_OK(t *testing.T) {
	ctx := mutantContext.SetLogger(context.Background(), logger.New("-"))

	dna := []string{
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTG",
	}

	service := &MutantsServices{}

	result := service.ValidateDna(ctx, dna)

	assert.NoError(t, result)
}

func TestValidateDna_ReturnError(t *testing.T) {
	ctx := mutantContext.SetLogger(context.Background(), logger.New("-"))

	dna := []string{
		"ATGCGA",
		"CAGDDC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTG",
	}

	service := &MutantsServices{}

	result := service.ValidateDna(ctx, dna)

	assert.Error(t, result)
}
