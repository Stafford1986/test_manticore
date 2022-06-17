package manticore_native

import (
	"context"
	"errors"
	"fmt"
	"github.com/Stafford1986/test_manticore/pb"
	"github.com/Stafford1986/test_manticore/usecase/entity"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strconv"
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
	ctxWithCancel, cancelFn := context.WithCancel(ctx)
	defer cancelFn()
	q, err := req.BuildInsertQuery()
	if err != nil {
		return err
	}
	_, err = repo.db.QueryContext(ctxWithCancel, q)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1064 && strings.Contains(mysqlErr.Message, "duplicate") {
			return entity.ErrDuplicateId
		}
		return err
	}

	return nil
}

func (repo *VacancyRepository) BulkInsert(ctx context.Context, req *pb.VacancyList) error {
	ctxWithCancel, cancelFn := context.WithCancel(ctx)
	defer cancelFn()
	q, err := req.BuildBulkInsertQuery()
	if err != nil {
		return err
	}
	_, err = repo.db.QueryContext(ctxWithCancel, q)
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

func (repo *VacancyRepository) Find(ctx context.Context, sr *pb.VacancySearchEntity) ([]uint32, error) {
	fr, err := sr.BuildSearchQuery(1000)
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

	var result []uint32

	for {
		if rows.Next() {
			fieldMap := make(map[string]interface{}, len(col))

			err = sqlx.MapScan(rows, fieldMap)
			if err != nil {
				return nil, err
			}

			val := fieldMap["id"]
			v, ok := val.([]byte)
			if !ok {
				return nil, errors.New("err convert id")
			}
			p, err := strconv.ParseUint(string(v), 10, 32)
			if err != nil {
				return nil, errors.New("err convert value to Id")
			}

			result = append(result, uint32(p))

			continue
		}

		break
	}

	return result, nil
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
