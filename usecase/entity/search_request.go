package entity

import "strings"

const (
	SortCreatedAt = "created_at"
	SortUpdatedAt = "updated_at"

	orderAsc  = "ASC"
	orderDesc = "DESC"

	filterBySalaryMin  = "salary_min"
	filterBySalaryMax  = "salary_max"
	filterBySalaryFrom = "salary_from"
	filterBySalaryTo   = "salary_to"
)

var (
	forbiddenChars = []rune{'/', '(', ')', ';', '"', '<', '>', '\\', '{', '}'}
)

//easyjson:json
type SearchRequest struct {
	Query        string        `json:"query"`
	IsActive     bool          `json:"is_active"`
	CityId       int32         `json:"city_id"`
	SortParams   *SortParams   `json:"sort_params"`
	FilterParams *FilterParams `json:"filter_params"`
}

//easyjson:json
type SortParams struct {
	Field string `json:"field"`
	Order string `json:"order"`
}

//easyjson:json
type FilterParams struct {
	Field string `json:"field"`
	Order string `json:"order"`
}

func (sr *SearchRequest) Validate() error {
	if sr.Query == "" {
		return ErrValidateSearchReq
	}

	for _, r := range forbiddenChars {
		if strings.ContainsRune(sr.Query, r) {
			return ErrForbiddenCharacter
		}
	}

	if sr.SortParams != nil {
		switch strings.ToLower(sr.SortParams.Field) {
		case SortCreatedAt, SortUpdatedAt:
		default:
			return ErrValidateSearchReq
		}

		switch strings.ToUpper(sr.SortParams.Order) {
		case orderAsc, orderDesc:
		default:
			return ErrValidateSearchReq
		}
	}

	if sr.FilterParams != nil {
		switch strings.ToLower(sr.FilterParams.Field) {
		case filterBySalaryMax, filterBySalaryMin:
		case filterBySalaryFrom, filterBySalaryTo:
		default:
			return ErrValidateSearchReq
		}

		switch strings.ToUpper(sr.FilterParams.Order) {
		case orderAsc, orderDesc:
		default:
			return ErrValidateSearchReq
		}

	}

	return nil
}
