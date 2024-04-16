package usecase

import (
	"context"
	"trekkstay/core"
	"trekkstay/modules/hotel/domain/entity"
	"trekkstay/modules/hotel_emp/constant"
)

type UpdateHotelUseCase interface {
	ExecuteUpdateHotel(ctx context.Context, hotel entity.HotelEntity) error
}

type updateHotelUseCaseImpl struct {
	hotelReaderEmpRepo hotelEmpReaderRepository
	hotelWriterRepo    hotelWriterRepository
}

var _ UpdateHotelUseCase = (*updateHotelUseCaseImpl)(nil)

func NewUpdateHotelUseCase(hotelReaderEmpRepo hotelEmpReaderRepository, hotelWriterRepo hotelWriterRepository) UpdateHotelUseCase {
	return &updateHotelUseCaseImpl{
		hotelReaderEmpRepo: hotelReaderEmpRepo,
		hotelWriterRepo:    hotelWriterRepo,
	}
}

func (useCase updateHotelUseCaseImpl) ExecuteUpdateHotel(ctx context.Context, hotel entity.HotelEntity) error {
	requester := ctx.Value(core.CurrentRequesterKeyStruct{}).(core.Requester)

	owner, _ := useCase.hotelReaderEmpRepo.FindHotelEmpByCondition(ctx, map[string]interface{}{
		"id": requester.GetUserID(),
	})

	if hotel.ID != owner.HotelID {
		return constant.ErrPermissionDenied(nil)
	}

	if err := useCase.hotelWriterRepo.UpdateHotel(ctx, hotel); err != nil {
		return constant.ErrSomethingWentWrong(err)
	}

	return nil
}
