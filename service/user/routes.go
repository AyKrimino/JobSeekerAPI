package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/AyKrimino/JobSeekerAPI/types"
	"github.com/AyKrimino/JobSeekerAPI/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(s types.UserStore) *Handler {
	return &Handler{store: s}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var req types.RegisterUserRequest

	// parse json data into req
	if err := utils.ParseJSON(r, &req); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// data validation

	// check if email already exists
	_, err := h.store.GetUserByEmail(req.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", req.Email))
		return
	}

	// create jobseeker or company
	switch req.Role {
	case "JobSeeker":
		userID, err := h.store.CreateUser(parseUserFromRequest(&req))
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		if err := h.store.CreateJobSeeker(parseJobSeekerFromRequest(&req, userID)); err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}
	case "Company":
		userID, err := h.store.CreateUser(parseUserFromRequest(&req))
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		if err := h.store.CreateCompany(parseCompanyFromRequest(&req, userID)); err != nil {
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
		IsActive:  true,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
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
