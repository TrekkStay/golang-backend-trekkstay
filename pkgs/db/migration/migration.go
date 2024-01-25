package migration

import (
	"gorm.io/gorm"
	hotel "trekkstay/modules/hotel/domain/entity"
	user "trekkstay/modules/user/domain/entity"
)

// Migration is a function that performs migrations on the given database.
//
// It takes a pointer to a gorm.DB object as a parameter.
// It returns an error indicating the success or failure of the migration operation.
func Migration(db *gorm.DB) error {
	err := db.AutoMigrate(
		user.UserEntity{},
		hotel.HotelEntity{},
		hotel.RoomEntity{},
		hotel.HotelEmployeeEntity{},
		hotel.HotelFacilityEntity{},
		hotel.RoomFacilityEntity{},
	)

	return err
}
