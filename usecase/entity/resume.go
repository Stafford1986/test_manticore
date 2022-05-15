package entity

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

const (
	QueryTypeInsert QueryType = "insert"
	QueryTypeUpsert QueryType = "upsert"
)

type QueryType string

//easyjson:json
type ResumeEntity struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Active    int    `json:"active"`
	SalaryMin int32  `json:"salary_min"`
	SalaryMax int32  `json:"salary_max"`
	CreatedAT int32  `json:"created_at"`
	UpdatedAt int32  `json:"updated_at"`
	Skills    string `json:"skills"`
	City      int32  `json:"city"`
	CityLevel int32  `json:"city_level"`
}

//easyjson:json
type Resumes []*ResumeEntity

func (r *ResumeEntity) Validate() error {
	if r.Id == 0 {
		return &ValidationErr{
			Err: "err. empty id field",
		}
	}

	if r.Name == "" {
		return &ValidationErr{
			Err: "err. empty name field",
		}
	}

	if r.SalaryMin == 0 {
		return &ValidationErr{
			Err: "err. empty salary min field",
		}
	}

	if r.SalaryMax == 0 {
		return &ValidationErr{
			Err: "err. empty salary max field",
		}
	}

	if r.CreatedAT == 0 {
		return &ValidationErr{
			Err: "err. empty created_at field",
		}
	}

	if r.UpdatedAt == 0 {
		return &ValidationErr{
			Err: "err. empty updated_at field",
		}
	}

	if r.Skills == "" {
		return &ValidationErr{
			Err: "err. empty skills field",
		}
	}

	if r.City == 0 {
		return &ValidationErr{
			Err: "err. empty city field",
		}
	}

	if r.CityLevel == 0 {
		return &ValidationErr{
			Err: "err. empty city level field",
		}
	}

	for _, s := range forbiddenChars {
		if strings.ContainsRune(r.Name, s) || strings.ContainsRune(r.Skills, s) {
			return ErrForbiddenCharacter
		}
	}

	return nil
}

func (re *ResumeEntity) BuildQuery(qt QueryType) (string, error) {
	var err error
	sb := bytes.NewBufferString("")

	switch qt {
	case QueryTypeInsert:
		_, err = sb.WriteString("INSERT INTO " +
			"resumes(id, name, active, salary_min, salary_max, created_at, updated_at, skills, city, city_level) VALUES(")
	case QueryTypeUpsert:
		_, err = sb.WriteString("REPLACE INTO " +
			"resumes(id, name, active, salary_min, salary_max, created_at, updated_at, skills, city, city_level) VALUES(")
	default:
		return "", fmt.Errorf("err. wrong query type: %s", qt)
	}

	_, err = sb.WriteString(strconv.FormatInt(re.Id, 10))
	_, err = sb.WriteString(", ")
	_, err = sb.WriteString("'")
	_, err = sb.WriteString(re.Name)
	_, err = sb.WriteString("'")
	_, err = sb.WriteString(", ")
	_, err = sb.WriteString(strconv.Itoa(re.Active))
	_, err = sb.WriteString(", ")
	_, err = sb.WriteString(strconv.FormatInt(int64(re.SalaryMin), 10))
	_, err = sb.WriteString(", ")
	_, err = sb.WriteString(strconv.FormatInt(int64(re.SalaryMax), 10))
	_, err = sb.WriteString(", ")
	_, err = sb.WriteString(strconv.FormatInt(int64(re.CreatedAT), 10))
	_, err = sb.WriteString(", ")
	_, err = sb.WriteString(strconv.FormatInt(int64(re.UpdatedAt), 10))
	_, err = sb.WriteString(", ")
	_, err = sb.WriteString("'")
	_, err = sb.WriteString(re.Skills)
	_, err = sb.WriteString("'")
	_, err = sb.WriteString(", ")
	_, err = sb.WriteString(strconv.FormatInt(int64(re.City), 10))
	_, err = sb.WriteString(", ")
	_, err = sb.WriteString(strconv.FormatInt(int64(re.CityLevel), 10))
	_, err = sb.WriteString(");")

	if err != nil {
		return "", err
	}

	return sb.String(), nil

	//return fmt.Sprintf(`INSERT INTO resumes(
	//   id, name, active, salary_min, salary_max, created_at, updated_at, skills, city, city_level)
	//   VALUES(%d, '%s', %d, %d, %d, %d, %d, '%s', %d, %d)`,
	//	re.Id,
	//	re.Name,
	//	re.Active,
	//	re.SalaryMin,
	//	re.SalaryMax,
	//	re.CreatedAT,
	//	re.UpdatedAt,
	//	re.Skills,
	//	re.City,
	//	re.CityLevel,
	//)
}
