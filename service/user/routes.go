package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"

	"github.com/AyKrimino/JobSeekerAPI/service/auth"
	"github.com/AyKrimino/JobSeekerAPI/service/company"
	"github.com/AyKrimino/JobSeekerAPI/service/jobseeker"
	"github.com/AyKrimino/JobSeekerAPI/types"
	"github.com/AyKrimino/JobSeekerAPI/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	UserRepo types.UserRepository
	JobSeekerRepo types.JobSeekerRepository
	CompanyRepo types.CompanyRepository
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{
		UserRepo: NewUserStore(db),
		JobSeekerRepo: jobseeker.NewJobseekerStore(db),
		CompanyRepo: company.NewCompany(db),
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var req types.RegisterUserRequest

	if err := utils.ParseJSON(r, &req); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(req); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid request %v", errors))
		return
	}

	_, err := h.UserRepo.GetUserByEmail(req.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", req.Email))
		return
	}

	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	req.Password = hashedPassword

	switch strings.ToLower(req.Role) {
	case "jobseeker":
		userID, err := h.UserRepo.CreateUser(parseUserFromRequest(&req))
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		if err := h.JobSeekerRepo.CreateJobSeeker(parseJobSeekerFromRequest(&req, userID)); err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}
	case "company":
		userID, err := h.UserRepo.CreateUser(parseUserFromRequest(&req))
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		if err := h.CompanyRepo.CreateCompany(parseCompanyFromRequest(&req, userID)); err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}
	default:
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid role"))
		return
	}
	utils.WriteJSON(w, http.StatusCreated, nil)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"success": "login"})
}

func parseUserFromRequest(req *types.RegisterUserRequest) *types.User {
	return &types.User{
		Email:     req.Email,
		Password:  req.Password,
		Role:      req.Role,
	}
}

func parseJobSeekerFromRequest(req *types.RegisterUserRequest, userID int) *types.JobSeeker {
	return &types.JobSeeker{
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		ProfileSummary: req.ProfileSummary,
		Skills:         req.Skills,
		Experience:     req.Experience,
		Education:      req.Education,
		UserID:         userID,
	}
}

func parseCompanyFromRequest(req *types.RegisterUserRequest, userID int) *types.Company {
	return &types.Company{
		Name:         req.Name,
		Headquarters: req.Headquarters,
		Website:      req.Website,
		Industry:     req.Industry,
		CompanySize:  req.CompanySize,
		UserID:       userID,
	}
}
