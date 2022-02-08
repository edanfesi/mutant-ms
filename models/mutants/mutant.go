package mutants

type DnaSequence struct {
	Dna []string `json:"dna"`
}

type IsMutantResponse struct {
	IsMutant bool `json:"is_mutant"`
}

type MutantDna struct {
	ID       int    `json:"id"`
	Dna      string `json:"dna"`
	IsMutant bool   `json:"is_mutant" sql:"is_mutant"`
}

type MutantStats struct {
	CountMutantDna int     `json:"count_mutant_dna"`
	CountHumanDna  int     `json:"count_human_dna"`
	Ratio          float32 `json:"ratio"`
}
