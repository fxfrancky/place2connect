package handlers

import (
	"context"
	"fmt"

	"github.com/fxfrancky/place2connect-api/config"
	"github.com/fxfrancky/place2connect-api/internal/initializers"
	"github.com/fxfrancky/place2connect-api/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/redis/go-redis/v9"
)

func (h *Handler) NewRoutes(config *config.Config, swagg swagger.Config) *fiber.App {

	apiPath := "/api/" + config.APIVersion
	app := fiber.New()

	app = middleware.EnableCors(app)

	// Main /api/v1 route Group
	v1 := app.Group(apiPath)
	ctx := context.TODO()
	value, err := initializers.RedisClient.Get(ctx, "test").Result()

	if err == redis.Nil {
		fmt.Println("key: test does not exist")
	} else if err != nil {
		panic(err)
	}
	// Check that the api is healthy /api/v1/healthchecker
	v1.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": value,
		})
	})

	// Add static images
	v1.Static("/images", "./images")

	// add docs /api/v1/swagger
	v1.Get("/swagger/*", swagger.HandlerDefault)
	v1.Get("/swagger/*", swagger.New(swagg))

	// Authentification and user management
	grpUsers := v1.Group("/auth")                  // /api/v1/auth
	grpUsers.Post("/register", h.SignUpUser)       // /api/v1/auth/register
	grpUsers.Post("/login", h.SignInUser)          // /api/v1/auth/login
	grpUsers.Get("/refresh", h.RefreshAccessToken) // /api/v1/auth/refresh
	// Protected users routes
	grpUsers.Get("/logout", middleware.DeserializeUser, h.LogoutUser) // /api/v1/auth/logout

	// user management secure
	usersGrp := v1.Group("/users")                                              // /api/v1/users
	usersGrp.Get("/me", middleware.DeserializeUser, h.GetMe)                    // /api/v1/users/me
	usersGrp.Put("/:id", middleware.DeserializeUser, h.UpdateUser)              // /api/v1/users/:id [put]
	usersGrp.Delete("/:id", middleware.DeserializeUser, h.DeleteUser)           // /api/v1/users/:id [delete]
	usersGrp.Get("/email/:email", middleware.DeserializeUser, h.GetUserByEmail) // /api/v1/users/:email [get]
	usersGrp.Get("/:id", middleware.DeserializeUser, h.GetUserById)             // /api/v1/users/:email [get]
	usersGrp.Get("/:limit?/:offset?", middleware.DeserializeUser, h.AllUsers)   // /api/v1/users/:limit?/:offset? [get]

	// user management secure
	userFriendGrp := v1.Group("/userfriends/")
	userFriendGrp.Get("/:userID/:limit?/:offset?", middleware.DeserializeUser, h.AllUserFriends) // /api/v1/userfriends/:user_id/:limit?/:offset? [get]
	userFriendGrp.Post("", middleware.DeserializeUser, h.AddFriendToUser)                        // /api/v1/userfriends/ [post]
	userFriendGrp.Delete("", middleware.DeserializeUser, h.RemoveFriendToUser)                   // /api/v1/userfriends/ [delete]

	// post management secure
	postsGrp := v1.Group("/posts")
	postsGrp.Post("", middleware.DeserializeUser, h.CreatePost) // /api/v1/posts [post]
	postsGrp.Get("/:limit?/:offset?", middleware.DeserializeUser, h.FeedPosts)

	userPostsGrp := v1.Group("/userposts")
	userPostsGrp.Get("/:user_id/:limit?/:offset?", middleware.DeserializeUser, h.AllUserPosts) // /api/v1/userposts/:user_id/:limit?/:offset? [get]

	userLikesGrp := v1.Group("/userlikes")
	userLikesGrp.Post("", middleware.DeserializeUser, h.AddLikeToPost) // /api/v1/userlikes/ [post]

	userCommentsGrp := v1.Group("/usercomments")
	userCommentsGrp.Post("", middleware.DeserializeUser, h.AddCommentToPost) // /api/v1/usercomments/ [post]

	imageUploadGrp := v1.Group("/imageUpload")
	imageUploadGrp.Post("", h.UploadImage) // /api/v1/imageUpload/ [post]

	// Close Other Routes
	v1.All("*", func(c *fiber.Ctx) error {
		path := c.Path()
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "fail",
			"message": fmt.Sprintf("Path: %v does not exists on this server", path),
		})
	})

	return app

}
