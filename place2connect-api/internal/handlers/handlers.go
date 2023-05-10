package handlers

import (
	postRepository "github.com/fxfrancky/place2connect-api/internal/repository/post"
	userRepository "github.com/fxfrancky/place2connect-api/internal/repository/user"
)

type Handler struct {
	userRepository userRepository.UserRepository
	postRepository postRepository.PostRepository
}

func NewHandler(userRepo userRepository.UserRepository, postRepo postRepository.PostRepository) *Handler {
	return &Handler{
		userRepository: userRepo,
		postRepository: postRepo,
	}
}
