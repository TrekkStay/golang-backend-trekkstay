package usecase

import (
	"context"
	"trekkstay/modules/hotel_emp/constant"
	"trekkstay/modules/hotel_emp/domain/entity"
)

type FilterHotelEmpUseCase interface {
	ExecuteFilterHotelEmp(ctx context.Context, hotelID string) ([]entity.HotelEmpEntity, error)
}

type filterHotelEmpUseCaseImpl struct {
	readerRepo hotelEmpReaderRepository
}

var _ FilterHotelEmpUseCase = (*filterHotelEmpUseCaseImpl)(nil)

func NewFilterHotelEmpUseCase(readerRepo hotelEmpReaderRepository) FilterHotelEmpUseCase {
	return &filterHotelEmpUseCaseImpl{
		readerRepo: readerRepo,
	}
}

func (useCase filterHotelEmpUseCaseImpl) ExecuteFilterHotelEmp(ctx context.Context, hotelID string) ([]entity.HotelEmpEntity, error) {
	hotelEmployees, err := useCase.readerRepo.FindHotelEmpByHotelID(ctx, hotelID)
	if err != nil {
		return nil, constant.ErrSomethingWentWrong(err)
	}

	return hotelEmployees, nil
}
