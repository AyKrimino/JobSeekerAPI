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
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=200"`
	Role     string `json:"role" validate:"required,oneofci=JobSeeker Company"`

	// JobSeeker-specific fields
	FirstName      string   `json:"firstName,omitempty" validate:"alpha"`
	LastName       string   `json:"lastName,omitempty" validate:"alpha"`
	ProfileSummary string   `json:"profileSummary,omitempty" validate:"max=500"`
	Skills         []string `json:"skills,omitempty"`
	Experience     int      `json:"experience,omitempty" validate:"gte=0,lte=50"`
	Education      string   `json:"education,omitempty" validate:"max=255"`

	// Company-specific fields
	Name         string `json:"name,omitempty" validate:"alpha,max=255"`
	Headquarters string `json:"headquarters,omitempty" validate:"max=255"`
	Website      string `json:"website,omitempty" validate:"max=255,url"`
	Industry     string `json:"industry,omitempty" validate:"max=255"`
	CompanySize  string `json:"companySize,omitempty" validate:"max=50"`
}
