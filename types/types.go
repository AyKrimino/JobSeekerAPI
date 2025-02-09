package types

import "time"

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type JobSeeker struct {
	ID             int      `json:"id"`
	FirstName      string   `json:"firstName"`
	LastName       string   `json:"lastName"`
	ProfileSummary string   `json:"profileSummary"`
	Skills         []string `json:"skills"`
	Experience     int      `json:"experience"`
	Education      string   `json:"education"`
	UserID         int      `json:"userId"`
}

type Company struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Headquarters string `json:"headquarters"`
	Website      string `json:"website"`
	Industry     string `json:"industry"`
	CompanySize  string `json:"companySize"`
	UserID       int    `json:"userId"`
}

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(id int) (*User, error)
	CreateUser(u *User) (id int, err error)
	CreateJobSeeker(js *JobSeeker) error
	CreateCompany(cpy *Company) error
	GetUserRoleById(id int) (string, error)
}

type RegisterUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`

	// JobSeeker-specific fields
	FirstName      string   `json:"firstName,omitempty"`
	LastName       string   `json:"lastName,omitempty"`
	ProfileSummary string   `json:"profileSummary,omitempty"`
	Skills         []string `json:"skills,omitempty"`
	Experience     int      `json:"experience,omitempty"`
	Education      string   `json:"education,omitempty"`

	// Company-specific fields
	Name         string `json:"name,omitempty"`
	Headquarters string `json:"headquarters,omitempty"`
	Website      string `json:"website,omitempty"`
	Industry     string `json:"industry,omitempty"`
	CompanySize  string `json:"companySize,omitempty"`
}
