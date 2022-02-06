package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mutant-ms/services/mutants/mocks"
	testUtils "mutant-ms/utils/tests"
	"net/http"
	"testing"

	mutantModel "mutant-ms/models/mutants"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestIsMutant_ErrorOnBody(t *testing.T) {
	uri := "/mutant"

	payload := []int{1, 2, 3}

	payloadByte, err := json.Marshal(payload)
	assert.NoError(t, err)

	mutantServiceMock := new(mocks.Services)

	_, _, rec, c := testUtils.SetupServerTest(http.MethodPost, uri, bytes.NewReader(payloadByte))

	mutantController := NewMutantController()
	mutantController.services = mutantServiceMock
	mutantController.IsMutant(c)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestIsMutant_ErrorOnValidateDna(t *testing.T) {
	uri := "/mutant"

	payload := mutantModel.DnaSequence{
		Dna: []string{
			"ATGCGA",
			"CAGTGC",
			"TTATGT",
			"AGAAGG",
			"CCCCTA",
			"TCACSS",
		},
	}
	payloadByte, err := json.Marshal(payload)
	assert.NoError(t, err)

	mutantServiceMock := new(mocks.Services)
	mutantServiceMock.On("ValidateDna", mock.Anything, payload.Dna).Return(fmt.Errorf("test error"))

	_, _, rec, c := testUtils.SetupServerTest(http.MethodPost, uri, bytes.NewReader(payloadByte))

	mutantController := NewMutantController()
	mutantController.services = mutantServiceMock
	mutantController.IsMutant(c)

	var response mutantModel.IsMutantResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestIsMutant_OK(t *testing.T) {
	uri := "/mutant"

	payload := mutantModel.DnaSequence{
		Dna: []string{
			"ATGCGA",
			"CAGTGC",
			"TTATGT",
			"AGAAGG",
			"CCCCTA",
			"TCACTG",
		},
	}
	payloadByte, err := json.Marshal(payload)
	assert.NoError(t, err)

	mutantServiceMock := new(mocks.Services)
	mutantServiceMock.On("ValidateDna", mock.Anything, payload.Dna).Return(nil)
	mutantServiceMock.On("IsMutant", mock.Anything, payload.Dna).Return(true)

	_, _, rec, c := testUtils.SetupServerTest(http.MethodPost, uri, bytes.NewReader(payloadByte))

	mutantController := NewMutantController()
	mutantController.services = mutantServiceMock
	mutantController.IsMutant(c)

	var response mutantModel.IsMutantResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, true, response.IsMutant)
}
