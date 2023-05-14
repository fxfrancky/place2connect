package handlers

import (
	"github.com/fxfrancky/place2connect-api/internal/models"
	"github.com/gofiber/fiber/v2"
)

// User Request
type SignUpInput struct {
	LastName      string        `json:"lastName" validate:"required,min=3"`
	FirstName     string        `json:"firstName" validate:"required,min=3"`
	Email         string        `json:"email" validate:"email,omitempty"`
	Password      string        `json:"password" validate:"required,min=8"`
	Provider      string        `json:"provider"`
	IsAdmin       bool          `json:"isAdmin"`
	Photo         string        `json:"photo,omitempty"`
	Verified      bool          `json:"isVerified,omitempty"`
	PicturePath   string        `json:"picturePath" validate:"required,min=3"`
	Friends       []models.User `json:"friends,omitempty"`
	Location      string        `json:"location" validate:"required,min=3"`
	Occupation    string        `json:"occupation" validate:"required,min=3"`
	ViewedProfile uint          `json:"viewedProfile"`
	Impressions   uint          `json:"impressions"`
}

type SignInInput struct {
	Email    string `json:"email"  validate:"required"`
	Password string `json:"password"  validate:"required"`
}

// User Request
type UpdateUserRequest struct {
	FirstName     string        `json:"firstName" validate:"required"`
	LastName      string        `json:"lastName" validate:"required"`
	Email         string        `json:"email" validate:"email,omitempty"`
	Role          string        `json:"role,omitempty"`
	Photo         string        `json:"photo,omitempty"`
	Provider      string        `json:"provider"`
	IsAdmin       bool          `json:"isAdmin"`
	Password      string        `json:"password" validate:"required,min=8"`
	Verified      bool          `json:"isVerified,omitempty"`
	PicturePath   string        `json:"picturePath" validate:"required,min=3"`
	Friends       []models.User `json:"friends,omitempty"`
	Location      string        `json:"location" validate:"required,min=3"`
	Occupation    string        `json:"occupation" validate:"required,min=3"`
	ViewedProfile uint          `json:"viewedProfile"`
	Impressions   uint          `json:"impressions"`
}

func (u *UpdateUserRequest) BindUpdateUser(c *fiber.Ctx, usr *models.User) error {

	// Validate the user
	if err := c.BodyParser(u); err != nil {
		return err
	}

	// Map the user
	usr.FirstName = u.FirstName
	usr.LastName = u.LastName
	usr.Email = u.Email
	usr.IsAdmin = u.IsAdmin
	usr.Password = u.Password
	usr.Provider = u.Provider
	usr.Role = u.Role
	usr.Photo = u.Photo
	usr.Verified = u.Verified
	usr.PicturePath = u.PicturePath
	usr.Location = u.Location
	usr.Occupation = u.Occupation
	usr.ViewedProfile = u.ViewedProfile
	usr.Impressions = u.Impressions

	return nil
}

func (s *UpdateUserRequest) PopulateUpdateUser(user *models.User) {

	// Populate the user
	s.FirstName = user.FirstName
	s.LastName = user.LastName
	s.Email = user.Email
	s.IsAdmin = user.IsAdmin
	s.Password = user.Password
	s.Provider = user.Provider
	s.Role = user.Role
	s.Photo = user.Photo
	s.Verified = user.Verified
	s.PicturePath = user.PicturePath
	s.Location = user.Location
	s.Occupation = user.Occupation
	s.ViewedProfile = user.ViewedProfile
	s.Impressions = user.Impressions
}
