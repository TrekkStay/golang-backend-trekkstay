package usecase

import (
	"context"
	"trekkstay/core"
	"trekkstay/modules/reservation/domain/entity"
	"trekkstay/pkgs/s3"
	"trekkstay/utils"
)

type CreateReservationUseCase interface {
	ExecuteCreateReservation(ctx context.Context, reservation *entity.ReservationEntity) (*entity.ReservationEntity, error)
}

type createReservationUseCaseImpl struct {
	roomReaderRepo        HotelRoomReaderRepository
	hotelReaderRepo       HotelReaderRepository
	reservationWriterRepo ReservationWriterRepository
	uploadHandler         s3.UploadHandler
}

var _ CreateReservationUseCase = (*createReservationUseCaseImpl)(nil)

func NewCreateReservationUseCase(
	roomReaderRepo HotelRoomReaderRepository,
	hotelReaderRepo HotelReaderRepository,
	reservationWriterRepo ReservationWriterRepository,
	uploadHandler s3.UploadHandler,
) CreateReservationUseCase {
	return &createReservationUseCaseImpl{
		roomReaderRepo:        roomReaderRepo,
		hotelReaderRepo:       hotelReaderRepo,
		reservationWriterRepo: reservationWriterRepo,
		uploadHandler:         uploadHandler,
	}
}

func (useCase createReservationUseCaseImpl) ExecuteCreateReservation(ctx context.Context,
	reservation *entity.ReservationEntity) (*entity.ReservationEntity, error) {
	// Find room
	room, err := useCase.roomReaderRepo.FindHotelRoomByCondition(ctx,
		map[string]interface{}{"id": reservation.RoomID})

	if err != nil {
		return nil, err
	}

	// Find hotel
	hotel, err := useCase.hotelReaderRepo.FindHotelByCondition(ctx,
		map[string]interface{}{"id": room.HotelID})

	// Retrieve room information
	(*reservation).Room.HotelID = (*room).HotelID
	(*reservation).Room.HotelName = (*hotel).Name
	(*reservation).Room.Location = (*hotel).District.NameEn
	(*reservation).Room.Type = (*room).Type
	(*reservation).Room.OriginalPrice = (*room).OriginalPrice
	(*reservation).Room.Images = entity.MediaJSON((*room).Images)
	(*reservation).Room.BookingPrice = (*room).OriginalPrice - ((*room).OriginalPrice * (*room).DiscountRate / 100)

	(*reservation).Status = "UPCOMING"
	(*reservation).UserID = ctx.Value(core.CurrentRequesterKeyStruct{}).(core.Requester).GetUserID()
	(*reservation).TotalPrice = (*reservation).Room.BookingPrice * (*reservation).Quantity

	if err := useCase.reservationWriterRepo.InsertReservation(ctx, reservation); err != nil {
		return nil, err
	}

	qrCode, _ := utils.GenerateQRCode(reservation.ID)
	url, err := useCase.uploadHandler.UploadImageToS3(qrCode)
	(*reservation).QRCodeURL = *url

	if err := useCase.reservationWriterRepo.UpdateReservation(ctx, *reservation); err != nil {
		return nil, err
	}

	return reservation, nil
}
