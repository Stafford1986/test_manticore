// Code generated by generator, DO NOT EDIT.
package pb

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func (re *ResumeEntity) ParseDbResult(m map[string]interface{}) (*ResumeEntity, error) {
	var (
		res = &ResumeEntity{}
	)
	for k, v := range m {
		switch k {
		case "id":
			val, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err convert id")
			}
			p, err := strconv.ParseUint(string(val), 10, 32)
			if err != nil {
				return nil, errors.New("err convert value to Id")
			}
			res.Id = uint32(p)
		case "name":
			val, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err convert name")
			}
			res.Name = string(val)
		case "active":
			val, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err convert active")
			}
			p, err := strconv.Atoi(string(val))
			if err != nil {
				return nil, errors.New("err convert value to Active")
			}
			if p == 0 {
				res.Active = false
			} else {
				res.Active = true
			}
		case "created_at":
			val, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err convert created_at")
			}
			p, err := strconv.ParseInt(string(val), 10, 64)
			if err != nil {
				return nil, errors.New("err convert value to CreatedAt")
			}
			res.CreatedAt = p
		case "updated_at":
			val, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err convert updated_at")
			}
			p, err := strconv.ParseInt(string(val), 10, 64)
			if err != nil {
				return nil, errors.New("err convert value to UpdatedAt")
			}
			res.UpdatedAt = p
		case "job_name":
			val, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err convert job_name")
			}
			res.JobName = string(val)
		case "job_desc_duties":
			val, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err convert job_desc_duties")
			}
			res.JobDescDuties = string(val)
		case "job_desc_achievements":
			val, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err convert job_desc_achievements")
			}
			res.JobDescAchievements = string(val)
		case "company_name":
			val, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err convert company_name")
			}
			res.CompanyName = string(val)
		case "institute_name":
			val, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err convert institute_name")
			}
			res.InstituteName = string(val)
		case "certificate_name":
			val, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err convert certificate_name")
			}
			res.CertificateName = string(val)
		case "customer_id":
			val, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err convert customer_id")
			}
			p, err := strconv.ParseUint(string(val), 10, 32)
			if err != nil {
				return nil, errors.New("err convert value to CustomerId")
			}
			res.CustomerId = uint32(p)
		case "languages":
			val, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err convert languages")
			}
			var (
				r []uint32
			)
			s := strings.Split(string(val), ",")
			for _, sv := range s {
				p, err := strconv.ParseUint(string(sv), 10, 32)
				if err != nil {
					return nil, errors.New("err convert value to Languages")
				}
				r = append(r, uint32(p))
			}
			res.Languages = r
		case "salary_curr":
			val, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err convert salary_curr")
			}
			res.SalaryCurr = string(val)
		case "payment_period":
			val, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err convert payment_period")
			}
			p, err := strconv.ParseUint(string(val), 10, 32)
			if err != nil {
				return nil, errors.New("err convert value to PaymentPeriod")
			}
			res.PaymentPeriod = uint32(p)
		case "industries":
			val, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err convert industries")
			}
			var (
				r []uint32
			)
			s := strings.Split(string(val), ",")
			for _, sv := range s {
				p, err := strconv.ParseUint(string(sv), 10, 32)
				if err != nil {
					return nil, errors.New("err convert value to Industries")
				}
				r = append(r, uint32(p))
			}
			res.Industries = r
		case "business_trip":
			val, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err convert business_trip")
			}
			p, err := strconv.ParseUint(string(val), 10, 32)
			if err != nil {
				return nil, errors.New("err convert value to BusinessTrip")
			}
			res.BusinessTrip = uint32(p)
		case "default_work_type":
			val, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err convert default_work_type")
			}
			p, err := strconv.ParseUint(string(val), 10, 32)
			if err != nil {
				return nil, errors.New("err convert value to DefaultWorkType")
			}
			res.DefaultWorkType = uint32(p)
		case "regular_trip":
			val, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err convert regular_trip")
			}
			p, err := strconv.ParseUint(string(val), 10, 32)
			if err != nil {
				return nil, errors.New("err convert value to RegularTrip")
			}
			res.RegularTrip = uint32(p)
		case "work_license":
			val, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err convert work_license")
			}
			var (
				r []uint32
			)
			s := strings.Split(string(val), ",")
			for _, sv := range s {
				p, err := strconv.ParseUint(string(sv), 10, 32)
				if err != nil {
					return nil, errors.New("err convert value to WorkLicense")
				}
				r = append(r, uint32(p))
			}
			res.WorkLicense = r
		case "resume_language":
			val, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err convert resume_language")
			}
			p, err := strconv.ParseUint(string(val), 10, 32)
			if err != nil {
				return nil, errors.New("err convert value to ResumeLanguage")
			}
			res.ResumeLanguage = uint32(p)
		case "metro":
			val, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err convert metro")
			}
			var (
				r []uint32
			)
			s := strings.Split(string(val), ",")
			for _, sv := range s {
				p, err := strconv.ParseUint(string(sv), 10, 32)
				if err != nil {
					return nil, errors.New("err convert value to Metro")
				}
				r = append(r, uint32(p))
			}
			res.Metro = r
		case "work_type":
			val, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err convert work_type")
			}
			var (
				r []uint32
			)
			s := strings.Split(string(val), ",")
			for _, sv := range s {
				p, err := strconv.ParseUint(string(sv), 10, 32)
				if err != nil {
					return nil, errors.New("err convert value to WorkType")
				}
				r = append(r, uint32(p))
			}
			res.WorkType = r
		case "additional_calc_form":
			val, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err convert additional_calc_form")
			}
			var (
				r []uint32
			)
			s := strings.Split(string(val), ",")
			for _, sv := range s {
				p, err := strconv.ParseUint(string(sv), 10, 32)
				if err != nil {
					return nil, errors.New("err convert value to AdditionalCalcForm")
				}
				r = append(r, uint32(p))
			}
			res.AdditionalCalcForm = r
		case "about_short":
			val, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err convert about_short")
			}
			res.AboutShort = string(val)
		case "salary":
			val, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err convert salary")
			}
			p, err := strconv.ParseUint(string(val), 10, 32)
			if err != nil {
				return nil, errors.New("err convert value to Salary")
			}
			res.Salary = uint32(p)
		case "skills":
			val, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err convert skills")
			}
			res.Skills = string(val)
		case "city":
			val, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err convert city")
			}
			p, err := strconv.ParseUint(string(val), 10, 32)
			if err != nil {
				return nil, errors.New("err convert value to City")
			}
			res.City = uint32(p)
		case "specialization":
			val, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err convert specialization")
			}
			p, err := strconv.ParseUint(string(val), 10, 32)
			if err != nil {
				return nil, errors.New("err convert value to Specialization")
			}
			res.Specialization = uint32(p)
		case "status":
			val, ok := v.([]byte)
			if !ok {
				return nil, errors.New("err convert status")
			}
			p, err := strconv.ParseUint(string(val), 10, 32)
			if err != nil {
				return nil, errors.New("err convert value to Status")
			}
			res.Status = uint32(p)
		default:
			return nil, fmt.Errorf("unknown field %s", k)
		}
	}
	return res, nil
}
