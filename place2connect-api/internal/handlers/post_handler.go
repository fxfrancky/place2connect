package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	handlersRequests "github.com/fxfrancky/place2connect-api/internal/handlers/requests"
	handlersResponses "github.com/fxfrancky/place2connect-api/internal/handlers/responses"
	"github.com/fxfrancky/place2connect-api/internal/models"
	"github.com/fxfrancky/place2connect-api/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// CreatePost func to create a new Post
// @Summary create a new Post
// @Description CreatePost create a new Post
// @ID create-post
// @Tags Post
// @Accept  json
// @Produce  json
// @Param Post body PostRequest true "The PostRequest"
// @Param	Authorization	header		string	true	"Authentication header"
// @Success 201 {object} handlersResponses.PostResponse
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		403				{string}	string	"Status Forbidden"
// @Failure		409				{string}	string	"Status Conflict"
// @Failure		422				{string}	string	"Status UnprocessableEntity"
// @Failure		502				{string}	string	"Status BadGateway"
// @Security ApiKeyAuth
// @Router /api/v1/posts [post]
func (h *Handler) CreatePost(c *fiber.Ctx) error {

	var post models.Post
	payload := &handlersRequests.PostRequest{}

	if err := handlersRequests.BindPostsRequest(c, &post, payload); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}
	errors := utils.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})
	}

	user, err := h.userRepository.GetUserById(int(post.UserID))
	if user == nil && err == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound("User"))
	}
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	image := payload.PicturePath

	err = h.postRepository.CreatePost(user, post.Description, image)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}

	postResp := handlersResponses.NewPostResponse(&post)

	return c.Status(http.StatusCreated).JSON(postResp)
}

// FeedPosts func to display all Post
// @Summary Display all Post
// @Description Display all Post. Auth is required
// @ID all-posts
// @Tags Post
// @Accept  json
// @Produce  json
// @Param limit query integer false "Limit number of posts returned (default is 20)"
// @Param offset query integer false "Offset/Skip number of posts (default is 0)"
// @Param	Authorization	header		string	true	"Authentication header"
// @Success   200       {object} handlersResponses.PostListResponse
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		500				{string}	string	"Status Internal Server Error"
// @Failure		502				{string}	string	"Status BadGateway"
// @Security ApiKeyAuth
// @Router /api/v1/posts/{limit}/{offset} [get]
func (h *Handler) FeedPosts(c *fiber.Ctx) error {
	var (
		posts []models.Post
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

	posts, count, err = h.postRepository.FeedPosts(offset, limit)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	return c.Status(http.StatusOK).JSON(handlersResponses.NewPostsListResponse(posts, count))
}

// AllUserPosts func to display all Posts of a User
// @Summary Display all Posts of a User
// @Description Display all Posts of a User By userId. Auth is required
// @ID all-user-posts
// @Tags User
// @Accept  json
// @Produce  json
// @Param user_id path integer true "ID of the user for posts"
// @Param limit query integer false "Limit number of Posts returned (default is 20)"
// @Param offset query integer false "Offset/Skip number of posts (default is 0)"
// @Param	Authorization	header		string	true	"Authentication header"
// @Success   200       {object} handlersResponses.PostListResponse
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		500				{string}	string	"Status Internal Server Error"
// @Failure		502				{string}	string	"Status BadGateway"
// @Security ApiKeyAuth
// @Router /api/v1/posts/users/{user_id}/{limit}/{offset} [get]
func (h *Handler) AllUserPosts(c *fiber.Ctx) error {
	var (
		posts []models.Post
		count int64
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
	userId, err := utils.StringToInt(c.Params("user_id"))
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

	posts, count, err = h.postRepository.GetUserPosts(userId, offset, limit)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	return c.Status(http.StatusOK).JSON(handlersResponses.NewPostsListResponse(posts, count))
}

// AddLikeToPost func to add a Like a Post
// @Summary Add a new Like to a Post
// @Description AddLikeToPost add a new like to Post
// @ID add-like
// @Tags Post
// @Accept  json
// @Produce  json
// @Param likeRequest body handlersRequests.LikeRequest true "LikeRequest"
// @Param	Authorization	header		string	true	"Authentication header"
// @Success 201 {object} handlersResponses.PostResponse
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		403				{string}	string	"Status Forbidden"
// @Failure		409				{string}	string	"Status Conflict"
// @Failure		422				{string}	string	"Status UnprocessableEntity"
// @Failure		502				{string}	string	"Status BadGateway"
// @Security ApiKeyAuth
// @Router /api/v1/userlikes [post]
func (h *Handler) AddLikeToPost(c *fiber.Ctx) error {

	var payload *handlersRequests.LikeRequest

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.NewError(err))
	}

	userID := int(payload.UserID)
	postID := int(payload.PostID)
	isLikable := payload.IsLikablePost

	user, err := h.userRepository.GetUserById(userID)
	if user == nil && err == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound("User"))
	}
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	post, err := h.postRepository.GetPostById(postID)
	log.Println("Post returned ", post)
	if post == nil && err == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound("Post"))
	}
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	//Check if user already liked this post
	idx := utils.Find(post.Likes, func(value interface{}) bool {
		return value.(models.Like).UserID == uint(userID)
	})
	// User never liked it
	if idx < 0 {
		log.Println(`User nerver liked this post `, userID, postID, isLikable)
		err = h.postRepository.AddLikeToPost(user, post, isLikable)
	} else {
		log.Println(`User Already liked this post `, userID, postID, isLikable)
		err = h.postRepository.RemoveLikeToPost(user, post)

	}
	log.Println(`Number of Like `, len(post.Likes))

	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}
	return c.Status(http.StatusCreated).JSON(handlersResponses.NewPostResponse(post))
}

// AddCommentToPost func to add a new Comment to Post
// @Summary Add a new Comment to a Post
// @Description AddCommentToPost add a new comment to Post
// @ID add-comment
// @Tags Post
// @Accept  json
// @Produce  json
// @Param commentRequest body handlersRequests.CommentRequest true "CommentRequest"
// @Param	Authorization	header		string	true	"Authentication header"
// @Success 201 {object} handlersResponses.PostResponse
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		403				{string}	string	"Status Forbidden"
// @Failure		409				{string}	string	"Status Conflict"
// @Failure		422				{string}	string	"Status UnprocessableEntity"
// @Failure		502				{string}	string	"Status BadGateway"
// @Security ApiKeyAuth
// @Router /api/v1/usercomments [post]
func (h *Handler) AddCommentToPost(c *fiber.Ctx) error {

	var payload *handlersRequests.CommentRequest

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.NewError(err))
	}

	userId := int(payload.UserID)
	commentText := payload.CommentDescription

	user, err := h.userRepository.GetUserById(userId)
	if user == nil && err == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound("User"))
	}

	postId := int(payload.PostID)

	post, err := h.postRepository.GetPostById(postId)
	if post == nil && err == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound("Post"))
	}
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	err = h.postRepository.AddCommentToPost(user, post, commentText)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}
	return c.Status(http.StatusCreated).JSON(handlersResponses.NewPostResponse(post))
}

// UploadImage func to upload an image
// @Summary Upload an image
// @Description UploadImage func to upload an image
// @ID upload-image
// @Tags Post
// @Accept  json
// @Produce  json
// @Param	Authorization	header		string	true	"Authentication header"
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		403				{string}	string	"Status Forbidden"
// @Failure		409				{string}	string	"Status Conflict"
// @Failure		422				{string}	string	"Status UnprocessableEntity"
// @Failure		502				{string}	string	"Status BadGateway"
// @Security ApiKeyAuth
// @Router /api/v1/imageUpload/ [post]
func (h *Handler) UploadImage(c *fiber.Ctx) error {

	// parse incomming image file

	file, err := c.FormFile("picture")
	// log.Println(file.Filename)

	if err != nil {
		log.Println("image upload error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

	// generate new uuid for image name
	uniqueId := uuid.New()

	// remove "- from imageName"

	filename := strings.Replace(uniqueId.String(), "-", "", -1)

	// extract image extension from original file filename

	fileExt := strings.Split(file.Filename, ".")[1]

	// generate image from filename and extension
	image := fmt.Sprintf("%s.%s", filename, fileExt)

	// save image to ./images dir
	err = c.SaveFile(file, fmt.Sprintf("./images/%s", image))

	if err != nil {
		log.Println("image save error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

	// generate image url to serve to client using CDN

	imageUrl := fmt.Sprintf("http://localhost:4000/images/%s", image)

	// create meta data and send to client

	data := map[string]interface{}{

		"imageName": image,
		"imageUrl":  imageUrl,
		"header":    file.Header,
		"size":      file.Size,
	}

	return c.JSON(fiber.Map{"status": 201, "message": "Image uploaded successfully", "data": data})
}
