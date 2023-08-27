package api

import (
	"database/sql"
	"github.com/accmeboot/issueshift/internal/api/response"
	"github.com/accmeboot/issueshift/internal/service"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(us *service.UserService) *UserHandler {
	return &UserHandler{userService: us}
}

func (uh *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	mockEmail := "mis@ff.com"
	user, err := uh.userService.GetUserByEmail(mockEmail)
	if err != nil {
		response.PassToContext(r, err)
		return
	}

	var DTO struct {
		ID        int64          `json:"id"`
		Email     string         `json:"email"`
		Name      string         `json:"name"`
		CreatedAt sql.NullTime   `json:"created_at"`
		AvatarUrl sql.NullString `json:"avatar_url,omitempty"`
	}

	DTO.ID = user.ID
	DTO.Email = user.Email
	DTO.Name = user.Name
	DTO.CreatedAt = user.CreatedAt
	DTO.AvatarUrl = user.AvatarUrl

	response.WriteJSON(w, r, http.StatusOK, response.Envelope{"user": DTO})
}

func (uh *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var DTO struct {
		Email    string `json:"email"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	err := response.ReadJSON(w, r, &DTO)
	if err != nil {
		response.PassToContext(r, err)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(DTO.Password), 12)
	if err != nil {
		response.PassToContext(r, err)
		return
	}
	err = uh.userService.CreateUser(DTO.Email, DTO.Name, "", hash)
	if err != nil {
		response.PassToContext(r, err)
		return
	}

	response.WriteJSON(w, r, http.StatusCreated, response.Envelope{})
}
