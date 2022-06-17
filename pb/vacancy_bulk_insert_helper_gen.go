// Code generated by generator, DO NOT EDIT.
package pb

import (
	"bytes"
	gomultierror "github.com/hashicorp/go-multierror"
	"strconv"
	"strings"
)

func (res *VacancyList) BuildBulkInsertQuery() (string, error) {
	var (
		resErr       error
		allResValues = make([]string, 0, len(res.Items))
	)
	sb := bytes.NewBufferString("INSERT INTO vacancies(id, name, active, created_at, updated_at, skills, city, city_level, brand, website_url, logo, industry_groups, company_id, specialization, metro, salary_before_tax, salary_curr, job_responsibility, job_requirement, work_condition, city_visibility, vacancy_language, business_trips, self_employed, ip_employed, payment_period, salary_from, salary_to, default_work_type, work_type, experience, min_customer_languages, driver_license, driver_exp, have_car, restrictions, list_respond_button) VALUES")
	for _, r := range res.Items {
		var (
			resValues = make([]string, 0, 37)
		)
		sbb := bytes.NewBufferString("(")
		resValues = append(resValues, strconv.FormatInt(int64(r.Id), 10))
		resValues = append(resValues, func() string {
			sb := bytes.NewBufferString("'")
			if len(r.Name) == 0 {
				_, err := sb.WriteString(" '")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
				return sb.String()
			}
			_, err := sb.WriteString(r.Name)
			_, err = sb.WriteString("'")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}())
		resValues = append(resValues, func() string {
			if r.Active == true {
				return "1"
			}
			return "0"
		}())
		resValues = append(resValues, strconv.FormatInt(int64(r.CreatedAt), 10))
		resValues = append(resValues, strconv.FormatInt(int64(r.UpdatedAt), 10))
		resValues = append(resValues, func() string {
			sb := bytes.NewBufferString("'")
			if len(r.Skills) == 0 {
				_, err := sb.WriteString(" '")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
				return sb.String()
			}
			_, err := sb.WriteString(r.Skills)
			_, err = sb.WriteString("'")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}())
		resValues = append(resValues, strconv.FormatInt(int64(r.City), 10))
		resValues = append(resValues, strconv.FormatInt(int64(r.CityLevel), 10))
		resValues = append(resValues, func() string {
			sb := bytes.NewBufferString("'")
			if len(r.Brand) == 0 {
				_, err := sb.WriteString(" '")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
				return sb.String()
			}
			_, err := sb.WriteString(r.Brand)
			_, err = sb.WriteString("'")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}())
		resValues = append(resValues, func() string {
			sb := bytes.NewBufferString("'")
			if len(r.WebsiteUrl) == 0 {
				_, err := sb.WriteString(" '")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
				return sb.String()
			}
			_, err := sb.WriteString(r.WebsiteUrl)
			_, err = sb.WriteString("'")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}())
		resValues = append(resValues, func() string {
			sb := bytes.NewBufferString("'")
			if len(r.Logo) == 0 {
				_, err := sb.WriteString(" '")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
				return sb.String()
			}
			_, err := sb.WriteString(r.Logo)
			_, err = sb.WriteString("'")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}())
		resValues = append(resValues, func() string {
			sb := bytes.NewBufferString("(")
			if len(r.IndustryGroups) == 0 {
				_, err := sb.WriteString("0)")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
				return sb.String()
			}
			for i, v := range r.IndustryGroups {
				if i == len(r.IndustryGroups)-1 {
					_, err := sb.WriteString(strconv.FormatInt(int64(v), 10))
					_, err = sb.WriteString(")")
					if err != nil {
						resErr = gomultierror.Append(resErr, err)
					}
					break
				}
				_, err := sb.WriteString(strconv.FormatInt(int64(v), 10))
				_, err = sb.WriteString(",")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
			}
			return sb.String()
		}())
		resValues = append(resValues, strconv.FormatInt(int64(r.CompanyId), 10))
		resValues = append(resValues, func() string {
			sb := bytes.NewBufferString("(")
			if len(r.Specialization) == 0 {
				_, err := sb.WriteString("0)")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
				return sb.String()
			}
			for i, v := range r.Specialization {
				if i == len(r.Specialization)-1 {
					_, err := sb.WriteString(strconv.FormatInt(int64(v), 10))
					_, err = sb.WriteString(")")
					if err != nil {
						resErr = gomultierror.Append(resErr, err)
					}
					break
				}
				_, err := sb.WriteString(strconv.FormatInt(int64(v), 10))
				_, err = sb.WriteString(",")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
			}
			return sb.String()
		}())
		resValues = append(resValues, func() string {
			sb := bytes.NewBufferString("(")
			if len(r.Metro) == 0 {
				_, err := sb.WriteString("0)")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
				return sb.String()
			}
			for i, v := range r.Metro {
				if i == len(r.Metro)-1 {
					_, err := sb.WriteString(strconv.FormatInt(int64(v), 10))
					_, err = sb.WriteString(")")
					if err != nil {
						resErr = gomultierror.Append(resErr, err)
					}
					break
				}
				_, err := sb.WriteString(strconv.FormatInt(int64(v), 10))
				_, err = sb.WriteString(",")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
			}
			return sb.String()
		}())
		resValues = append(resValues, func() string {
			if r.SalaryBeforeTax == true {
				return "1"
			}
			return "0"
		}())
		resValues = append(resValues, strconv.FormatInt(int64(r.SalaryCurr), 10))
		resValues = append(resValues, func() string {
			sb := bytes.NewBufferString("'")
			if len(r.JobResponsibility) == 0 {
				_, err := sb.WriteString(" '")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
				return sb.String()
			}
			_, err := sb.WriteString(r.JobResponsibility)
			_, err = sb.WriteString("'")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}())
		resValues = append(resValues, func() string {
			sb := bytes.NewBufferString("'")
			if len(r.JobRequirement) == 0 {
				_, err := sb.WriteString(" '")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
				return sb.String()
			}
			_, err := sb.WriteString(r.JobRequirement)
			_, err = sb.WriteString("'")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}())
		resValues = append(resValues, func() string {
			sb := bytes.NewBufferString("'")
			if len(r.WorkCondition) == 0 {
				_, err := sb.WriteString(" '")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
				return sb.String()
			}
			_, err := sb.WriteString(r.WorkCondition)
			_, err = sb.WriteString("'")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}())
		resValues = append(resValues, func() string {
			sb := bytes.NewBufferString("(")
			if len(r.CityVisibility) == 0 {
				_, err := sb.WriteString("0)")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
				return sb.String()
			}
			for i, v := range r.CityVisibility {
				if i == len(r.CityVisibility)-1 {
					_, err := sb.WriteString(strconv.FormatInt(int64(v), 10))
					_, err = sb.WriteString(")")
					if err != nil {
						resErr = gomultierror.Append(resErr, err)
					}
					break
				}
				_, err := sb.WriteString(strconv.FormatInt(int64(v), 10))
				_, err = sb.WriteString(",")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
			}
			return sb.String()
		}())
		resValues = append(resValues, strconv.FormatInt(int64(r.VacancyLanguage), 10))
		resValues = append(resValues, strconv.FormatInt(int64(r.BusinessTrips), 10))
		resValues = append(resValues, func() string {
			if r.SelfEmployed == true {
				return "1"
			}
			return "0"
		}())
		resValues = append(resValues, func() string {
			if r.IpEmployed == true {
				return "1"
			}
			return "0"
		}())
		resValues = append(resValues, strconv.FormatInt(int64(r.PaymentPeriod), 10))
		resValues = append(resValues, strconv.FormatInt(int64(r.SalaryFrom), 10))
		resValues = append(resValues, strconv.FormatInt(int64(r.SalaryTo), 10))
		resValues = append(resValues, strconv.FormatInt(int64(r.DefaultWorkType), 10))
		resValues = append(resValues, func() string {
			sb := bytes.NewBufferString("(")
			if len(r.WorkType) == 0 {
				_, err := sb.WriteString("0)")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
				return sb.String()
			}
			for i, v := range r.WorkType {
				if i == len(r.WorkType)-1 {
					_, err := sb.WriteString(strconv.FormatInt(int64(v), 10))
					_, err = sb.WriteString(")")
					if err != nil {
						resErr = gomultierror.Append(resErr, err)
					}
					break
				}
				_, err := sb.WriteString(strconv.FormatInt(int64(v), 10))
				_, err = sb.WriteString(",")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
			}
			return sb.String()
		}())
		resValues = append(resValues, strconv.FormatInt(int64(r.Experience), 10))
		resValues = append(resValues, func() string {
			sb := bytes.NewBufferString("(")
			if len(r.MinCustomerLanguages) == 0 {
				_, err := sb.WriteString("0)")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
				return sb.String()
			}
			for i, v := range r.MinCustomerLanguages {
				if i == len(r.MinCustomerLanguages)-1 {
					_, err := sb.WriteString(strconv.FormatInt(int64(v), 10))
					_, err = sb.WriteString(")")
					if err != nil {
						resErr = gomultierror.Append(resErr, err)
					}
					break
				}
				_, err := sb.WriteString(strconv.FormatInt(int64(v), 10))
				_, err = sb.WriteString(",")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
			}
			return sb.String()
		}())
		resValues = append(resValues, func() string {
			sb := bytes.NewBufferString("(")
			if len(r.DriverLicense) == 0 {
				_, err := sb.WriteString("0)")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
				return sb.String()
			}
			for i, v := range r.DriverLicense {
				if i == len(r.DriverLicense)-1 {
					_, err := sb.WriteString(strconv.FormatInt(int64(v), 10))
					_, err = sb.WriteString(")")
					if err != nil {
						resErr = gomultierror.Append(resErr, err)
					}
					break
				}
				_, err := sb.WriteString(strconv.FormatInt(int64(v), 10))
				_, err = sb.WriteString(",")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
			}
			return sb.String()
		}())
		resValues = append(resValues, strconv.FormatInt(int64(r.DriverExp), 10))
		resValues = append(resValues, func() string {
			if r.HaveCar == true {
				return "1"
			}
			return "0"
		}())
		resValues = append(resValues, func() string {
			sb := bytes.NewBufferString("(")
			if len(r.Restrictions) == 0 {
				_, err := sb.WriteString("0)")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
				return sb.String()
			}
			for i, v := range r.Restrictions {
				if i == len(r.Restrictions)-1 {
					_, err := sb.WriteString(strconv.FormatInt(int64(v), 10))
					_, err = sb.WriteString(")")
					if err != nil {
						resErr = gomultierror.Append(resErr, err)
					}
					break
				}
				_, err := sb.WriteString(strconv.FormatInt(int64(v), 10))
				_, err = sb.WriteString(",")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
			}
			return sb.String()
		}())
		resValues = append(resValues, func() string {
			if r.ListRespondButton == true {
				return "1"
			}
			return "0"
		}())
		_, err := sbb.WriteString(strings.Join(resValues, ", "))
		_, err = sbb.WriteString(")")
		if err != nil {
			resErr = gomultierror.Append(resErr, err)
		}
		allResValues = append(allResValues, sbb.String())
	}
	_, err := sb.WriteString(strings.Join(allResValues, ", "))
	_, err = sb.WriteString(";")
	if err != nil {
		resErr = gomultierror.Append(resErr, err)
	}
	if resErr != nil {
		return "", resErr
	}
	return sb.String(), nil
}
func (res *VacancyList) BuildBulkUpsertQuery() (string, error) {
	var (
		resErr       error
		allResValues = make([]string, 0, len(res.Items))
	)
	sb := bytes.NewBufferString("REPLACE INTO vacancies(id, name, active, created_at, updated_at, skills, city, city_level, brand, website_url, logo, industry_groups, company_id, specialization, metro, salary_before_tax, salary_curr, job_responsibility, job_requirement, work_condition, city_visibility, vacancy_language, business_trips, self_employed, ip_employed, payment_period, salary_from, salary_to, default_work_type, work_type, experience, min_customer_languages, driver_license, driver_exp, have_car, restrictions, list_respond_button) VALUES")
	for _, r := range res.Items {
		var (
			resValues = make([]string, 0, 37)
		)
		sbb := bytes.NewBufferString("(")
		resValues = append(resValues, strconv.FormatInt(int64(r.Id), 10))
		resValues = append(resValues, func() string {
			sb := bytes.NewBufferString("'")
			if len(r.Name) == 0 {
				_, err := sb.WriteString(" '")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
				return sb.String()
			}
			_, err := sb.WriteString(r.Name)
			_, err = sb.WriteString("'")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}())
		resValues = append(resValues, func() string {
			if r.Active == true {
				return "1"
			}
			return "0"
		}())
		resValues = append(resValues, strconv.FormatInt(int64(r.CreatedAt), 10))
		resValues = append(resValues, strconv.FormatInt(int64(r.UpdatedAt), 10))
		resValues = append(resValues, func() string {
			sb := bytes.NewBufferString("'")
			if len(r.Skills) == 0 {
				_, err := sb.WriteString(" '")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
				return sb.String()
			}
			_, err := sb.WriteString(r.Skills)
			_, err = sb.WriteString("'")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}())
		resValues = append(resValues, strconv.FormatInt(int64(r.City), 10))
		resValues = append(resValues, strconv.FormatInt(int64(r.CityLevel), 10))
		resValues = append(resValues, func() string {
			sb := bytes.NewBufferString("'")
			if len(r.Brand) == 0 {
				_, err := sb.WriteString(" '")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
				return sb.String()
			}
			_, err := sb.WriteString(r.Brand)
			_, err = sb.WriteString("'")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}())
		resValues = append(resValues, func() string {
			sb := bytes.NewBufferString("'")
			if len(r.WebsiteUrl) == 0 {
				_, err := sb.WriteString(" '")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
				return sb.String()
			}
			_, err := sb.WriteString(r.WebsiteUrl)
			_, err = sb.WriteString("'")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}())
		resValues = append(resValues, func() string {
			sb := bytes.NewBufferString("'")
			if len(r.Logo) == 0 {
				_, err := sb.WriteString(" '")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
				return sb.String()
			}
			_, err := sb.WriteString(r.Logo)
			_, err = sb.WriteString("'")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}())
		resValues = append(resValues, func() string {
			sb := bytes.NewBufferString("(")
			if len(r.IndustryGroups) == 0 {
				_, err := sb.WriteString("0)")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
				return sb.String()
			}
			for i, v := range r.IndustryGroups {
				if i == len(r.IndustryGroups)-1 {
					_, err := sb.WriteString(strconv.FormatInt(int64(v), 10))
					_, err = sb.WriteString(")")
					if err != nil {
						resErr = gomultierror.Append(resErr, err)
					}
					break
				}
				_, err := sb.WriteString(strconv.FormatInt(int64(v), 10))
				_, err = sb.WriteString(",")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
			}
			return sb.String()
		}())
		resValues = append(resValues, strconv.FormatInt(int64(r.CompanyId), 10))
		resValues = append(resValues, func() string {
			sb := bytes.NewBufferString("(")
			if len(r.Specialization) == 0 {
				_, err := sb.WriteString("0)")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
				return sb.String()
			}
			for i, v := range r.Specialization {
				if i == len(r.Specialization)-1 {
					_, err := sb.WriteString(strconv.FormatInt(int64(v), 10))
					_, err = sb.WriteString(")")
					if err != nil {
						resErr = gomultierror.Append(resErr, err)
					}
					break
				}
				_, err := sb.WriteString(strconv.FormatInt(int64(v), 10))
				_, err = sb.WriteString(",")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
			}
			return sb.String()
		}())
		resValues = append(resValues, func() string {
			sb := bytes.NewBufferString("(")
			if len(r.Metro) == 0 {
				_, err := sb.WriteString("0)")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
				return sb.String()
			}
			for i, v := range r.Metro {
				if i == len(r.Metro)-1 {
					_, err := sb.WriteString(strconv.FormatInt(int64(v), 10))
					_, err = sb.WriteString(")")
					if err != nil {
						resErr = gomultierror.Append(resErr, err)
					}
					break
				}
				_, err := sb.WriteString(strconv.FormatInt(int64(v), 10))
				_, err = sb.WriteString(",")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
			}
			return sb.String()
		}())
		resValues = append(resValues, func() string {
			if r.SalaryBeforeTax == true {
				return "1"
			}
			return "0"
		}())
		resValues = append(resValues, strconv.FormatInt(int64(r.SalaryCurr), 10))
		resValues = append(resValues, func() string {
			sb := bytes.NewBufferString("'")
			if len(r.JobResponsibility) == 0 {
				_, err := sb.WriteString(" '")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
				return sb.String()
			}
			_, err := sb.WriteString(r.JobResponsibility)
			_, err = sb.WriteString("'")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}())
		resValues = append(resValues, func() string {
			sb := bytes.NewBufferString("'")
			if len(r.JobRequirement) == 0 {
				_, err := sb.WriteString(" '")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
				return sb.String()
			}
			_, err := sb.WriteString(r.JobRequirement)
			_, err = sb.WriteString("'")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}())
		resValues = append(resValues, func() string {
			sb := bytes.NewBufferString("'")
			if len(r.WorkCondition) == 0 {
				_, err := sb.WriteString(" '")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
				return sb.String()
			}
			_, err := sb.WriteString(r.WorkCondition)
			_, err = sb.WriteString("'")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}())
		resValues = append(resValues, func() string {
			sb := bytes.NewBufferString("(")
			if len(r.CityVisibility) == 0 {
				_, err := sb.WriteString("0)")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
				return sb.String()
			}
			for i, v := range r.CityVisibility {
				if i == len(r.CityVisibility)-1 {
					_, err := sb.WriteString(strconv.FormatInt(int64(v), 10))
					_, err = sb.WriteString(")")
					if err != nil {
						resErr = gomultierror.Append(resErr, err)
					}
					break
				}
				_, err := sb.WriteString(strconv.FormatInt(int64(v), 10))
				_, err = sb.WriteString(",")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
			}
			return sb.String()
		}())
		resValues = append(resValues, strconv.FormatInt(int64(r.VacancyLanguage), 10))
		resValues = append(resValues, strconv.FormatInt(int64(r.BusinessTrips), 10))
		resValues = append(resValues, func() string {
			if r.SelfEmployed == true {
				return "1"
			}
			return "0"
		}())
		resValues = append(resValues, func() string {
			if r.IpEmployed == true {
				return "1"
			}
			return "0"
		}())
		resValues = append(resValues, strconv.FormatInt(int64(r.PaymentPeriod), 10))
		resValues = append(resValues, strconv.FormatInt(int64(r.SalaryFrom), 10))
		resValues = append(resValues, strconv.FormatInt(int64(r.SalaryTo), 10))
		resValues = append(resValues, strconv.FormatInt(int64(r.DefaultWorkType), 10))
		resValues = append(resValues, func() string {
			sb := bytes.NewBufferString("(")
			if len(r.WorkType) == 0 {
				_, err := sb.WriteString("0)")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
				return sb.String()
			}
			for i, v := range r.WorkType {
				if i == len(r.WorkType)-1 {
					_, err := sb.WriteString(strconv.FormatInt(int64(v), 10))
					_, err = sb.WriteString(")")
					if err != nil {
						resErr = gomultierror.Append(resErr, err)
					}
					break
				}
				_, err := sb.WriteString(strconv.FormatInt(int64(v), 10))
				_, err = sb.WriteString(",")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
			}
			return sb.String()
		}())
		resValues = append(resValues, strconv.FormatInt(int64(r.Experience), 10))
		resValues = append(resValues, func() string {
			sb := bytes.NewBufferString("(")
			if len(r.MinCustomerLanguages) == 0 {
				_, err := sb.WriteString("0)")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
				return sb.String()
			}
			for i, v := range r.MinCustomerLanguages {
				if i == len(r.MinCustomerLanguages)-1 {
					_, err := sb.WriteString(strconv.FormatInt(int64(v), 10))
					_, err = sb.WriteString(")")
					if err != nil {
						resErr = gomultierror.Append(resErr, err)
					}
					break
				}
				_, err := sb.WriteString(strconv.FormatInt(int64(v), 10))
				_, err = sb.WriteString(",")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
			}
			return sb.String()
		}())
		resValues = append(resValues, func() string {
			sb := bytes.NewBufferString("(")
			if len(r.DriverLicense) == 0 {
				_, err := sb.WriteString("0)")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
				return sb.String()
			}
			for i, v := range r.DriverLicense {
				if i == len(r.DriverLicense)-1 {
					_, err := sb.WriteString(strconv.FormatInt(int64(v), 10))
					_, err = sb.WriteString(")")
					if err != nil {
						resErr = gomultierror.Append(resErr, err)
					}
					break
				}
				_, err := sb.WriteString(strconv.FormatInt(int64(v), 10))
				_, err = sb.WriteString(",")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
			}
			return sb.String()
		}())
		resValues = append(resValues, strconv.FormatInt(int64(r.DriverExp), 10))
		resValues = append(resValues, func() string {
			if r.HaveCar == true {
				return "1"
			}
			return "0"
		}())
		resValues = append(resValues, func() string {
			sb := bytes.NewBufferString("(")
			if len(r.Restrictions) == 0 {
				_, err := sb.WriteString("0)")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
				return sb.String()
			}
			for i, v := range r.Restrictions {
				if i == len(r.Restrictions)-1 {
					_, err := sb.WriteString(strconv.FormatInt(int64(v), 10))
					_, err = sb.WriteString(")")
					if err != nil {
						resErr = gomultierror.Append(resErr, err)
					}
					break
				}
				_, err := sb.WriteString(strconv.FormatInt(int64(v), 10))
				_, err = sb.WriteString(",")
				if err != nil {
					resErr = gomultierror.Append(resErr, err)
				}
			}
			return sb.String()
		}())
		resValues = append(resValues, func() string {
			if r.ListRespondButton == true {
				return "1"
			}
			return "0"
		}())
		_, err := sbb.WriteString(strings.Join(resValues, ", "))
		_, err = sbb.WriteString(")")
		if err != nil {
			resErr = gomultierror.Append(resErr, err)
		}
		allResValues = append(allResValues, sbb.String())
	}
	_, err := sb.WriteString(strings.Join(allResValues, ", "))
	_, err = sb.WriteString(";")
	if err != nil {
		resErr = gomultierror.Append(resErr, err)
	}
	if resErr != nil {
		return "", resErr
	}
	return sb.String(), nil
}