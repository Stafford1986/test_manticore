package manticore_native

import (
	"context"
	"errors"
	"fmt"
	"github.com/Stafford1986/test_manticore/pb"
	"github.com/Stafford1986/test_manticore/usecase/entity"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strconv"
	"strings"
)

const (
	resumeIndex = "resumes"
)

type ResumeRepository struct {
	db *sqlx.DB
}

func NewResumeRepository(db *sqlx.DB) *ResumeRepository {
	return &ResumeRepository{
		db: db,
	}
}

func (repo *ResumeRepository) Insert(ctx context.Context, req *pb.ResumeEntity) error {
	q, err := req.BuildInsertQuery()
	fmt.Println(q)
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

func (repo *ResumeRepository) Upsert(ctx context.Context, req *pb.ResumeEntity) error {
	q, err := req.BuildUpsertQuery()
	if err != nil {
		return err
	}
	_, err = repo.db.QueryContext(ctx, q)
	if err != nil {
		return err
	}

	oq := createOptimizeQuery(resumeIndex)
	_, err = repo.db.QueryContext(ctx, oq)
	if err != nil {
		return err
	}

	return nil
}

func (repo *ResumeRepository) Find(ctx context.Context, sr *pb.ResumeSearchEntity) ([]uint32, error) {
	fr, err := sr.BuildSearchQuery(1000)
	fmt.Println(fr)
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

func (repo *ResumeRepository) GetSuggestions(ctx context.Context, req string) ([]string, error) {
	rows, err := repo.db.QueryContext(ctx, fmt.Sprintf("call suggest('%s', '%s')", req, resumeIndex))
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

func createOptimizeQuery(index string) string {
	return fmt.Sprintf("OPTIMIZE INDEX %s;", index)
}
