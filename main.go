package main

import (
	"fmt"
	"time"
)

type RoomType string
type RoomStatus string
type ReservationStatus string

const (
	Single       RoomType = "single"
	Double       RoomType = "double"
	Suite        RoomType = "suite"
	Deluxe       RoomType = "deluxe"
	Presidential RoomType = "presidential"
)
const (
	Available        RoomStatus = "available"
	Occupied         RoomStatus = "occupied"
	UnderMaintenance RoomStatus = "underMaintenance"
	Reserved         RoomStatus = "reserved"
)
const (
	Pending    ReservationStatus = "pending"
	Confirmed  ReservationStatus = "confirmed"
	CheckedIn  ReservationStatus = "checkedIn"
	CheckedOut ReservationStatus = "checkedOut"
	Cancelled  ReservationStatus = "cancelled"
)

type Chargeable interface {
	GetPrice() float64
	GetDescription() string
}
type DiscountPolicy interface {
	CalculateDiscount(reservation Reservation) float64
}
type LoyaltyDiscountPolicy struct{}

func (p LoyaltyDiscountPolicy) CalculateDiscount(r Reservation) float64 {
	points := r.Guest.LoyaltyPoints

	switch {
	case points >= 200:
		return 0.15
	case points >= 100:
		return 0.10
	case points >= 50:
		return 0.05
	default:
		return 0
	}
}

// لو عايز تغيّر “قرار” → Strategy / Policy
// لو عايز تضيف “سلوك” → Wrapper / Delegate

// 1️⃣ Policy / Strategy — إمتى؟
// تستخدمها لما:
// السلوك قرار
// القرار ممكن يتغيّر
// نفس الكائن يخضع لقرارات مختلفة
// مفيش state بيتغيّر في الكائن نفسه
// كلمات مفتاحية:
// calculate
// decide
// choose
// evaluate
// apply rule

// 2️⃣ Wrapper / Delegation — إمتى؟
// تستخدمه لما:
// عايز تضيف سلوك
// السلوك يتنفّذ قبل / بعد / حول السلوك الأصلي
// السلوك يؤثّر على lifecycle
// السلوك بيتراكم (stackable)
// كلمات مفتاحية:
// before / after
// retry
// schedule
// log
// cache
// authorize

type Room struct {
	RoomNumber    string
	Type          RoomType
	Status        RoomStatus
	Floor         int
	PricePerNight float64
	MaxOccupancy  int
	Amenities     []string
}

func (r Room) GetDescription() string {
	return fmt.Sprintf(
		"Room %s (%s) - Floor %d - $%.2f/night",
		r.RoomNumber,
		r.Type,
		r.Floor,
		r.PricePerNight,
	)
}
func (r Room) GetPrice() float64 {
	return r.PricePerNight
}
func (r Room) IsAvailable() bool {
	return r.Status == Available
}

// we need to study stratgey pattern
// بدل ما تسأل:
// هل Room تحسب السعر إزاي؟
// اسأل:
// مين مسؤول عن سياسة التسعير؟
// الإجابة:
// مش Room
// Room:
// تعرف سعرها الأساسي
// تعرف نوعها
// لكن لا تعرف سياسة الفندق
func (r *Room) ChangeStatus(newStatus RoomStatus) { // pointer because we need to change the state
	r.Status = newStatus
}

type Guest struct {
	GuestId       string
	Name          string
	Email         string
	Phone         string
	IdNumber      string
	LoyaltyPoints int
}

func (g Guest) GetGuestInfo() string {
	return fmt.Sprintf(
		"%s (%s) - Points: %d",
		g.Name,
		g.Email,
		g.LoyaltyPoints,
	)
}
func (g *Guest) AddLoyaltyPoints(points int) {
	g.LoyaltyPoints += points
}

// نفصل بين:
// Fact (حقيقة)
// Decision (قرار)
// Fact
// LoyaltyPoints
// Guest type
// Stay length
// Decision
// Discount rate
// Promotions
// Offers
// 📌 الـ Fact في الـ Entity
// 📌 الـ Decision في Policy / Strategy

// إمتى أعرف إن الحاجة دي Policy؟
// اسأل نفسك 3 أسئلة 👇
// 1️⃣ هل القاعدة ممكن تتغير؟
// لو أيوة → Policy
// 2️⃣ هل القاعدة تخص البزنس مش الكائن؟
// لو أيوة → Policy
// 3️⃣ هل نفس الكائن ممكن يخضع لقرارات مختلفة؟
// لو أيوة → Strategy

type Service struct {
	ServiceID   string
	Name        string
	Price       float64
	Description string
}

func (s Service) GetPrice() float64 {
	return s.Price
}

func (s Service) GetDescription() string {
	return s.Name + " - " + s.Description
}

type Reservation struct {
	ReservationID string
	Guest         Guest
	Room          Room
	CheckInDate   time.Time
	CheckOutDate  time.Time
	Status        ReservationStatus
	Services      []Service
	TotalGuests   int
}

func (r Reservation) GetNumberOfNights() int {
	return int(r.CheckOutDate.Sub(r.CheckInDate).Hours() / 24)
}
func (r Reservation) GetRoomCost() float64 {
	return float64(r.GetNumberOfNights()) * r.Room.GetPrice()
}
func (r Reservation) GetServicesCost() float64 {
	var total float64
	for _, s := range r.Services {
		total += s.GetPrice()
	}
	return total
}

func (r Reservation) GetTotal(policy DiscountPolicy) float64 {
	baseTotal := r.GetRoomCost() + r.GetServicesCost()
	discountRate := policy.CalculateDiscount(r)
	return baseTotal * (1 - discountRate)
}

func (r *Reservation) AddService(service Service) {
	r.Services = append(r.Services, service)
}
func (r *Reservation) CheckIn() {
	r.Status = CheckedIn
	r.Room.ChangeStatus(Occupied)
}

func (r *Reservation) CheckOut() {
	r.Status = CheckedOut
	r.Room.ChangeStatus(Available)
}

func (r *Reservation) Cancel() {
	r.Status = Cancelled
}

type Hotel struct {
	Name         string
	Address      string
	Rooms        []Room
	Guests       []Guest
	Reservations []Reservation
	Services     []Service
}

func (h *Hotel) AddRoom(room Room) {
	h.Rooms = append(h.Rooms, room)
}
func (h *Hotel) RegisterGuest(guest Guest) {
	h.Guests = append(h.Guests, guest)
}
func (h *Hotel) AddService(service Service) {
	h.Services = append(h.Services, service)
}

func (h *Hotel) CreateReservation(guest Guest, room *Room, checkIn, checkOut time.Time, totalGuests int) (*Reservation, error) {
	if !checkOut.After(checkIn) {
		return nil, fmt.Errorf("check-out must be after check-in")
	}

	if !room.IsAvailable() {
		return nil, fmt.Errorf("room %s is not available", room.RoomNumber)
	}
	if totalGuests > room.MaxOccupancy {
		return nil, fmt.Errorf("room %s exceeds max occupancy", room.RoomNumber)
	}

	room.ChangeStatus(Reserved)
	reservation := &Reservation{
		ReservationID: fmt.Sprintf("R-%d", len(h.Reservations)+1),
		Guest:         guest,
		Room:          *room,
		CheckInDate:   checkIn,
		CheckOutDate:  checkOut,
		Status:        Confirmed,
		TotalGuests:   totalGuests,
	}
	h.Reservations = append(h.Reservations, *reservation)

	return reservation, nil

}
func (h *Hotel) GetAvailableRooms() []Room {
	var rooms []Room
	for _, room := range h.Rooms {
		if room.IsAvailable() {
			rooms = append(rooms, room)
		}
	}
	return rooms
}
func (h *Hotel) GetAvailableRoomsByType(roomType RoomType) []Room {
	var result []Room
	for _, room := range h.Rooms {
		if room.Type == roomType && room.IsAvailable() {
			result = append(result, room)
		}
	}
	return result
}

func (h *Hotel) CheckInGuest(reservationID string) error {
	for _, r := range h.Reservations {
		if r.ReservationID == reservationID {
			if r.Status != Confirmed {
				return fmt.Errorf("cannot check in reservation %s", reservationID)
			}
			r.CheckIn()
			return nil
		}
	}
	return fmt.Errorf("reservation not found")
}
func (h *Hotel) CheckOutGuest(reservationID string) error {
	for _, r := range h.Reservations {
		if r.ReservationID == reservationID {
			if r.Status != CheckedIn {
				return fmt.Errorf("cannot check out reservation %s", reservationID)
			}
			r.CheckOut()
			return nil
		}
	}
	return fmt.Errorf("reservation not found")
}

func (h *Hotel) CancelReservation(reservationID string) error {
	for _, r := range h.Reservations {
		if r.ReservationID == reservationID {
			if r.Status == CheckedOut {
				return fmt.Errorf("cannot cancel completed reservation")
			}
			r.Cancel()
			return nil
		}
	}
	return fmt.Errorf("reservation not found")
}
func (h *Hotel) GetCurrentOccupancy() float64 {
	var OccupiedPercent int

	for _, room := range h.Rooms {
		if room.Status == Occupied {
			OccupiedPercent++
		}
	}
	return (float64(OccupiedPercent) / float64(len(h.Rooms))) * 100
}

func main() {

}
