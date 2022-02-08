package mutants

import (
	"context"
	"fmt"

	"github.com/kisielk/sqlstruct"

	mutantModel "mutant-ms/models/mutants"
	"mutant-ms/repositories/builders"
	"mutant-ms/storage"
	uPostgres "mutant-ms/utils/postgres"
)

type MutantsPostgres struct {
	db storage.PostgresDrivers
}

func NewMutantsPostgres(db storage.PostgresDrivers) *MutantsPostgres {
	return &MutantsPostgres{
		db: db,
	}
}

func (mRepositories *MutantsPostgres) Save(ctx context.Context, mutantDna mutantModel.MutantDna) error {
	_, err := mRepositories.db.ExecContext(ctx, builders.SaveDna(mutantDna))
	fmt.Println(err)
	if err != nil {
		return fmt.Errorf("[mutant_postgres][save][err:%v]", err.Error())
	}

	return nil
}

func (mRepositories *MutantsPostgres) GetAll(ctx context.Context) ([]mutantModel.MutantDna, error) {
	rows, err := mRepositories.db.QueryContext(ctx, builders.GetDNAs())
	if err != nil {
		return make([]mutantModel.MutantDna, 0), fmt.Errorf("[mutant_postgres][get][err:%v]", err.Error())
	}

	defer uPostgres.CloseRows(rows)

	mutantsDNA := make([]mutantModel.MutantDna, 0)
	row := mutantModel.MutantDna{}
	for rows.Next() {
		err = sqlstruct.Scan(&row, rows)
		fmt.Println(rows.Columns())
		if err != nil {
			return make([]mutantModel.MutantDna, 0), fmt.Errorf("[mutant_postgres][scan][err:%v]", err.Error())
		}
		mutantsDNA = append(mutantsDNA, row)
	}

	return mutantsDNA, nil
}
