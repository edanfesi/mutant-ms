package builders

import (
	"fmt"

	mutantModel "mutant-ms/models/mutants"
)

func SaveDna(mutantDna mutantModel.MutantDna) string {
	query := `
	INSERT
	INTO
		mutant_dna (dna, is_mutant)
	VALUES ('%s', %v)
	ON CONFLICT DO NOTHING
	`

	return fmt.Sprintf(query, mutantDna.Dna, mutantDna.IsMutant)
}

func GetDNAs() string {
	return `
	SELECT
		id,
		dna,
		is_mutant
	FROM
		mutant_dna
	`
}
