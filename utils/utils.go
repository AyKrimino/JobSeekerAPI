package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/AyKrimino/JobSeekerAPI/types"
	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()
var IsAlpha = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

func init() {
	Validate.RegisterStructValidation(ValidateRegisterUserRequest, types.RegisterUserRequest{})
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}

func ParseJSON(r *http.Request, v any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(v)
}

func EncodeStringSliceToJSON(s []string) ([]byte, error) {
	jsonData, err := json.Marshal(s)
	if err != nil {
		return nil, fmt.Errorf("failed to encode string slice to JSON: %w", err)
	}
	return jsonData, nil
}

func DecodeJSONTOStringSlice(jsonData []byte) ([]string, error) {
	var result []string

	if err := json.Unmarshal(jsonData, &result); err != nil {
		return nil, fmt.Errorf("failed to decode JSON data to string slice: %w", err)
	}
	return result, nil
}

func IsValidURL(u string) bool {
	_, err := url.ParseRequestURI(u)
	return err != nil
}

func ValidateRegisterUserRequest(sl validator.StructLevel) {
	req := sl.Current().Interface().(types.RegisterUserRequest)

	if strings.EqualFold(req.Role, "JobSeeker") {
		if req.FirstName == "" {
			sl.ReportError(req.FirstName, "FirstName", "firstName", "required_for_jobseeker", "")
		}
		if !IsAlpha(req.FirstName) {
			sl.ReportError(req.FirstName, "FirstName", "firstName", "firstName_not_alpha", "")
		}
		if len(req.FirstName) > 100 {
			sl.ReportError(
				req.FirstName,
				"FirstName",
				"firstName",
				"firstName_length_must_be_lte_100",
				"",
			)
		}

		if req.LastName == "" {
			sl.ReportError(req.LastName, "LastName", "lastName", "required_for_jobseeker", "")
		}
		if !IsAlpha(req.LastName) {
			sl.ReportError(req.LastName, "LastName", "lastName", "lastName_not_alpha", "")
		}
		if len(req.LastName) > 100 {
			sl.ReportError(
				req.LastName,
				"LastName",
				"lastName",
				"lastName_length_must_be_lte_100",
				"",
			)
		}

		if len(req.ProfileSummary) > 500 {
			sl.ReportError(
				req.ProfileSummary,
				"ProfileSummary",
				"profileSummary",
				"profileSummary_length_must_be_lte_500",
				"",
			)
		}

		if req.Experience > 50 || req.Experience < 0 {
			sl.ReportError(
				req.Experience,
				"Experience",
				"experience",
				"experience_must_be_between_0_and_50",
				"",
			)
		}

		if len(req.Education) > 255 {
			sl.ReportError(
				req.Education,
				"Education",
				"education",
				"education_length_must_be_lte_255",
				"",
			)
		}

		if req.Name != "" {
			sl.ReportError(req.Name, "Name", "name", "forbidden_for_jobseeker", "")
		}
		if req.Headquarters != "" {
			sl.ReportError(
				req.Headquarters,
				"Headquarters",
				"headquarters",
				"forbidden_for_jobseeker",
				"",
			)
		}
		if req.Website != "" {
			sl.ReportError(req.Website, "Website", "website", "forbidden_for_jobseeker", "")
		}
		if req.Industry != "" {
			sl.ReportError(req.Industry, "Industry", "industry", "forbidden_for_jobseeker", "")
		}
		if req.CompanySize != "" {
			sl.ReportError(
				req.CompanySize,
				"CompanySize",
				"companySize",
				"forbidden_for_jobseeker",
				"",
			)
		}
	}
	if strings.EqualFold(req.Role, "Company") {
		if req.Name == "" {
			sl.ReportError(req.Name, "Name", "name", "required_for_company", "")
		}
		if !IsAlpha(req.Name) {
			sl.ReportError(req.Name, "Name", "name", "name_not_alpha", "")
		}
		if len(req.Name) > 255 {
			sl.ReportError(req.Name, "Name", "name", "name_length_must_be_lte_255", "")
		}

		if len(req.Headquarters) > 255 {
			sl.ReportError(
				req.Headquarters,
				"Headquarters",
				"headquarters",
				"headquarters_length_must_be_lte_255",
				"",
			)
		}

		if len(req.Website) > 255 {
			sl.ReportError(req.Website, "Website", "website", "website_length_must_be_lte_255", "")
		}
		if !IsValidURL(req.Website) {
			sl.ReportError(req.Website, "Website", "website", "website_must_be_in_url_format", "")
		}

		if len(req.Industry) > 255 {
			sl.ReportError(
				req.Industry,
				"Industry",
				"industry",
				"industry_length_must_be_lte_255",
				"",
			)
		}

		if len(req.CompanySize) > 255 {
			sl.ReportError(
				req.CompanySize,
				"CompanySize",
				"companySize",
				"companySize_length_must_be_lte_255",
				"",
			)
		}

		if req.FirstName != "" {
			sl.ReportError(req.FirstName, "FirstName", "firstName", "forbidden_for_company", "")
		}
		if req.LastName != "" {
			sl.ReportError(req.LastName, "LastName", "lastName", "forbidden_for_company", "")
		}
		if req.ProfileSummary != "" {
			sl.ReportError(
				req.ProfileSummary,
				"ProfileSummary",
				"profileSummary",
				"forbidden_for_company",
				"",
			)
		}
		if req.Skills != nil {
			sl.ReportError(req.Skills, "Skills", "skills", "forbidden_for_company", "")
		}
		if req.Experience != 0 {
			sl.ReportError(req.Experience, "Experience", "experience", "forbidden_for_company", "")
		}
		if req.Education != "" {
			sl.ReportError(req.Education, "Education", "education", "forbidden_for_company", "")
		}
	}
}
