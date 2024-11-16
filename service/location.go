package service

import (
	"20241115/model"
	"20241115/repository"
)

type LocationService struct {
	LocationRepo repository.Location
}

func InitLocationService(repo repository.Location) *LocationService {
	return &LocationService{LocationRepo: repo}
}

func (service LocationService) Get(eventId int) (*model.Location, error) {
	return service.LocationRepo.Get(eventId)
}
