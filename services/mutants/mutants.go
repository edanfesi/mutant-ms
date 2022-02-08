package mutants

import (
	"context"
	"fmt"
	"math"
	"regexp"
	"strings"

	"mutant-ms/models/mutants"
	mutantsRepositories "mutant-ms/repositories/mutants"
)

type MutantsServices struct {
	repositories mutantsRepositories.Repositories
}

func NewMutantService(repositories mutantsRepositories.Repositories) *MutantsServices {
	return &MutantsServices{
		repositories: repositories,
	}
}

func (mutantsServices *MutantsServices) IsMutant(ctx context.Context, dna []string) error {
	dnaLen := len(dna)

	visitedBases := make([][]int, dnaLen)
	for i := range visitedBases {
		visitedBases[i] = make([]int, dnaLen)
	}

	mutantDNAToSave := mutants.MutantDna{
		Dna:      strings.Join(dna, ", "),
		IsMutant: false,
	}

	mutantDNA := 0
	for i := 0; i < dnaLen; i++ {
		for j := 0; j < dnaLen; j++ {
			nextPosX := i + 1
			nextPosY := j + 1
			if i+1 < dnaLen && dna[i+1][j] != 1 {
				horizontalDNALen := 1 + getLengthDNA(dna[i][j], nextPosX, j, 1, 0, dna, visitedBases)

				if horizontalDNALen >= 4 {
					mutantDNA += 1
				}
			}

			if j+1 < dnaLen && dna[i][j+1] != 1 {
				verticalDNALength := 1 + getLengthDNA(dna[i][j], i, nextPosY, 0, 1, dna, visitedBases)

				if verticalDNALength >= 4 {
					mutantDNA += 1
				}
			}

			if i+1 < dnaLen && j+1 < dnaLen && dna[i+1][j+1] != 1 {
				diagonalDNALength := 1 + getLengthDNA(
					dna[i][j], nextPosX, nextPosY, 1, 1, dna, visitedBases)

				if diagonalDNALength >= 4 {
					mutantDNA += 1
				}
			}

			visitedBases[i][j] = 1

			if mutantDNA > 1 {
				mutantDNAToSave.IsMutant = true
				err := mutantsServices.repositories.Save(ctx, mutantDNAToSave)
				if err != nil {
					return err
				}

				return nil
			}
		}
	}

	err := mutantsServices.repositories.Save(ctx, mutantDNAToSave)
	if err != nil {
		return err
	}

	return fmt.Errorf("forbidden")
}

func (mutantsservices *MutantsServices) ValidateDna(ctx context.Context, dna []string) error {
	dnaLen := len(dna)

	r, _ := regexp.Compile(fmt.Sprintf("[A,T,C,G]{%d}", dnaLen))

	for i := 0; i < dnaLen; i++ {
		if !r.MatchString(dna[i]) {
			return fmt.Errorf("the base %s is not a valid base", dna[i])
		}
	}

	return nil
}

func (mutantsServices *MutantsServices) GetStats(ctx context.Context) (mutants.MutantStats, error) {
	savedDna, err := mutantsServices.repositories.GetAll(ctx)
	if err != nil {
		return mutants.MutantStats{}, err
	}

	mutantsDna := getMutantDna(savedDna)

	countMutantDna := len(mutantsDna)
	countHumanDna := len(savedDna)

	if countHumanDna == 0 {
		countHumanDna = 1
	}

	ratio := float32(countMutantDna) / float32(countHumanDna)
	mutantStats := mutants.MutantStats{
		CountMutantDna: countMutantDna,
		CountHumanDna:  len(savedDna),
		Ratio:          float32(math.Floor(float64(ratio*100)) / 100),
	}

	return mutantStats, nil
}

func getLengthDNA(currentValue byte, posX int, posY int, movX int, movY int, dna []string, visited [][]int) int {
	if currentValue != dna[posX][posY] {
		return 0
	}

	dnaLen := len(dna)

	newPosX := posX + movX
	newPosY := posY + movY

	visited[posX][posY] = 1

	if newPosX >= dnaLen || newPosY >= dnaLen {
		return 1
	}

	return 1 + getLengthDNA(dna[posX][posY], posX+movX, posY+movY, movX, movY, dna, visited)
}

func getMutantDna(humansDna []mutants.MutantDna) []mutants.MutantDna {
	mutantsAdn := make([]mutants.MutantDna, 0)

	for _, val := range humansDna {
		if val.IsMutant {
			mutantsAdn = append(mutantsAdn, val)
		}
	}

	return mutantsAdn
}
