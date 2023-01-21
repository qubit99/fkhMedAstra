package repository

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"medastra/models"
)

const UserTable = "user"
const LoginTable = "login"
const SlotTable = "slots"
const DoctorTable = "doctor"
const BookingTable = "booking"

type RepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepository(Db *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{
		Db: Db,
	}
}

func createSearchQueryFromRequest(db *gorm.DB, searchreq *models.DoctorSearchRequest) *gorm.DB {
	if len(searchreq.Specialities) > 0 {
		db = db.Where("speciality in (?)", searchreq.Specialities)
	}
	if len(searchreq.Cities) > 0 {
		db = db.Where("city in (?)", searchreq.Cities)
	}

	if searchreq.SortBy == "" {
		searchreq.SortBy = "id"
	}
	if searchreq.SortOrder == "" {
		searchreq.SortOrder = "desc"
	}
	db = db.Order(searchreq.SortBy + " " + searchreq.SortOrder)
	if searchreq.Limit != 0 {
		db = db.Limit(searchreq.Limit).Offset(searchreq.Offset)
	}
	return db
}

func (u RepositoryImpl) SaveProfile(user *models.UserProfile) error {
	err := u.Db.Table(UserTable).Create(&user).Error
	return err
}
func (u RepositoryImpl) UpdateProfile(user *models.UserProfile) error {
	err := u.Db.Table(UserTable).Model(&user).Updates(&user).Error
	return err
}

func (u RepositoryImpl) DeleteProfile(Id int) error {
	var user models.UserProfile
	err := u.Db.Table(UserTable).Where("id=?", Id).Delete(&user).Error
	return err
}
func (u RepositoryImpl) FindProfileByUserName(username string) (*models.UserProfile, error) {
	var result models.UserProfile
	err := u.Db.Table(UserTable).
		Where("username = ?", username).
		First(&result).
		Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}
func (u RepositoryImpl) CreateLogin(userLogin *models.User) error {
	return u.Db.Table(LoginTable).Create(userLogin).Error
}
func (u RepositoryImpl) CheckLogin(username string, password string) error {
	var userLogin models.User
	err := u.Db.Table(LoginTable).Where("username=?", username).First(&userLogin).Error
	if err != nil {
		return err
	}
	if userLogin.Password != password {
		return fmt.Errorf("Incorrect password or incorrect username")
	}
	return nil
}

func (u RepositoryImpl) CreateBooking(testBooking *models.Booking) error {
	err := u.Db.Table(BookingTable).Create(testBooking).Error
	return err
}

func (u RepositoryImpl) GetBookings(username string) ([]models.Booking, error) {
	var bookings []models.Booking
	result := u.Db.Table(BookingTable).Where("username=?", username).Find(&bookings)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return bookings, nil
}

func (u RepositoryImpl) CreateDoctor(doctor *models.Doctor) error {
	return u.Db.Table(DoctorTable).Create(doctor).Error
}
func (u RepositoryImpl) SearchDoctors(searchreq *models.DoctorSearchRequest) ([]*models.Doctor, error) {
	query := createSearchQueryFromRequest(u.Db.Table(DoctorTable), searchreq)
	result := make([]*models.Doctor, 0)
	err := query.Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, err
}

func (u RepositoryImpl) GetSlotsByDoctor(doctorId int) ([]*models.Slot, error) {
	result := make([]*models.Slot, 0)
	err := u.Db.Table(SlotTable).Where("doctor_id=?", doctorId).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u RepositoryImpl) BookSlot(username string, slotId int) error {
	//ideally this all should happend in a single transaction

	//check capacity
	var slot models.Slot

	err := u.Db.Table(SlotTable).Where("slot_id=?", slotId).First(&slot).Error
	if err != nil {
		return err
	}

	if slot.Capacity < 0 {
		return fmt.Errorf("Slot Full")
	}

	slot.Capacity--
	err = u.Db.Table(SlotTable).Model(&slot).Updates(&slot).Error
	if err != nil {
		return err
	}

	booking := models.Booking{slotId, username, slot.DoctorId}
	return u.Db.Table(BookingTable).Create(&booking).Error
}
func (u RepositoryImpl) getSlotById(id int) (*models.Slot, error) {
	var slot models.Slot
	err := u.Db.Table(SlotTable).Where("slot_id=?", id).First(&slot).Error
	if err != nil {
		return nil, err
	}
	return &slot, nil
}
func (u RepositoryImpl) getDoctorById(id int) (*models.Doctor, error) {
	var doctor models.Doctor
	err := u.Db.Table(DoctorTable).Where("id=?", id).First(&doctor).Error
	if err != nil {
		return nil, err
	}
	return &doctor, nil
}
func (u RepositoryImpl) GetUserBookings(username string) ([]*models.UserBookingResponse, error) {
	bookings := make([]models.Booking, 0)
	err := u.Db.Table(BookingTable).Where("username=?", username).Find(&bookings).Error
	if err != nil {
		return nil, err
	}
	bookingresp := make([]*models.UserBookingResponse, 0)
	for _, booking := range bookings {
		slot, err := u.getSlotById(booking.SlotId)
		if err != nil {
			log.Println(err)
			continue
		}
		doctor, err := u.getDoctorById(booking.DoctorId)
		if err != nil {
			log.Println(err)
			continue
		}
		bookingresp = append(bookingresp, &models.UserBookingResponse{slot, doctor})
	}
	return bookingresp, nil
}
func (u RepositoryImpl) GetDoctorBookings(doctorId int) ([]*models.DoctorBookingResponse, error) {
	bookings := make([]models.Booking, 0)
	err := u.Db.Table(BookingTable).Where("doctor_id=?", doctorId).Find(&bookings).Error
	if err != nil {
		return nil, err
	}
	bookingresp := make([]*models.DoctorBookingResponse, 0)
	for _, booking := range bookings {
		slot, err := u.getSlotById(booking.SlotId)
		if err != nil {
			log.Println(err)
			continue
		}
		user, err := u.FindProfileByUserName(booking.Username)
		if err != nil {
			log.Println(err)
			continue
		}
		bookingresp = append(bookingresp, &models.DoctorBookingResponse{slot, user})
	}
	return bookingresp, nil
}
