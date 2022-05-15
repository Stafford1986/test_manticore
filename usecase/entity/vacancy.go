package entity

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

//easyjson:json
type VacancyEntity struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Requirement string `json:"requirement"`
	Active      int    `json:"active"`
	SalaryFrom  int32  `json:"salary_from"`
	SalaryTo    int32  `json:"salary_to"`
	CreatedAT   int32  `json:"created_at"`
	UpdatedAt   int32  `json:"updated_at"`
	Skills      string `json:"skills"`
	City        int32  `json:"city"`
	CityLevel   int32  `json:"city_level"`
}

//easyjson:json
type Vacancies []*VacancyEntity

func (v *VacancyEntity) Validate() error {
	if v.Id == 0 {
		return &ValidationErr{
			Err: "err. empty id field",
		}
	}

	if v.Name == "" {
		return &ValidationErr{
			Err: "err. empty name field",
		}
	}

	if v.Description == "" {
		return &ValidationErr{
			Err: "err. empty description field",
		}
	}

	if v.Requirement == "" {
		return &ValidationErr{
			Err: "err. empty requirement field",
		}
	}

	if v.SalaryFrom == 0 {
		return &ValidationErr{
			Err: "err. empty salary from field",
		}
	}

	if v.SalaryTo == 0 {
		return &ValidationErr{
			Err: "err. empty salary to field",
		}
	}

	if v.CreatedAT == 0 {
		return &ValidationErr{
			Err: "err. empty created_at field",
		}
	}

	if v.UpdatedAt == 0 {
		return &ValidationErr{
			Err: "err. empty updated_at field",
		}
	}

	if v.Skills == "" {
		return &ValidationErr{
			Err: "err. empty skills field",
		}
	}

	if v.City == 0 {
		return &ValidationErr{
			Err: "err. empty city field",
		}
	}

	if v.CityLevel == 0 {
		return &ValidationErr{
			Err: "err. empty city level field",
		}
	}

	for _, s := range forbiddenChars {
		if strings.ContainsRune(v.Name, s) || strings.ContainsRune(v.Description, s) ||
			strings.ContainsRune(v.Requirement, s) || strings.ContainsRune(v.Skills, s) {
			return ErrForbiddenCharacter
		}
	}

	return nil
}

func (v *VacancyEntity) BuildQuery(qt QueryType) (string, error) {
	var err error
	sb := bytes.NewBufferString("")

	switch qt {
	case QueryTypeInsert:
		_, err = sb.WriteString("INSERT INTO " +
			"vacancies(id, name, description, requirement, " +
			"active, salary_from, salary_to, created_at, updated_at, skills, city, city_level) VALUES(")
	case QueryTypeUpsert:
		_, err = sb.WriteString("REPLACE INTO " +
			"vacancies(id, name, description, requirement, " +
			"active, salary_from, salary_to, created_at, updated_at, skills, city, city_level) VALUES(")
	default:
		return "", fmt.Errorf("err. wrong query type: %s", qt)
	}

	_, err = sb.WriteString(strconv.FormatInt(v.Id, 10))
	_, err = sb.WriteString(", ")
	_, err = sb.WriteString("'")
	_, err = sb.WriteString(v.Name)
	_, err = sb.WriteString("'")
	_, err = sb.WriteString(", ")
	_, err = sb.WriteString("'")
	_, err = sb.WriteString(v.Description)
	_, err = sb.WriteString("'")
	_, err = sb.WriteString(", ")
	_, err = sb.WriteString("'")
	_, err = sb.WriteString(v.Requirement)
	_, err = sb.WriteString("'")
	_, err = sb.WriteString(", ")
	_, err = sb.WriteString(strconv.Itoa(v.Active))
	_, err = sb.WriteString(", ")
	_, err = sb.WriteString(strconv.FormatInt(int64(v.SalaryFrom), 10))
	_, err = sb.WriteString(", ")
	_, err = sb.WriteString(strconv.FormatInt(int64(v.SalaryTo), 10))
	_, err = sb.WriteString(", ")
	_, err = sb.WriteString(strconv.FormatInt(int64(v.CreatedAT), 10))
	_, err = sb.WriteString(", ")
	_, err = sb.WriteString(strconv.FormatInt(int64(v.UpdatedAt), 10))
	_, err = sb.WriteString(", ")
	_, err = sb.WriteString("'")
	_, err = sb.WriteString(v.Skills)
	_, err = sb.WriteString("'")
	_, err = sb.WriteString(", ")
	_, err = sb.WriteString(strconv.FormatInt(int64(v.City), 10))
	_, err = sb.WriteString(", ")
	_, err = sb.WriteString(strconv.FormatInt(int64(v.CityLevel), 10))
	_, err = sb.WriteString(");")

	if err != nil {
		return "", err
	}

	return sb.String(), nil

}
