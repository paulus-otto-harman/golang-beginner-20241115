package service

import "20241115/repository"

type LocationService struct {
	LocationRepo repository.Location
}

func InitLocationService(repo repository.Location) LocationService {
	return LocationService{LocationRepo: repo}
}
