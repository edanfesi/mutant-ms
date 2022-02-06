package mutants

import (
	"context"
	"testing"

	mutantContext "mutant-ms/utils/context"
	"mutant-ms/utils/logger"

	"github.com/stretchr/testify/assert"
)

func TestIsMutant_ReturnTrue(t *testing.T) {
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

	result := service.IsMutant(ctx, dna)

	assert.Equal(t, true, result)
}

func TestIsMutant_ReturnFalse(t *testing.T) {
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

	result := service.IsMutant(ctx, dna)

	assert.Equal(t, false, result)
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
