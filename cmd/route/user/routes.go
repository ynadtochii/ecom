package user

import (
	"net/http"
	"strconv"

	"github.com/ynadtochii/ecom/db/models"
	"github.com/ynadtochii/ecom/services"

	util "github.com/ynadtochii/ecom/util"
)

type Handler struct {
	userService *services.UserService
}

func NewHandler(userService *services.UserService) *Handler {
	return &Handler{userService: userService}
}

func (h *Handler) GetUsers(router *http.ServeMux) {
	router.HandleFunc("/users", h.getUsers)
}

func (h *Handler) GetUserById(router *http.ServeMux) {
	router.HandleFunc("/users/{id}", h.getUserById)
}

func (h *Handler) CreateUser(router *http.ServeMux) {
	router.HandleFunc("POST /users", h.createUser)
}

func (h *Handler) UpdateUser(router *http.ServeMux) {
	router.HandleFunc("PUT /users/{id}", h.updateUser)
}

func (h *Handler) DeleteUser(router *http.ServeMux) {
	router.HandleFunc("DELETE /users/{id}", h.deleteUser)
}

func (h *Handler) deleteUser(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("id")

	userIDUint, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		util.RespondError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

  user,  err := h.userService.DeleteUser(uint(userIDUint))
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, "Failed to delete user")
		return
	}

	util.RespondJSON(w, http.StatusOK, user)
}

func (h *Handler) updateUser(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("id")

	var updatedUser models.User
	err := util.DecodeJSONBody(r, &updatedUser)
	if err != nil {
		util.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	userIDUint, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		util.RespondError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	updatedUser.ID = uint(userIDUint)

	user, err := h.userService.UpdateUser(&updatedUser)
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, "Failed to update user")
		return
	}

	util.RespondJSON(w, http.StatusOK, user)
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := util.DecodeJSONBody(r, &newUser)
	if err != nil {
		util.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	createdUser, err := h.userService.CreateUser(&newUser)
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	util.RespondJSON(w, http.StatusCreated, createdUser)
}

func (h *Handler) getUserById(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("id")

	userIDUint, err := strconv.ParseUint(userID, 10, 64)

	if err != nil {
		util.RespondError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	user, err := h.userService.GetUserByID(uint(userIDUint))

	if err != nil {
		util.RespondError(w, http.StatusNotFound, "User not found")
		return
	}
	util.RespondJSON(w, http.StatusOK, user)
}

func (h *Handler) getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, "Failed to get users")
		return
	}
	util.RespondJSON(w, http.StatusOK, users)
}
