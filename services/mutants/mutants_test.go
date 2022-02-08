package mutants

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"mutant-ms/models/mutants"
	"mutant-ms/repositories/mutants/mocks"
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

	dnaString := strings.Join(dna, ", ")

	toSave := mutants.MutantDna{
		ID:       0,
		Dna:      dnaString,
		IsMutant: true,
	}

	mutantRepoMock := new(mocks.Repositories)
	mutantRepoMock.On("Save", mock.Anything, toSave).Return(nil)

	service := &MutantsServices{
		repositories: mutantRepoMock,
	}

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

	dnaString := strings.Join(dna, ", ")

	toSave := mutants.MutantDna{
		ID:       0,
		Dna:      dnaString,
		IsMutant: false,
	}

	mutantRepoMock := new(mocks.Repositories)
	mutantRepoMock.On("Save", mock.Anything, toSave).Return(nil)

	service := &MutantsServices{
		repositories: mutantRepoMock,
	}

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

	mutantRepoMock := new(mocks.Repositories)

	service := &MutantsServices{
		repositories: mutantRepoMock,
	}

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

	mutantRepoMock := new(mocks.Repositories)

	service := &MutantsServices{
		repositories: mutantRepoMock,
	}

	result := service.ValidateDna(ctx, dna)

	assert.Error(t, result)
}
