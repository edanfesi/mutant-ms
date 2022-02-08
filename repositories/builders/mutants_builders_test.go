package builders

import (
	"testing"

	"github.com/stretchr/testify/assert"

	mutantModel "mutant-ms/models/mutants"
)

func TestSaveDnaOk(t *testing.T) {
	mutantDna := mutantModel.MutantDna{
		ID:       1,
		Dna:      "AAAAAA, BBBBBB",
		IsMutant: true,
	}

	expected := `
	INSERT
	INTO
		mutant_dna (dna, is_mutant)
	VALUES ('AAAAAA, BBBBBB', true)
	ON CONFLICT DO NOTHING
	`

	query := SaveDna(mutantDna)
	assert.Equal(t, expected, query)
}

func TestGetDNAs(t *testing.T) {
	expected := `
	SELECT
		id,
		dna,
		is_mutant
	FROM
		mutant_dna
	`

	query := GetDNAs()
	assert.Equal(t, expected, query)
}
