package service

import (
	"example.com/gallery/dto"
	"example.com/gallery/entity"
	"example.com/gallery/repository"
	"fmt"
	"github.com/mashingan/smapping"
	"log"
)

// PostService is a contract about something that this service can do
type PostService interface {
	Insert(post dto.PostCreateDTO) entity.Post
	Update(post dto.PostUpdateDTO) entity.Post
	Delete(post entity.Post)
	All() []entity.Post
	FindByID(ID uint64) entity.Post
	FindByTopicID(topicID uint64) []entity.Post
	GetTrendingPosts() []entity.Post
	IsAllowedToEdit(userID string, postID uint64) bool
}

// postService is a concrete implementation of PostService interface
type postService struct {
	postRepository repository.PostRepository
}

// NewPostService creates a new instance of PostService
func NewPostService(postRepo repository.PostRepository) PostService {
	return &postService{
		postRepository: postRepo,
	}
}

// Insert function creates a new post
func (service *postService) Insert(post dto.PostCreateDTO) entity.Post {
	postToInsert := entity.Post{}
	err := smapping.FillStruct(&postToInsert, smapping.MapFields(&post))
	if err != nil {
		log.Fatalf("Failed to map %v", err)
	}
	insertedPost := service.postRepository.InsertPost(postToInsert)
	return insertedPost
}

// Update function updates an existing post
func (service *postService) Update(post dto.PostUpdateDTO) entity.Post {
	postToUpdate := entity.Post{}
	err := smapping.FillStruct(&postToUpdate, smapping.MapFields(&post))
	if err != nil {
		log.Fatalf("Failed to map %v", err)
	}
	updatedPost := service.postRepository.UpdatePost(postToUpdate)
	return updatedPost
}

func (service *postService) Delete(post entity.Post) {
	service.postRepository.DeletePost(post)
}

func (service *postService) All() []entity.Post {
	return service.postRepository.AllPost()
}

func (service *postService) FindByID(ID uint64) entity.Post {
	return service.postRepository.FindPostByID(ID)
}

func (service *postService) FindByTopicID(topicID uint64) []entity.Post {
	return service.postRepository.FindPostByTopicID(topicID)
}

func (service *postService) GetTrendingPosts() []entity.Post {
	return service.postRepository.GetTrendingPosts()
}

func (service *postService) IsAllowedToEdit(userID string, postID uint64) bool {
	post := service.postRepository.FindPostByID(postID)
	id := fmt.Sprintf("%v", post.UserID)
	return userID == id
}
