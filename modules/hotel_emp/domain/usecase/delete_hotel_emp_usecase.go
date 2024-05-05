package usecase

import "context"

type DeleteHotelEmpUseCase interface {
	ExecuteDeleteHotelEmp(ctx context.Context, employeeID string) error
}

type deleteHotelEmpUseCaseImpl struct {
	writerRepo hotelEmpWriterRepository
}

var _ DeleteHotelEmpUseCase = (*deleteHotelEmpUseCaseImpl)(nil)

func NewDeleteHotelEmpUseCase(writerRepo hotelEmpWriterRepository) DeleteHotelEmpUseCase {
	return &deleteHotelEmpUseCaseImpl{
		writerRepo: writerRepo,
	}
}

func (useCase deleteHotelEmpUseCaseImpl) ExecuteDeleteHotelEmp(ctx context.Context, employeeID string) error {
	if err := useCase.writerRepo.DeleteHotelEmp(ctx, employeeID); err != nil {
		return err
	}

	return nil
}
