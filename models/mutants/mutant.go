package mutants

type DnaSequence struct {
	Dna []string `json:"dna"`
}

type IsMutantResponse struct {
	IsMutant bool `json:"is_mutant"`
}
