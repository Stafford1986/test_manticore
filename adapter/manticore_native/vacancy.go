package manticore_native

import (
	"context"
	"errors"
	"fmt"
	"github.com/Stafford1986/test_manticore/pb"
	"github.com/Stafford1986/test_manticore/usecase/entity"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strings"
)

const (
	vacancyIndex = "vacancies"
)

type VacancyRepository struct {
	db *sqlx.DB
}

func NewVacancyRepository(db *sqlx.DB) *VacancyRepository {
	return &VacancyRepository{
		db: db,
	}
}

func (repo *VacancyRepository) Insert(ctx context.Context, req *pb.VacancyEntity) error {
	q, err := req.BuildInsertQuery()
	if err != nil {
		return err
	}
	_, err = repo.db.QueryContext(ctx, q)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1064 && strings.Contains(mysqlErr.Message, "duplicate") {
			return entity.ErrDuplicateId
		}
		return err
	}

	return nil
}

func (repo *VacancyRepository) Upsert(ctx context.Context, req *pb.VacancyEntity) error {
	q, err := req.BuildUpsertQuery()
	if err != nil {
		return err
	}
	_, err = repo.db.QueryContext(ctx, q)
	if err != nil {
		return err
	}

	oq := createOptimizeQuery(vacancyIndex)
	_, err = repo.db.QueryContext(ctx, oq)
	if err != nil {
		return err
	}

	return nil
}

func (repo *VacancyRepository) Find(ctx context.Context, sr *pb.VacancySearchEntity) (*pb.VacancySearchResponse, error) {
	fr, err := sr.BuildSearchQuery()
	if err != nil {
		return nil, err
	}

	rows, err := repo.db.QueryContext(ctx, fr)
	if err != nil {
		return nil, err
	}

	col, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var result []*pb.VacancyEntity

	for {
		if rows.Next() {
			vc := &pb.VacancyEntity{}
			fieldMap := make(map[string]interface{}, len(col))

			err = sqlx.MapScan(rows, fieldMap)
			if err != nil {
				return nil, err
			}

			vac, err := vc.ParseDbResult(fieldMap)
			if err != nil {
				return nil, err
			}

			result = append(result, vac)

			continue
		}

		break
	}

	return &pb.VacancySearchResponse{
		Items: result,
	}, nil
}

func (repo *VacancyRepository) GetSuggestions(ctx context.Context, req string) ([]string, error) {
	rows, err := repo.db.QueryContext(ctx, fmt.Sprintf("call suggest('%s', '%s')", req, vacancyIndex))
	if err != nil {
		return nil, err
	}

	fmt.Println(rows)

	var result []string

	for {
		if rows.Next() {
			var (
				suggest  string
				distance int
				docs     int
			)
			err := rows.Scan(&suggest, &distance, &docs)
			if err != nil {
				return nil, err
			}

			result = append(result, suggest)

			continue
		}

		break
	}

	return result, err
}