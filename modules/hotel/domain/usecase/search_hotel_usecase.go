package usecase

import (
	"context"
	"math"
	"sort"
	"trekkstay/core"
	"trekkstay/modules/hotel/domain/entity"
	"trekkstay/utils"
)

type SearchHotelUseCase interface {
	ExecuteSearchHotel(ctx context.Context, filter entity.HotelSearchEntity, page, limit int) (*core.Pagination, error)
}

type searchHotelUseCaseImpl struct {
	readerRepo hotelReaderRepository
}

var _ SearchHotelUseCase = (*searchHotelUseCaseImpl)(nil)

func NewSearchHotelUseCase(readerRepo hotelReaderRepository) SearchHotelUseCase {
	return &searchHotelUseCaseImpl{
		readerRepo: readerRepo,
	}
}

func (useCase searchHotelUseCaseImpl) ExecuteSearchHotel(ctx context.Context, filter entity.HotelSearchEntity, page, limit int) (*core.Pagination, error) {
	hotels, err := useCase.readerRepo.SearchHotel(ctx, filter, page, limit)

	if err != nil {
		return nil, err
	}

	// Sort hotel
	if filter.AttractionLat != nil && filter.AttractionLng != nil {
		hotelList := hotels.Rows.([]entity.HotelEntity)

		sort.Slice(hotelList, func(i, j int) bool {
			// Sort by distance
			return utils.HaversineDistance(
				*filter.AttractionLat,
				*filter.AttractionLng,
				hotelList[i].Coordinates.Lat,
				hotelList[i].Coordinates.Lng,
			) < utils.HaversineDistance(
				*filter.AttractionLat,
				*filter.AttractionLng,
				hotelList[j].Coordinates.Lat,
				hotelList[j].Coordinates.Lng,
			)
		})

		for i := range hotelList {
			hotelList[i].Attraction = new(entity.AttractionsJSON)
			hotelList[i].Attraction.Distance = math.Round(utils.HaversineDistance(
				*filter.AttractionLat,
				*filter.AttractionLng,
				hotelList[i].Coordinates.Lat,
				hotelList[i].Coordinates.Lng,
			)*1000) / 1000

			hotelList[i].Attraction.Name = *filter.AttractionName
			if hotelList[i].Attraction.Distance == 0 {
				hotelList[i].Attraction.Distance = 0.001
			}
		}

		hotels.Rows = hotelList
	}

	if len(hotels.Rows.([]entity.HotelEntity)) == 0 {
		hotels.Rows = []entity.HotelEntity{}
	}

	return hotels, nil
}
