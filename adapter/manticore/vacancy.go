package manticore

import (
	"context"
	"errors"
	"fmt"
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

func (repo *VacancyRepository) Insert(ctx context.Context, req *entity.VacancyEntity) error {
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

func (repo *VacancyRepository) Upsert(ctx context.Context, req *entity.VacancyEntity) error {
	q, err := req.BuildQuery(entity.QueryTypeUpsert)
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

func (repo *VacancyRepository) Find(ctx context.Context, sr *entity.SearchRequest) ([]*entity.VacancyEntity, error) {
	err := sr.Validate()
	if err != nil {
		return nil, err
	}

	fr, err := createFindQuery(sr, vacancyIndex)
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

	var result []*entity.VacancyEntity

	for {
		if rows.Next() {
			fieldMap := make(map[string]interface{}, len(col))

			err = sqlx.MapScan(rows, fieldMap)
			if err != nil {
				return nil, err
			}

			vac, err := scanMapIntoVacancy(fieldMap)
			if err != nil {
				return nil, err
			}

			result = append(result, vac)

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

func scanMapIntoVacancy(m map[string]interface{}) (*entity.VacancyEntity, error) {
	res := &entity.VacancyEntity{}
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
		case "description":
			description, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err. convert description")
			}
			res.Description = string(description)
		case "requirement":
			requirement, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err. convert requirement")
			}
			res.Description = string(requirement)
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
		case "salary_from":
			sf, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err. convert salary from")
			}

			s, err := strconv.ParseInt(string(sf), 10, 32)
			if err != nil {
				return nil, err
			}
			res.SalaryFrom = int32(s)
		case "salary_to":
			st, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err. convert salary to")
			}

			s, err := strconv.ParseInt(string(st), 10, 32)
			if err != nil {
				return nil, err
			}
			res.SalaryTo = int32(s)
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
