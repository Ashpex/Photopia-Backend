package service

import (
	"example.com/gallery/dto"
	"example.com/gallery/entity"
	"example.com/gallery/repository"
	"github.com/mashingan/smapping"
	"log"
)

type FollowService interface {
	Follow(follower dto.FollowDTO) entity.Follower
	UnFollow(follower entity.Follower)
	AllFollowers(userID uint64) []entity.Follower
}

type followerService struct {
	followerRepository repository.FollowerRepository
}

func NewFollowService(followerRepo repository.FollowerRepository) FollowService {
	return &followerService{
		followerRepository: followerRepo,
	}
}

// Follow is a function that will follow a user
func (service *followerService) Follow(follower dto.FollowDTO) entity.Follower {
	followerToFollow := entity.Follower{}
	err := smapping.FillStruct(&followerToFollow, smapping.MapFields(&follower))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	followedFollower := service.followerRepository.Follow(followerToFollow)
	return followedFollower
}

// UnFollow is a function that will unfollow a user
func (service *followerService) UnFollow(follower entity.Follower) {
	service.followerRepository.Unfollow(follower)
}

func (service *followerService) AllFollowers(userID uint64) []entity.Follower {
	return service.followerRepository.AllFollower(userID)
}
