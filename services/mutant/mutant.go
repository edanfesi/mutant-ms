package mutant

import (
	"context"
	"fmt"
	"regexp"
)

type MutantServices struct{}

func (mutantServices *MutantServices) IsMutant(ctx context.Context, dna []string) (bool, error) {
	dnaLen := len(dna)

	visitedBases := make([][]int, dnaLen)
	for i := range visitedBases {
		visitedBases[i] = make([]int, dnaLen)
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
				return true, nil
			}
		}
	}

	return false, nil
}

func (mutantservices *MutantServices) ValidateDna(ctx context.Context, dna []string) error {
	dnaLen := len(dna)

	r, _ := regexp.Compile(fmt.Sprintf("[A,T,C,G]{%d}", dnaLen))

	for i := 0; i < dnaLen; i++ {
		if !r.MatchString(dna[i]) {
			return fmt.Errorf("the base %s is not a valid base", dna[i])
		}
	}

	return nil
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
