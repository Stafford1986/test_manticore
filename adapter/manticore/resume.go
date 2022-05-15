package manticore

import (
	"context"
	"errors"
	"fmt"
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

func (repo *ResumeRepository) Insert(ctx context.Context, req *entity.ResumeEntity) error {
	q, err := req.BuildQuery(entity.QueryTypeInsert)
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

func (repo *ResumeRepository) Upsert(ctx context.Context, req *entity.ResumeEntity) error {
	q, err := req.BuildQuery(entity.QueryTypeUpsert)
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

func (repo *ResumeRepository) Find(ctx context.Context, sr *entity.SearchRequest) ([]*entity.ResumeEntity, error) {
	err := sr.Validate()
	if err != nil {
		return nil, err
	}

	fr, err := createFindQuery(sr, resumeIndex)
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

	var result []*entity.ResumeEntity

	for {
		if rows.Next() {
			fieldMap := make(map[string]interface{}, len(col))

			err = sqlx.MapScan(rows, fieldMap)
			if err != nil {
				return nil, err
			}

			r, err := scanMapIntoResume(fieldMap)
			if err != nil {
				return nil, err
			}

			result = append(result, r)

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

func createFindQuery(sr *entity.SearchRequest, index string) (string, error) {
	sb := strings.Builder{}
	isActive := func() int {
		if sr.IsActive {
			return 1
		}

		return 0
	}()
	_, err := sb.WriteString(fmt.Sprintf(`SELECT * FROM %s WHERE MATCH('*%s*') AND active = %d`,
		index, sr.Query, isActive))
	if err != nil {
		return "", err
	}

	if sr.CityId != 0 {
		_, err = sb.WriteString(fmt.Sprintf(" AND city = %d", sr.CityId))
		if err != nil {
			return "", err
		}
	}

	if sr.SortParams == nil {
		_, err = sb.WriteString(fmt.Sprintf(" ORDER BY %s DESC, city_level ASC", entity.SortUpdatedAt))
		if err != nil {
			return "", err
		}
	} else {
		_, err = sb.WriteString(fmt.Sprintf(" ORDER BY %s %s, city_level ASC", sr.SortParams.Field, sr.SortParams.Order))
		if err != nil {
			return "", err
		}
	}

	if sr.FilterParams != nil {
		_, err = sb.WriteString(fmt.Sprintf(",%s %s;", sr.FilterParams.Field, sr.FilterParams.Order))
		if err != nil {
			return "", err
		}

		return sb.String(), nil
	}

	_, err = sb.WriteString(";")
	if err != nil {
		return "", err
	}

	return sb.String(), nil
}

func createOptimizeQuery(index string) string {
	return fmt.Sprintf("OPTIMIZE INDEX %s;", index)
}

func scanMapIntoResume(m map[string]interface{}) (*entity.ResumeEntity, error) {
	res := &entity.ResumeEntity{}
	for k, v := range m {
		switch k {
		case "id":
			id, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err. convert id")
			}

			p, err := strconv.ParseInt(string(id), 10, 64)
			if err != nil {
				return nil, err
			}
			res.Id = p
		case "name":
			name, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err. convert name")
			}
			res.Name = string(name)
		case "active":
			active, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err. convert active")
			}
			a, err := strconv.Atoi(string(active))
			if err != nil {
				return nil, err
			}
			res.Active = a
		case "salary_max":
			sf, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err. convert salary max")
			}

			s, err := strconv.ParseInt(string(sf), 10, 32)
			if err != nil {
				return nil, err
			}
			res.SalaryMax = int32(s)
		case "salary_min":
			st, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err. convert salary min")
			}

			s, err := strconv.ParseInt(string(st), 10, 32)
			if err != nil {
				return nil, err
			}
			res.SalaryMin = int32(s)
		case "created_at":
			ca, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err. convert created at")
			}

			s, err := strconv.ParseInt(string(ca), 10, 32)
			if err != nil {
				return nil, err
			}
			res.CreatedAT = int32(s)
		case "updated_at":
			ca, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err. convert updated_at")
			}

			s, err := strconv.ParseInt(string(ca), 10, 32)
			if err != nil {
				return nil, err
			}
			res.UpdatedAt = int32(s)
		case "skills":
			skills, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err. convert skills")
			}
			res.Skills = string(skills)
		case "city":
			ct, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err. convert city")
			}

			s, err := strconv.ParseInt(string(ct), 10, 32)
			if err != nil {
				return nil, err
			}
			res.City = int32(s)
		case "city_level":
			ct, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err. convert city level")
			}

			s, err := strconv.ParseInt(string(ct), 10, 32)
			if err != nil {
				return nil, err
			}
			res.CityLevel = int32(s)
		default:
			return nil, fmt.Errorf("err. wrong field: %s", k)
		}

	}

	return res, nil
}
