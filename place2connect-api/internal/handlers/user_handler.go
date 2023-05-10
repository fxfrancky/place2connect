package handlers

import (
	"errors"
	"net/http"
	"strconv"

	handlersRequests "github.com/fxfrancky/place2connect-api/internal/handlers/requests"
	handlersResponses "github.com/fxfrancky/place2connect-api/internal/handlers/responses"
	"github.com/fxfrancky/place2connect-api/internal/models"
	"github.com/fxfrancky/place2connect-api/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

// GetMe func to get current user.
// @Description Get current user.
// @Summary Logout  of the the API.
// @Tags User
// @Accept json
// @Produce json
// @Param	Authorization	header		string	true	"Authentication header"
// @Security ApiKeyAuth
// @Router /api/v1/users/me [get]
func (h *Handler) GetMe(c *fiber.Ctx) error {
	user := c.Locals("user").(handlersResponses.UserResponse)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": user}})
}

// GetUser func to get a user By Email
// @Summary Get a user
// @Description Get a user. Auth required
// @ID get-user-by-email
// @Tags User
// @Accept  json
// @Produce  json
// @Param email path string true "Email of the user to get"
// @Param	Authorization	header		string	true	"Authentication header"
// @Success 200 {object} handlersResponses.UserResponse
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		500				{string}	string	"Status Internal Server Error"
// @Security ApiKeyAuth
// @Router /api/v1/users/email/{email} [get]
func (h *Handler) GetUserByEmail(c *fiber.Ctx) error {

	email := c.Params("email")

	u, err := h.userRepository.GetUserByEmail(email)

	if u == nil && err == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound("User"))
	}
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	return c.Status(http.StatusOK).JSON(handlersResponses.NewUserResponse(u))
}

// GetUser func to get a user By ID
// @Summary Get a user ID
// @Description Get a user By ID. Auth required
// @ID get-user-by-id
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path integer true "Id of the user to get"
// @Param	Authorization	header		string	true	"Authentication header"
// @Success 200 {object} handlersResponses.UserResponse
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		500				{string}	string	"Status Internal Server Error"
// @Security ApiKeyAuth
// @Router /api/v1/users/{id} [get]
func (h *Handler) GetUserById(c *fiber.Ctx) error {

	id, err := utils.StringToInt(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	u, err := h.userRepository.GetUserById(id)

	if u == nil && err == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound("User"))
	}
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	return c.Status(http.StatusOK).JSON(handlersResponses.NewUserResponse(u))
}

// UpdateUser func to update a user
// @Summary Update a user
// @Description Update a User. Auth is required
// @ID update-user
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path string true "Id of the user to update"
// @Param user body handlersRequests.UpdateUserRequest true "User to update"
// @Param	Authorization	header		string	true	"Authentication header"
// @Success   200       {object} handlersResponses.UserResponse
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		403				{string}	string	"Status Forbidden"
// @Failure		404				{string}	string	"Status Not Found"
// @Failure		409				{string}	string	"Status Conflict"
// @Failure		422				{string}	string	"Status UnprocessableEntity"
// @Failure		502				{string}	string	"Status BadGateway"
// @Security ApiKeyAuth
// @Router /api/v1/users/{id} [put]
func (h *Handler) UpdateUser(c *fiber.Ctx) error {

	id, err := utils.StringToInt(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	u, err := h.userRepository.GetUserById(id)

	if u == nil && err == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound("User"))
	}
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	payload := &handlersRequests.UpdateUserRequest{}
	payload.PopulateUpdateUser(u)
	if err := payload.BindUpdateUser(c, u); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}
	errors := utils.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})
	}

	if err = h.userRepository.UpdateUser(u); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	return c.Status(http.StatusOK).JSON(handlersResponses.NewUserResponse(u))
}

// DeleteUser func to delete a User
// @Summary Delete a User
// @Description Delete a User. Auth is required
// @ID delete-user
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path string true "Id of the user to delete"
// @Param	Authorization	header		string	true	"Authentication header"
// @Success   200       {object} handlersResponses.UserResponse
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		403				{string}	string	"Status Forbidden"
// @Failure		404				{string}	string	"Status Not Found"
// @Failure		409				{string}	string	"Status Conflict"
// @Failure		422				{string}	string	"Status UnprocessableEntity"
// @Failure		502				{string}	string	"Status BadGateway"
// @Security ApiKeyAuth
// @Router /api/v1/auth/users/{id} [delete]
func (h *Handler) DeleteUser(c *fiber.Ctx) error {

	id, err := utils.StringToInt(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	u, err := h.userRepository.GetUserById(id)

	if u == nil && err == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound("User"))
	}
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	err = h.userRepository.DeleteUser(u)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	return c.Status(http.StatusOK).JSON(map[string]interface{}{"result": "user deleted!"})
}

// AllUser func to display all User
// @Summary Display all Users
// @Description Display all Users. Auth is required
// @ID all-users
// @Tags User
// @Accept  json
// @Produce  json
// @Param limit query integer false "Limit number of products returned (default is 20)"
// @Param offset query integer false "Offset/Skip number of products (default is 0)"
// @Param	Authorization	header		string	true	"Authentication header"
// @Success   200       {object} handlersResponses.UserListResponse
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		500				{string}	string	"Status Internal Server Error"
// @Failure		502				{string}	string	"Status BadGateway"
// @Security ApiKeyAuth
// @Router /api/v1/auth/users/{limit}/{offset} [get]
func (h *Handler) AllUsers(c *fiber.Ctx) error {
	var (
		users []models.User
		count int64
	)
	offset, err := strconv.Atoi(c.Params("offset"))
	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(c.Params("limit"))
	if err != nil {
		limit = 20
	}

	users, count, err = h.userRepository.AllUsers(offset, limit)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	return c.Status(http.StatusOK).JSON(handlersResponses.NewUserListResponse(users, count))
}

// AllUserFriends func to display all Friends of a User
// @Summary Display all Friends of a User
// @Description Display all Friends of a User By userId. Auth is required
// @ID all-user-friends
// @Tags User
// @Accept  json
// @Produce  json
// @Param userID path integer true "ID of the user for friends"
// @Param limit query integer false "Limit number of Friends returned (default is 20)"
// @Param offset query integer false "Offset/Skip number of friends (default is 0)"
// @Param	Authorization	header		string	true	"Authentication header"
// @Success   200       {object}  handlersResponses.UserListResponse
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		500				{string}	string	"Status Internal Server Error"
// @Failure		502				{string}	string	"Status BadGateway"
// @Security ApiKeyAuth
// @Router /api/v1/users/{userID}/{limit}/{offset} [get]
func (h *Handler) AllUserFriends(c *fiber.Ctx) error {
	var (
		friends []models.User
		count   int64
	)

	limit, err := strconv.Atoi(c.Params("limit"))
	if err != nil {
		limit = 20
		err = nil
	}

	offset, err := strconv.Atoi(c.Params("offset"))
	if err != nil {
		offset = 0
		err = nil
	}

	// Get The user
	userId, err := utils.StringToInt(c.Params("userID"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.BadRequest())
	}

	user, err := h.userRepository.GetUserById(userId)
	if user == nil && err == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound("User"))
	}
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	friends, count, err = h.userRepository.GetUserFriends(userId, offset, limit)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	return c.Status(http.StatusOK).JSON(handlersResponses.NewUserListResponse(friends, count))
}

// AddFriendToUser func to add a new Friend to User
// @Summary Add a new Friend to a User
// @Description AddFriendToUser add a new friend to User
// @ID add-friend
// @Tags User
// @Accept  json
// @Produce  json
// @Param userFriend body handlersRequests.AddUserFriendRequest true "Add User to Friend"
// @Param	Authorization	header		string	true	"Authentication header"
// @Success 201 {object} handlersResponses.UserResponse
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		403				{string}	string	"Status Forbidden"
// @Failure		409				{string}	string	"Status Conflict"
// @Failure		422				{string}	string	"Status UnprocessableEntity"
// @Failure		502				{string}	string	"Status BadGateway"
// @Security ApiKeyAuth
// @Router /api/v1/userfriends [post]
func (h *Handler) AddFriendToUser(c *fiber.Ctx) error {

	// Get The user
	var payload *handlersRequests.AddUserFriendRequest
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.NewError(err))
	}

	userId := int(payload.UserID)
	friendId := int(payload.FriendID)

	user, err := h.userRepository.GetUserById(userId)
	if user == nil && err == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound("User"))
	}
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	userFriend, err := h.userRepository.GetUserById(friendId)
	if userFriend == nil && err == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound("User"))
	}
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	if friendId == userId {
		msg := "cannot add a user as himself own friend"
		return c.Status(http.StatusBadRequest).JSON(utils.NewError(errors.New(msg)))
	}

	err = h.userRepository.AddUserFriend(user, userFriend)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}
	err = h.userRepository.AddUserFriend(userFriend, user)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}
	return c.Status(http.StatusCreated).JSON(handlersResponses.NewUserResponse(user))
}

// RemoveFriendToUser func to remove a Friend to User
// @Summary Remove a new Friend to a User
// @Description RemoveFriendToUser remove a friend to a User
// @ID remove-friend
// @Tags User
// @Accept  json
// @Produce  json
// @Param userID path string true "Id of the User to remove the friend"
// @Param friendID path string true "Id of the UserFriend to remove"
// @Param	Authorization	header		string	true	"Authentication header"
// @Success 201 {object} handlersResponses.UserResponse
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		403				{string}	string	"Status Forbidden"
// @Failure		409				{string}	string	"Status Conflict"
// @Failure		422				{string}	string	"Status UnprocessableEntity"
// @Failure		502				{string}	string	"Status BadGateway"
// @Security ApiKeyAuth
// @Router /api/v1/userfriends/ [delete]
func (h *Handler) RemoveFriendToUser(c *fiber.Ctx) error {

	// Get The user
	var payload *handlersRequests.AddUserFriendRequest
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.NewError(err))
	}

	userId := int(payload.UserID)
	friendId := int(payload.FriendID)

	user, err := h.userRepository.GetUserById(userId)
	if user == nil && err == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound("User"))
	}
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	userFriend, err := h.userRepository.GetUserById(friendId)
	if userFriend == nil && err == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound("User"))
	}
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	err = h.userRepository.RemoveUserFriend(user, userFriend)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}
	err = h.userRepository.RemoveUserFriend(userFriend, user)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}
	return c.Status(http.StatusCreated).JSON(handlersResponses.NewUserResponse(user))
}
