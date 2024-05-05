package repository

import (
	"context"
	"gorm.io/gorm"
	"math"
	"sort"
	"trekkstay/core"
	"trekkstay/modules/hotel/domain/entity"
	database "trekkstay/pkgs/dbs/postgres"
)

type hotelReaderRepositoryImpl struct {
	db database.Database
}

var _ HotelReaderRepository = (*hotelReaderRepositoryImpl)(nil)

func NewHotelReaderRepository(db database.Database) HotelReaderRepository {
	return &hotelReaderRepositoryImpl{
		db: db,
	}
}

func (repo hotelReaderRepositoryImpl) FindHotelByCondition(ctx context.Context, condition map[string]interface{}) (*entity.HotelEntity, error) {
	var hotelEntity entity.HotelEntity

	if err := repo.db.Executor.
		WithContext(ctx).
		Where(condition).
		Preload("Rooms", func(db *gorm.DB) *gorm.DB {
			return db.
				Order("(hotel_rooms.original_price * hotel_rooms.discount_rate / 100) ASC").
				Limit(1)
		}).
		Preload("Province").
		Preload("District").
		Preload("Ward").
		First(&hotelEntity).Error; err != nil {
		return nil, err
	}

	return &hotelEntity, nil
}

func (repo hotelReaderRepositoryImpl) FindHotels(ctx context.Context,
	filter entity.HotelFilterEntity, page, limit int) (*core.Pagination, error) {
	var paging core.Pagination
	var hotels []entity.HotelEntity

	paging.Limit = limit
	paging.Page = page

	var scopeFunctions []func(d *gorm.DB) *gorm.DB

	if filter.Name != nil {
		scopeFunctions = append(scopeFunctions, func(d *gorm.DB) *gorm.DB {
			return d.Where("name LIKE ?", "%"+*filter.Name+"%")
		})
	}

	if filter.WardCode != nil {
		scopeFunctions = append(scopeFunctions, func(d *gorm.DB) *gorm.DB {
			return d.Where("ward_code = ?", *filter.WardCode)
		})
	}

	if filter.DistrictCode != nil {
		scopeFunctions = append(scopeFunctions, func(d *gorm.DB) *gorm.DB {
			return d.Where("district_code = ?", *filter.DistrictCode)
		})
	}

	if filter.ProvinceCode != nil {
		scopeFunctions = append(scopeFunctions, func(d *gorm.DB) *gorm.DB {
			return d.Where("province_code = ?", *filter.ProvinceCode)
		})
	}

	if filter.PriceOrder != nil {
		if *filter.PriceOrder == "asc" {
			paging.Sort = "min_price ASC, hotels.created_at DESC"
		} else {
			paging.Sort = "min_price DESC, hotels.created_at DESC"
		}
	}

	tx := repo.db.Executor.WithContext(ctx).Scopes(scopeFunctions...)
	txTotalRows := tx.Model(&entity.HotelEntity{}).Scopes(scopeFunctions...)
	result := tx.
		Select("hotels.*, MIN(hotel_rooms.original_price * hotel_rooms.discount_rate / 100) as min_price").
		Scopes(core.Paginate(&paging, txTotalRows)).
		Preload("Rooms", func(db *gorm.DB) *gorm.DB {
			return db.
				Order("(hotel_rooms.original_price * hotel_rooms.discount_rate / 100) ASC").
				Limit(1)
		}).
		Preload("Province").
		Preload("District").
		Preload("Ward").
		Joins("left join hotel_rooms on hotel_rooms.hotel_id = hotels.id").
		Group("hotels.id").
		Find(&hotels)

	// Calculate rating
	for i := range hotels {
		rating, avg := int64(0), float64(0)

		repo.db.Executor.
			WithContext(ctx).Table("ratings").Where("hotel_id = ?", hotels[i].ID).
			Count(&rating)
		repo.db.Executor.
			Table("ratings").Where("hotel_id = ?", hotels[i].ID).
			Select("AVG(point)").Row().Scan(&avg)

		hotels[i].Rating = &entity.RatingJSON{
			TotalReview: int(rating),
			AvgPoint:    avg,
		}
	}

	paging.Rows = hotels

	return &paging, result.Error
}

func (repo hotelReaderRepositoryImpl) SearchHotel(ctx context.Context,
	filter entity.HotelSearchEntity, page, limit int) (*core.Pagination, error) {
	var paging core.Pagination
	var hotels []entity.HotelEntity

	if limit == 0 {
		limit = 10
	}

	if page == 0 {
		page = 1
	}

	paging.Limit = limit
	paging.Page = page

	code := ""
	if filter.LocationCode != nil {
		code = *filter.LocationCode
	}

	var result = repo.db.Executor.Raw(`
		SELECT h.* FROM hotels h
		WHERE
		  EXISTS (
			SELECT r.hotel_id FROM hotel_rooms r
			WHERE
			  r.hotel_id = h.id
			  AND NOT EXISTS (
				SELECT sum(res.quantity) FROM reservations res
				WHERE
				  res.room_id = r.id    
				AND (res.check_in_date::DATE, res.check_out_date::DATE) OVERLAPS (?::DATE, ?::DATE)
				GROUP BY res.room_id
				HAVING sum(res.quantity) > r.quantity - ?
			  )
		  )
		AND (h.province_code = ? OR h.district_code = ? OR h.ward_code = ?)
		LIMIT ? OFFSET ?
	`,
		*filter.CheckInDate,
		*filter.CheckOutDate,
		*filter.NumOfRooms,
		code, code, code,
		limit, (page-1)*limit,
	).Scan(&hotels)

	for i := range hotels {
		var province entity.ProvinceEntity
		var district entity.DistrictEntity
		var ward entity.WardEntity

		repo.db.Executor.WithContext(ctx).Model(&entity.ProvinceEntity{}).Where("code = ?", hotels[i].ProvinceCode).First(&province)
		repo.db.Executor.WithContext(ctx).Model(&entity.DistrictEntity{}).Where("code = ?", hotels[i].DistrictCode).First(&district)
		repo.db.Executor.WithContext(ctx).Model(&entity.WardEntity{}).Where("code = ?", hotels[i].WardCode).First(&ward)

		hotels[i].Province = province
		hotels[i].District = district
		hotels[i].Ward = ward

		var rooms []entity.HotelRoomEntity
		repo.db.Executor.WithContext(ctx).Model(&entity.HotelRoomEntity{}).
			Where("hotel_id = ?", hotels[i].ID).
			Limit(1).
			Order("(hotel_rooms.original_price * hotel_rooms.discount_rate / 100) ASC").
			Find(&rooms)

		hotels[i].Rooms = rooms
	}

	// Sort price
	if filter.PriceOrder != nil {
		if *filter.PriceOrder == "asc" {
			paging.Sort = "min_price ASC, hotels.created_at DESC"

			sort.Slice(hotels, func(i, j int) bool {
				// Get the price of the first room in each hotel for comparison
				priceI := hotels[i].Rooms[0].OriginalPrice * hotels[i].Rooms[0].DiscountRate / 100
				priceJ := hotels[j].Rooms[0].OriginalPrice * hotels[j].Rooms[0].DiscountRate / 100
				// Compare the prices to determine the order
				return priceI < priceJ
			})
		} else {
			paging.Sort = "min_price DESC, hotels.created_at DESC"

			sort.Slice(hotels, func(i, j int) bool {
				// Get the price of the first room in each hotel for comparison
				priceI := hotels[i].Rooms[0].OriginalPrice * hotels[i].Rooms[0].DiscountRate / 100
				priceJ := hotels[j].Rooms[0].OriginalPrice * hotels[j].Rooms[0].DiscountRate / 100
				// Compare the prices to determine the order
				return priceI > priceJ
			})
		}
	}

	// Calculate rating
	for i := range hotels {
		rating, avg := int64(0), float64(0)

		repo.db.Executor.
			WithContext(ctx).Table("ratings").Where("hotel_id = ?", hotels[i].ID).
			Count(&rating)
		repo.db.Executor.
			Table("ratings").Where("hotel_id = ?", hotels[i].ID).
			Select("AVG(point)").Row().Scan(&avg)

		hotels[i].Rating = &entity.RatingJSON{
			TotalReview: int(rating),
			AvgPoint:    avg,
		}
	}

	paging.TotalRows = int64(len(hotels))
	paging.TotalPages = int(math.Ceil(float64(paging.TotalRows) / float64(limit)))
	paging.Rows = hotels

	return &paging, result.Error
}
