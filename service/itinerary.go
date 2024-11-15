package service

import "20241115/repository"

type ItineraryService struct {
	ItineraryRepo repository.Itinerary
}

func InitItineraryService(repo repository.Itinerary) *ItineraryService {
	return &ItineraryService{ItineraryRepo: repo}
}

func (service *ItineraryService) Get() {

}
