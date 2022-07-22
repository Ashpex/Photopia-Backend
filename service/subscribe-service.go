package service

import (
	"github.com/mashingan/smapping"
	"gitlab.zalopay.vn/top/intern/vybnt/gallery-backend/gallery/dto"
	"gitlab.zalopay.vn/top/intern/vybnt/gallery-backend/gallery/entity"
	"gitlab.zalopay.vn/top/intern/vybnt/gallery-backend/gallery/repository"
	"log"
)

type SubscribeService interface {
	Subscribe(subscribeDTO dto.SubscribeDTO) entity.Subscribe
	Unsubscribe(unsubscribeDTO entity.Subscribe)
	AllSubscribes(topicID uint64) []entity.Subscribe
	CountSubscribes(topicID uint64) int
	AllSubscribesByUser(userID uint64) []entity.Subscribe
}

type subscribeService struct {
	subscribeRepository repository.SubscribeRepository
}

func NewSubscribeService(subscribeRepository repository.SubscribeRepository) SubscribeService {
	return &subscribeService{
		subscribeRepository: subscribeRepository,
	}
}

func (service *subscribeService) Subscribe(subscribeDTO dto.SubscribeDTO) entity.Subscribe {
	subscribeToSave := entity.Subscribe{}
	err := smapping.FillStruct(&subscribeToSave, smapping.MapFields(&subscribeDTO))
	if err != nil {
		log.Fatalln("Failed map", err)
	}
	savedSubscribe := service.subscribeRepository.Subscribe(subscribeToSave)
	return savedSubscribe
}

func (service *subscribeService) Unsubscribe(subscribe entity.Subscribe) {
	unsubscribeToDelete := entity.Subscribe{}
	err := smapping.FillStruct(&unsubscribeToDelete, smapping.MapFields(&subscribe))
	if err != nil {
		log.Fatalln("Failed map", err)
	}
	service.subscribeRepository.Unsubscribe(unsubscribeToDelete)
}

func (service *subscribeService) AllSubscribes(topicID uint64) []entity.Subscribe {
	return service.subscribeRepository.AllSubscribes(topicID)
}

func (service *subscribeService) CountSubscribes(topicID uint64) int {
	return service.subscribeRepository.CountSubscribes(topicID)
}

func (service *subscribeService) AllSubscribesByUser(userID uint64) []entity.Subscribe {
	return service.subscribeRepository.AllSubscribesByUser(userID)
}
