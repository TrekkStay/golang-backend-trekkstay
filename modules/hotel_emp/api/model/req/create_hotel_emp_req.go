package req

type CreateHotelEmpReq struct {
	FullName   string `json:"full_name" validate:"required" extensions:"x-order=1"`
	Email      string `json:"email" validate:"required,email" extensions:"x-order=2"`
	Phone      string `json:"phone" validate:"required,phone" extensions:"x-order=3"`
	Contract   string `json:"contract" validate:"required,oneof=FULL_TIME PART_TIME INTERNSHIP" extensions:"x-order=4" example:"FULL_TIME | PART_TIME | INTERNSHIP"`
	BaseSalary int    `json:"base_salary" validate:"required" extensions:"x-order=5"`
}
