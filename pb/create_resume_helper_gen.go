// Code generated by generator, DO NOT EDIT.
package pb

import (
	"bytes"
	gomultierror "github.com/hashicorp/go-multierror"
	"strconv"
	"strings"
)

func (re *ResumeEntity) BuildInsertQuery() (string, error) {
	var (
		resErr    error
		resValues = make([]string, 0, 30)
	)
	sb := bytes.NewBufferString("INSERT INTO resumes(id, name, active, created_at, updated_at, job_name, job_desc_duties, job_desc_achievements, company_name, institute_name, certificate_name, customer_id, languages, salary_curr, payment_period, industries, business_trip, default_work_type, regular_trip, work_license, resume_language, metro, work_type, additional_calc_form, about_short, salary, skills, city, specialization, status) VALUES(")
	resValues = append(resValues, strconv.FormatInt(int64(re.Id), 10))
	resValues = append(resValues, func() string {
		sb := bytes.NewBufferString("'")
		if len(re.Name) == 0 {
			_, err := sb.WriteString(" '")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}
		_, err := sb.WriteString(re.Name)
		_, err = sb.WriteString("'")
		if err != nil {
			resErr = gomultierror.Append(resErr, err)
		}
		return sb.String()
	}())
	resValues = append(resValues, func() string {
		if re.Active == true {
			return "1"
		}
		return "0"
	}())
	resValues = append(resValues, strconv.FormatInt(int64(re.CreatedAt), 10))
	resValues = append(resValues, strconv.FormatInt(int64(re.UpdatedAt), 10))
	resValues = append(resValues, func() string {
		sb := bytes.NewBufferString("'")
		if len(re.JobName) == 0 {
			_, err := sb.WriteString(" '")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}
		_, err := sb.WriteString(re.JobName)
		_, err = sb.WriteString("'")
		if err != nil {
			resErr = gomultierror.Append(resErr, err)
		}
		return sb.String()
	}())
	resValues = append(resValues, func() string {
		sb := bytes.NewBufferString("'")
		if len(re.JobDescDuties) == 0 {
			_, err := sb.WriteString(" '")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}
		_, err := sb.WriteString(re.JobDescDuties)
		_, err = sb.WriteString("'")
		if err != nil {
			resErr = gomultierror.Append(resErr, err)
		}
		return sb.String()
	}())
	resValues = append(resValues, func() string {
		sb := bytes.NewBufferString("'")
		if len(re.JobDescAchievements) == 0 {
			_, err := sb.WriteString(" '")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}
		_, err := sb.WriteString(re.JobDescAchievements)
		_, err = sb.WriteString("'")
		if err != nil {
			resErr = gomultierror.Append(resErr, err)
		}
		return sb.String()
	}())
	resValues = append(resValues, func() string {
		sb := bytes.NewBufferString("'")
		if len(re.CompanyName) == 0 {
			_, err := sb.WriteString(" '")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}
		_, err := sb.WriteString(re.CompanyName)
		_, err = sb.WriteString("'")
		if err != nil {
			resErr = gomultierror.Append(resErr, err)
		}
		return sb.String()
	}())
	resValues = append(resValues, func() string {
		sb := bytes.NewBufferString("'")
		if len(re.InstituteName) == 0 {
			_, err := sb.WriteString(" '")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}
		_, err := sb.WriteString(re.InstituteName)
		_, err = sb.WriteString("'")
		if err != nil {
			resErr = gomultierror.Append(resErr, err)
		}
		return sb.String()
	}())
	resValues = append(resValues, func() string {
		sb := bytes.NewBufferString("'")
		if len(re.CertificateName) == 0 {
			_, err := sb.WriteString(" '")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}
		_, err := sb.WriteString(re.CertificateName)
		_, err = sb.WriteString("'")
		if err != nil {
			resErr = gomultierror.Append(resErr, err)
		}
		return sb.String()
	}())
	resValues = append(resValues, strconv.FormatInt(int64(re.CustomerId), 10))
	resValues = append(resValues, func() string {
		sb := bytes.NewBufferString("(")
		if len(re.Languages) == 0 {
			_, err := sb.WriteString("0)")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}
		for i, v := range re.Languages {
			if i == len(re.Languages)-1 {
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
		sb := bytes.NewBufferString("'")
		if len(re.SalaryCurr) == 0 {
			_, err := sb.WriteString(" '")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}
		_, err := sb.WriteString(re.SalaryCurr)
		_, err = sb.WriteString("'")
		if err != nil {
			resErr = gomultierror.Append(resErr, err)
		}
		return sb.String()
	}())
	resValues = append(resValues, strconv.FormatInt(int64(re.PaymentPeriod), 10))
	resValues = append(resValues, func() string {
		sb := bytes.NewBufferString("(")
		if len(re.Industries) == 0 {
			_, err := sb.WriteString("0)")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}
		for i, v := range re.Industries {
			if i == len(re.Industries)-1 {
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
	resValues = append(resValues, strconv.FormatInt(int64(re.BusinessTrip), 10))
	resValues = append(resValues, strconv.FormatInt(int64(re.DefaultWorkType), 10))
	resValues = append(resValues, strconv.FormatInt(int64(re.RegularTrip), 10))
	resValues = append(resValues, func() string {
		sb := bytes.NewBufferString("(")
		if len(re.WorkLicense) == 0 {
			_, err := sb.WriteString("0)")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}
		for i, v := range re.WorkLicense {
			if i == len(re.WorkLicense)-1 {
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
	resValues = append(resValues, strconv.FormatInt(int64(re.ResumeLanguage), 10))
	resValues = append(resValues, func() string {
		sb := bytes.NewBufferString("(")
		if len(re.Metro) == 0 {
			_, err := sb.WriteString("0)")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}
		for i, v := range re.Metro {
			if i == len(re.Metro)-1 {
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
		if len(re.WorkType) == 0 {
			_, err := sb.WriteString("0)")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}
		for i, v := range re.WorkType {
			if i == len(re.WorkType)-1 {
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
		if len(re.AdditionalCalcForm) == 0 {
			_, err := sb.WriteString("0)")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}
		for i, v := range re.AdditionalCalcForm {
			if i == len(re.AdditionalCalcForm)-1 {
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
		sb := bytes.NewBufferString("'")
		if len(re.AboutShort) == 0 {
			_, err := sb.WriteString(" '")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}
		_, err := sb.WriteString(re.AboutShort)
		_, err = sb.WriteString("'")
		if err != nil {
			resErr = gomultierror.Append(resErr, err)
		}
		return sb.String()
	}())
	resValues = append(resValues, strconv.FormatInt(int64(re.Salary), 10))
	resValues = append(resValues, func() string {
		sb := bytes.NewBufferString("'")
		if len(re.Skills) == 0 {
			_, err := sb.WriteString(" '")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}
		_, err := sb.WriteString(re.Skills)
		_, err = sb.WriteString("'")
		if err != nil {
			resErr = gomultierror.Append(resErr, err)
		}
		return sb.String()
	}())
	resValues = append(resValues, strconv.FormatInt(int64(re.City), 10))
	resValues = append(resValues, strconv.FormatInt(int64(re.Specialization), 10))
	resValues = append(resValues, strconv.FormatInt(int64(re.Status), 10))
	_, err := sb.WriteString(strings.Join(resValues, ", "))
	_, err = sb.WriteString(");")
	if err != nil {
		resErr = gomultierror.Append(resErr, err)
	}
	if resErr != nil {
		return "", resErr
	}
	return sb.String(), nil
}
func (re *ResumeEntity) BuildUpsertQuery() (string, error) {
	var (
		resErr    error
		resValues = make([]string, 0, 30)
	)
	sb := bytes.NewBufferString("REPLACE INTO resumes(id, name, active, created_at, updated_at, job_name, job_desc_duties, job_desc_achievements, company_name, institute_name, certificate_name, customer_id, languages, salary_curr, payment_period, industries, business_trip, default_work_type, regular_trip, work_license, resume_language, metro, work_type, additional_calc_form, about_short, salary, skills, city, specialization, status) VALUES(")
	resValues = append(resValues, strconv.FormatInt(int64(re.Id), 10))
	resValues = append(resValues, func() string {
		sb := bytes.NewBufferString("'")
		if len(re.Name) == 0 {
			_, err := sb.WriteString(" '")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}
		_, err := sb.WriteString(re.Name)
		_, err = sb.WriteString("'")
		if err != nil {
			resErr = gomultierror.Append(resErr, err)
		}
		return sb.String()
	}())
	resValues = append(resValues, func() string {
		if re.Active == true {
			return "1"
		}
		return "0"
	}())
	resValues = append(resValues, strconv.FormatInt(int64(re.CreatedAt), 10))
	resValues = append(resValues, strconv.FormatInt(int64(re.UpdatedAt), 10))
	resValues = append(resValues, func() string {
		sb := bytes.NewBufferString("'")
		if len(re.JobName) == 0 {
			_, err := sb.WriteString(" '")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}
		_, err := sb.WriteString(re.JobName)
		_, err = sb.WriteString("'")
		if err != nil {
			resErr = gomultierror.Append(resErr, err)
		}
		return sb.String()
	}())
	resValues = append(resValues, func() string {
		sb := bytes.NewBufferString("'")
		if len(re.JobDescDuties) == 0 {
			_, err := sb.WriteString(" '")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}
		_, err := sb.WriteString(re.JobDescDuties)
		_, err = sb.WriteString("'")
		if err != nil {
			resErr = gomultierror.Append(resErr, err)
		}
		return sb.String()
	}())
	resValues = append(resValues, func() string {
		sb := bytes.NewBufferString("'")
		if len(re.JobDescAchievements) == 0 {
			_, err := sb.WriteString(" '")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}
		_, err := sb.WriteString(re.JobDescAchievements)
		_, err = sb.WriteString("'")
		if err != nil {
			resErr = gomultierror.Append(resErr, err)
		}
		return sb.String()
	}())
	resValues = append(resValues, func() string {
		sb := bytes.NewBufferString("'")
		if len(re.CompanyName) == 0 {
			_, err := sb.WriteString(" '")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}
		_, err := sb.WriteString(re.CompanyName)
		_, err = sb.WriteString("'")
		if err != nil {
			resErr = gomultierror.Append(resErr, err)
		}
		return sb.String()
	}())
	resValues = append(resValues, func() string {
		sb := bytes.NewBufferString("'")
		if len(re.InstituteName) == 0 {
			_, err := sb.WriteString(" '")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}
		_, err := sb.WriteString(re.InstituteName)
		_, err = sb.WriteString("'")
		if err != nil {
			resErr = gomultierror.Append(resErr, err)
		}
		return sb.String()
	}())
	resValues = append(resValues, func() string {
		sb := bytes.NewBufferString("'")
		if len(re.CertificateName) == 0 {
			_, err := sb.WriteString(" '")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}
		_, err := sb.WriteString(re.CertificateName)
		_, err = sb.WriteString("'")
		if err != nil {
			resErr = gomultierror.Append(resErr, err)
		}
		return sb.String()
	}())
	resValues = append(resValues, strconv.FormatInt(int64(re.CustomerId), 10))
	resValues = append(resValues, func() string {
		sb := bytes.NewBufferString("(")
		if len(re.Languages) == 0 {
			_, err := sb.WriteString("0)")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}
		for i, v := range re.Languages {
			if i == len(re.Languages)-1 {
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
		sb := bytes.NewBufferString("'")
		if len(re.SalaryCurr) == 0 {
			_, err := sb.WriteString(" '")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}
		_, err := sb.WriteString(re.SalaryCurr)
		_, err = sb.WriteString("'")
		if err != nil {
			resErr = gomultierror.Append(resErr, err)
		}
		return sb.String()
	}())
	resValues = append(resValues, strconv.FormatInt(int64(re.PaymentPeriod), 10))
	resValues = append(resValues, func() string {
		sb := bytes.NewBufferString("(")
		if len(re.Industries) == 0 {
			_, err := sb.WriteString("0)")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}
		for i, v := range re.Industries {
			if i == len(re.Industries)-1 {
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
	resValues = append(resValues, strconv.FormatInt(int64(re.BusinessTrip), 10))
	resValues = append(resValues, strconv.FormatInt(int64(re.DefaultWorkType), 10))
	resValues = append(resValues, strconv.FormatInt(int64(re.RegularTrip), 10))
	resValues = append(resValues, func() string {
		sb := bytes.NewBufferString("(")
		if len(re.WorkLicense) == 0 {
			_, err := sb.WriteString("0)")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}
		for i, v := range re.WorkLicense {
			if i == len(re.WorkLicense)-1 {
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
	resValues = append(resValues, strconv.FormatInt(int64(re.ResumeLanguage), 10))
	resValues = append(resValues, func() string {
		sb := bytes.NewBufferString("(")
		if len(re.Metro) == 0 {
			_, err := sb.WriteString("0)")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}
		for i, v := range re.Metro {
			if i == len(re.Metro)-1 {
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
		if len(re.WorkType) == 0 {
			_, err := sb.WriteString("0)")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}
		for i, v := range re.WorkType {
			if i == len(re.WorkType)-1 {
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
		if len(re.AdditionalCalcForm) == 0 {
			_, err := sb.WriteString("0)")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}
		for i, v := range re.AdditionalCalcForm {
			if i == len(re.AdditionalCalcForm)-1 {
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
		sb := bytes.NewBufferString("'")
		if len(re.AboutShort) == 0 {
			_, err := sb.WriteString(" '")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}
		_, err := sb.WriteString(re.AboutShort)
		_, err = sb.WriteString("'")
		if err != nil {
			resErr = gomultierror.Append(resErr, err)
		}
		return sb.String()
	}())
	resValues = append(resValues, strconv.FormatInt(int64(re.Salary), 10))
	resValues = append(resValues, func() string {
		sb := bytes.NewBufferString("'")
		if len(re.Skills) == 0 {
			_, err := sb.WriteString(" '")
			if err != nil {
				resErr = gomultierror.Append(resErr, err)
			}
			return sb.String()
		}
		_, err := sb.WriteString(re.Skills)
		_, err = sb.WriteString("'")
		if err != nil {
			resErr = gomultierror.Append(resErr, err)
		}
		return sb.String()
	}())
	resValues = append(resValues, strconv.FormatInt(int64(re.City), 10))
	resValues = append(resValues, strconv.FormatInt(int64(re.Specialization), 10))
	resValues = append(resValues, strconv.FormatInt(int64(re.Status), 10))
	_, err := sb.WriteString(strings.Join(resValues, ", "))
	_, err = sb.WriteString(");")
	if err != nil {
		resErr = gomultierror.Append(resErr, err)
	}
	if resErr != nil {
		return "", resErr
	}
	return sb.String(), nil
}
