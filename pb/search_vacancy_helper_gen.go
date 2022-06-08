// Code generated by generator, DO NOT EDIT.
package pb

import (
	"bytes"
	"errors"
	"fmt"
	gomultierror "github.com/hashicorp/go-multierror"
	"strconv"
)

var (
	targetVacancyFieldMap     = make(map[string]string, 39)
	targetVacancyFieldMapType = make(map[string]string, 39)
)

func init() {
	targetVacancyFieldMap["salary_curr"] = ""
	targetVacancyFieldMap["experience"] = "filter"
	targetVacancyFieldMap["id"] = ""
	targetVacancyFieldMap["city"] = "filter"
	targetVacancyFieldMap["website_url"] = ""
	targetVacancyFieldMap["created_at"] = ""
	targetVacancyFieldMap["metro"] = "filter"
	targetVacancyFieldMap["salary_to"] = "filter"
	targetVacancyFieldMap["work_condition"] = "index"
	targetVacancyFieldMap["city_visibility"] = "filter"
	targetVacancyFieldMap["salary_from"] = "filter"
	targetVacancyFieldMap["default_work_type"] = "filter"
	targetVacancyFieldMap["skills"] = "index"
	targetVacancyFieldMap["industry_groups"] = "filter"
	targetVacancyFieldMap["job_requirement"] = "index"
	targetVacancyFieldMap["business_tripp"] = "filter"
	targetVacancyFieldMap["active"] = ""
	targetVacancyFieldMap["requirement"] = "index"
	targetVacancyFieldMap["city_level"] = ""
	targetVacancyFieldMap["work_type"] = ""
	targetVacancyFieldMap["min_customer_languages"] = "filter"
	targetVacancyFieldMap["driver_license"] = "filter"
	targetVacancyFieldMap["have_car"] = "filter"
	targetVacancyFieldMap["list_respond_button"] = ""
	targetVacancyFieldMap["brand"] = "index"
	targetVacancyFieldMap["self_employed"] = "filter"
	targetVacancyFieldMap["payment_period"] = "filter"
	targetVacancyFieldMap["job_responsibility"] = "index"
	targetVacancyFieldMap["restrictions"] = "filter"
	targetVacancyFieldMap["salary_before_tax"] = ""
	targetVacancyFieldMap["vacancy_language"] = "filter"
	targetVacancyFieldMap["updated_at"] = ""
	targetVacancyFieldMap["logo"] = ""
	targetVacancyFieldMap["specialization"] = "filter"
	targetVacancyFieldMap["ip_employed"] = "filter"
	targetVacancyFieldMap["driver_exp"] = "filter"
	targetVacancyFieldMap["name"] = "index"
	targetVacancyFieldMap["salary"] = ""
	targetVacancyFieldMap["company_id"] = ""
	targetVacancyFieldMapType["id"] = "uint32"
	targetVacancyFieldMapType["name"] = "string"
	targetVacancyFieldMapType["active"] = "bool"
	targetVacancyFieldMapType["created_at"] = "uint64"
	targetVacancyFieldMapType["updated_at"] = "uint64"
	targetVacancyFieldMapType["requirement"] = "string"
	targetVacancyFieldMapType["salary"] = "uint32"
	targetVacancyFieldMapType["skills"] = "string"
	targetVacancyFieldMapType["city"] = "uint32"
	targetVacancyFieldMapType["city_level"] = "uint32"
	targetVacancyFieldMapType["brand"] = "string"
	targetVacancyFieldMapType["website_url"] = "string"
	targetVacancyFieldMapType["logo"] = "string"
	targetVacancyFieldMapType["industry_groups"] = "[]uint32"
	targetVacancyFieldMapType["company_id"] = "uint32"
	targetVacancyFieldMapType["specialization"] = "[]uint32"
	targetVacancyFieldMapType["metro"] = "[]uint32"
	targetVacancyFieldMapType["salary_before_tax"] = "bool"
	targetVacancyFieldMapType["salary_curr"] = "uint32"
	targetVacancyFieldMapType["job_responsibility"] = "string"
	targetVacancyFieldMapType["job_requirement"] = "string"
	targetVacancyFieldMapType["work_condition"] = "string"
	targetVacancyFieldMapType["city_visibility"] = "[]uint32"
	targetVacancyFieldMapType["vacancy_language"] = "uint32"
	targetVacancyFieldMapType["business_tripp"] = "uint32"
	targetVacancyFieldMapType["self_employed"] = "bool"
	targetVacancyFieldMapType["ip_employed"] = "bool"
	targetVacancyFieldMapType["payment_period"] = "uint32"
	targetVacancyFieldMapType["salary_from"] = "uint32"
	targetVacancyFieldMapType["salary_to"] = "uint32"
	targetVacancyFieldMapType["default_work_type"] = "uint32"
	targetVacancyFieldMapType["work_type"] = "[]uint32"
	targetVacancyFieldMapType["experience"] = "uint32"
	targetVacancyFieldMapType["min_customer_languages"] = "[]uint32"
	targetVacancyFieldMapType["driver_license"] = "[]uint32"
	targetVacancyFieldMapType["driver_exp"] = "uint32"
	targetVacancyFieldMapType["have_car"] = "bool"
	targetVacancyFieldMapType["restrictions"] = "[]uint32"
	targetVacancyFieldMapType["list_respond_button"] = "bool"
}
func (re *VacancySearchEntity) BuildSearchQuery() (string, error) {
	var (
		resErr error
	)
	if len(re.Query) == 0 {
		return "", errors.New("err validation. empty search req")
	}
	query := re.Query
	sb := bytes.NewBufferString(fmt.Sprintf("SELECT * FROM resumes WHERE MATCH('*%s*')", query))
	for _, v := range re.FilterParams {
		if v.Filter == "" {
			return "", errors.New("err validation. empty filter field")
		}
		fv, ok := targetVacancyFieldMap[v.Filter]
		if !ok {
			return "", errors.New("err validation. wrong filter param name")
		}
		if fv != "filter" {
			return "", errors.New("err validation. field can't be used as filter")
		}
		switch v.Value.(type) {
		case *FilterParams_SingleInt:
			ft, ok := targetVacancyFieldMapType[v.Filter]
			if !ok {
				return "", errors.New("err validation. can't find filter name")
			}
			if ft != "uint32" {
				return "", errors.New("err validation. mismatch filter value type")
			}
			_, err := sb.WriteString(" AND ")
			_, err = sb.WriteString(v.Filter)
			_, err = sb.WriteString(" = ")
			_, err = sb.WriteString(strconv.FormatInt(int64(v.GetSingleInt()), 10))
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
		case *FilterParams_SingleString:
			ft, ok := targetVacancyFieldMapType[v.Filter]
			if !ok {
				return "", errors.New("err validation. can't find filter name")
			}
			if ft != "string" {
				return "", errors.New("err validation. mismatch filter value type")
			}
			_, err := sb.WriteString(" AND ")
			_, err = sb.WriteString(v.Filter)
			_, err = sb.WriteString(" = ")
			_, err = sb.WriteString(v.GetSingleString())
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
		case *FilterParams_IntArray:
			ft, ok := targetVacancyFieldMapType[v.Filter]
			if !ok {
				return "", errors.New("err validation. can't find filter name")
			}
			if ft != "[]uint32" {
				return "", errors.New("err validation. mismatch filter value type")
			}
			arr := v.GetIntArray()
			if arr == nil {
				return "", errors.New("err validation. got nil array")
			}
			arrVal := arr.GetValues()
			if len(arrVal) == 0 {
				return "", errors.New("err validation. got empty array")
			}
			var (
				sbp = bytes.NewBufferString("")
			)
			for i, v := range arrVal {
				var err error
				if i == 0 {
					_, err = sbp.WriteString("(")
				}
				_, err = sbp.WriteString(strconv.FormatInt(int64(v), 10))
				if i == len(arrVal)-1 {
					_, err = sbp.WriteString(")")
				} else {
					_, err = sbp.WriteString(",")
				}
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
			}
			_, err := sb.WriteString(" AND ")
			_, err = sb.WriteString(v.Filter)
			_, err = sb.WriteString(" IN")
			_, err = sb.WriteString(sbp.String())
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
		}
	}
	for _, v := range re.SortParams {
		if v.Field == "" {
			return "", errors.New("err validation. empty sort field")
		}
		fv, ok := targetVacancyFieldMap[v.Field]
		if !ok {
			return "", errors.New("err validation. wrong sort param name")
		}
		if fv != "" {
			return "", errors.New("err validation. field can't be used as sorted")
		}
		switch v.Order.(type) {
		case *SortParams_ASC:
			_, err := sb.WriteString(" ORDER BY ")
			_, err = sb.WriteString(v.Field)
			_, err = sb.WriteString(" ASC")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
		case *SortParams_DESC:
			_, err := sb.WriteString(" ORDER BY ")
			_, err = sb.WriteString(v.Field)
			_, err = sb.WriteString(" DESC")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
		}
	}
	_, err := sb.WriteString(";")
	if err != nil {
		resErr = gomultierror.Append(resErr, err)
	}
	if resErr != nil {
		return "", resErr
	}
	return sb.String(), nil
}