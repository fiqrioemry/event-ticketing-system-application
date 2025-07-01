package seeders

import (
	"time"

	"server/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedAll(db *gorm.DB) {
	// --- Users ---
	password, _ := bcrypt.GenerateFromPassword([]byte("123456"), 10)

	admin := models.User{
		ID:        uuid.MustParse("6b0a80e0-7609-4665-b1c3-b125b58d5b12"),
		Fullname:  "Event Admin",
		Email:     "admin@event.com",
		Password:  string(password),
		Role:      "admin",
		AvatarURL: "https://api.dicebear.com/6.x/initials/svg?seed=admin",
	}
	customer1 := models.User{
		ID:        uuid.MustParse("e267b1a5-263f-4701-a89a-91e61baf5ec0"),
		Fullname:  "Alice Customer",
		Email:     "alice@event.com",
		Password:  string(password),
		Role:      "user",
		AvatarURL: "https://api.dicebear.com/6.x/initials/svg?seed=alice",
	}
	customer2 := models.User{
		ID:        uuid.MustParse("84dfb0de-e90e-4d9d-95c3-7980d79cf66e"),
		Fullname:  "Bob Customer",
		Email:     "bob@event.com",
		Password:  string(password),
		Role:      "user",
		AvatarURL: "https://api.dicebear.com/6.x/initials/svg?seed=bob",
	}
	customer3 := models.User{
		ID:        uuid.MustParse("bdb598a3-1c86-4e95-93d2-65cda21b4b33"),
		Fullname:  "Charlie Customer",
		Email:     "charlie@event.com",
		Password:  string(password),
		Role:      "user",
		AvatarURL: "https://api.dicebear.com/6.x/initials/svg?seed=charlie",
	}
	customer4 := models.User{
		ID:        uuid.MustParse("bdb598a3-1c86-5e95-93d2-65cda21b4b33"),
		Fullname:  "Maulana Customer",
		Email:     "maulana@event.com",
		Password:  string(password),
		Role:      "user",
		AvatarURL: "https://api.dicebear.com/6.x/initials/svg?seed=maulana",
	}
	db.Create(&admin)
	db.Create(&customer1)
	db.Create(&customer2)
	db.Create(&customer3)
	db.Create(&customer4)

	eventImage1 := "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1751137753/event5_qptzmg.webp"
	eventImage2 := "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1751137753/event6_wvzoqp.webp"
	eventImage3 := "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1751137753/event6_wvzoqp.webp"
	eventImage4 := "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1751137753/event6_wvzoqp.webp"
	eventImage5 := "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1751137753/event4_n0s2kk.webp"
	eventImage6 := "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1751137752/event2_tqhnzg.webp"

	// --- Event 1 ---
	event1 := models.Event{
		ID:          uuid.MustParse("298b4cd2-4f6b-440e-8c79-3670a7634452"),
		Title:       "Jason Derulo Live in Concert",
		Image:       eventImage1,
		Description: "Jason Derulo is headlining an exciting music night with top artists. Join us for an unforgettable experience! It's going to be a night full of music, fun, and memories. All your favorite artists in one place! Mark your calendars and get ready for an amazing night!. Don't miss out on the chance to see your favorite artists live. Grab your tickets now!",
		Location:    "Jakarta",
		Date:        time.Now().AddDate(0, 0, -2),
		StartTime:   10,
		EndTime:     18,
		Status:      "active",
	}

	ticket1A := models.Ticket{
		ID:            uuid.MustParse("b93f5d2b-3216-404c-b2c0-c30d7f6d67f6"),
		EventID:       event1.ID,
		Name:          "VVIP",
		Price:         150000,
		Quota:         100,
		Limit:         2,
		Sold:          2,
		Refundable:    true,
		RefundPercent: 50,
	}
	ticket1B := models.Ticket{
		ID:            uuid.MustParse("f7c0b3ab-1aa0-40ae-b181-56cd6490ea21"),
		EventID:       event1.ID,
		Name:          "Regular",
		Price:         100000,
		Quota:         100,
		Limit:         2,
		Sold:          2,
		Refundable:    true,
		RefundPercent: 50,
	}

	db.Create(&event1)
	db.Create(&ticket1A)
	db.Create(&ticket1B)

	// --- Order 1 (Refunded) ---
	refundedAt := time.Now().AddDate(0, 0, -4)
	order1 := models.Order{
		ID:           uuid.MustParse("98b771d5-dc9b-4091-b48a-e3950c451716"),
		UserID:       customer1.ID,
		EventID:      event1.ID,
		Fullname:     "Alice Customer",
		Email:        "alice@event.com",
		Phone:        "08123456781",
		TotalPrice:   250000,
		Status:       "paid",
		IsRefunded:   true,
		RefundedAt:   &refundedAt,
		RefundAmount: 125000,
		RefundReason: "Customer request",
		CreatedAt:    time.Now().AddDate(0, 0, -5),
		UpdatedAt:    time.Now().AddDate(0, 0, -5),
	}
	db.Create(&order1)
	db.Create(&models.OrderDetail{
		ID:         uuid.MustParse("0e3324e3-e09d-470a-aee2-fb9399b3b5d2"),
		OrderID:    order1.ID,
		TicketID:   ticket1A.ID,
		TicketName: "VVIP",
		Quantity:   1,
		Price:      150000,
		CreatedAt:  time.Now().AddDate(0, 0, -5),
	})
	db.Create(&models.OrderDetail{
		ID:         uuid.MustParse("7c64b0ea-3528-4e7f-8cd3-79063d104cd7"),
		OrderID:    order1.ID,
		TicketID:   ticket1B.ID,
		TicketName: "Regular",
		Quantity:   1,
		Price:      100000,
		CreatedAt:  time.Now().AddDate(0, 0, -5),
	})

	payment1PaidAt := time.Now().AddDate(0, 0, -5)
	db.Create(&models.Payment{
		ID:       uuid.MustParse("80e2e3a7-e231-49ae-8980-885e4c47d0a6"),
		UserID:   customer1.ID,
		OrderID:  order1.ID,
		Fullname: "Alice Customer",
		Email:    "alice@event.com",
		Method:   "stripe",
		Amount:   250000,
		Status:   "paid",
		PaidAt:   &payment1PaidAt,
	})
	db.Create(&models.UserTicket{
		ID:       uuid.MustParse("3896ee3e-d5e1-42f2-8661-b4ae64f429b4"),
		UserID:   customer1.ID,
		EventID:  event1.ID,
		TicketID: ticket1A.ID,
		QRCode:   "QR-3896ee3e-d5e1-42f2-8661-b4ae64f429b4",
	})

	db.Create(&models.UserTicket{
		ID:       uuid.MustParse("2f36f12a-8eeb-4cfc-b1c0-8d25e81e06a1"),
		UserID:   customer1.ID,
		EventID:  event1.ID,
		TicketID: ticket1B.ID,
		QRCode:   "QR-2f36f12a-8eeb-4cfc-b1c0-8d25e81e06a1",
	})

	approvedOrder1 := time.Now().AddDate(0, 0, -4)
	db.Create(&models.WithdrawalRequest{
		ID:         uuid.MustParse("c1c73b88-7b5e-472c-8bd9-cab6e45b9e99"),
		UserID:     customer1.ID,
		Amount:     125000,
		Status:     "approved",
		Reason:     "Refund event",
		CreatedAt:  time.Now().AddDate(0, 0, -4),
		ApprovedAt: &approvedOrder1,
	})

	// --- Event 2 ---
	upcomingDate := time.Now().AddDate(0, 0, 5)
	newEvent := models.Event{
		ID:          uuid.MustParse("ddaf8eb0-a68e-4316-8dc7-834d183faaf6"),
		Title:       "Bandung Acoustic Night 2025",
		Image:       eventImage2,
		Description: "Biggest Acoustic Festival. Join us for an unforgettable evening filled with soothing acoustic performances by talented artists. Experience the magic of live music in a cozy atmosphere. Don't miss out on this musical journey! Tickets are selling fast, grab yours now!",
		Location:    "Bandung",
		Date:        upcomingDate,
		StartTime:   18,
		EndTime:     22,
		Status:      "active",
		CreatedAt:   time.Now().AddDate(0, 0, -2),
		UpdatedAt:   time.Now().AddDate(0, 0, -2),
	}
	db.Create(&newEvent)

	// Event 2 : tickets
	ticketA := models.Ticket{
		ID:            uuid.MustParse("d39ff313-db85-4a3f-9f33-fa8fb0c68019"),
		EventID:       uuid.MustParse("ddaf8eb0-a68e-4316-8dc7-834d183faaf6"),
		Name:          "VIP GOLD PACK",
		Price:         120000,
		Quota:         50,
		Limit:         2,
		Sold:          5,
		Refundable:    true,
		RefundPercent: 70,
	}
	ticketB := models.Ticket{
		ID:            uuid.MustParse("802a2b63-d941-4557-8dc7-0d0580d0580f"),
		EventID:       uuid.MustParse("ddaf8eb0-a68e-4316-8dc7-834d183faaf6"),
		Name:          "REGULAR SILVER PACK",
		Price:         90000,
		Quota:         50,
		Limit:         2,
		Sold:          3,
		Refundable:    true,
		RefundPercent: 70,
	}
	db.Create(&ticketA)
	db.Create(&ticketB)

	// --- Order 2 ---
	newOrder := models.Order{
		ID:           uuid.MustParse("24525c25-1944-4d57-8bbd-c4169cc46f04"),
		UserID:       uuid.MustParse("bdb598a3-1c86-4e95-93d2-65cda21b4b33"),
		EventID:      uuid.MustParse("ddaf8eb0-a68e-4316-8dc7-834d183faaf6"),
		Fullname:     "Charlie Customer",
		Email:        "charlie@event.com",
		Phone:        "08123456788",
		TotalPrice:   210000,
		IsRefunded:   false,
		RefundedAt:   nil,
		RefundAmount: 0,
		RefundReason: "",
		CreatedAt:    time.Now().AddDate(0, 0, -1),
		UpdatedAt:    time.Now().AddDate(0, 0, -1),
		Status:       "paid",
	}
	db.Create(&newOrder)
	// order 2 : order details
	db.Create(&models.OrderDetail{
		ID:         uuid.MustParse("798186a9-2a2e-42cd-bcd4-8ea98f47dbf5"),
		OrderID:    uuid.MustParse("24525c25-1944-4d57-8bbd-c4169cc46f04"),
		TicketID:   uuid.MustParse("d39ff313-db85-4a3f-9f33-fa8fb0c68019"),
		TicketName: "VIP GOLD PACK",
		Quantity:   1,
		Price:      120000,
		CreatedAt:  time.Now().AddDate(0, 0, -1),
	})

	db.Create(&models.OrderDetail{
		ID:         uuid.MustParse("5c6a27fe-82c2-4135-be8b-cfc40d1b57d4"),
		OrderID:    uuid.MustParse("24525c25-1944-4d57-8bbd-c4169cc46f04"),
		TicketID:   uuid.MustParse("802a2b63-d941-4557-8dc7-0d0580d0580f"),
		TicketName: "REGULAR SILVER PACK",
		Quantity:   1,
		Price:      90000,
		CreatedAt:  time.Now().AddDate(0, 0, -1),
	})

	// order 2 : payment
	order2PaidAt := time.Now().AddDate(0, 0, -1)
	db.Create(&models.Payment{
		ID:       uuid.MustParse("2ed5aec9-af02-4573-8fe0-c0f376af9bc1"),
		UserID:   uuid.MustParse("bdb598a3-1c86-4e95-93d2-65cda21b4b33"),
		OrderID:  uuid.MustParse("24525c25-1944-4d57-8bbd-c4169cc46f04"),
		Fullname: "Charlie Customer",
		Email:    "charlie@event.com",
		Method:   "stripe",
		Amount:   210000,
		Status:   "paid",
		PaidAt:   &order2PaidAt,
	})

	// order 2 : user tickets
	db.Create(&models.UserTicket{
		ID:       uuid.MustParse("7e3d5fe9-6915-46d2-a117-7eb5d645e3f6"),
		UserID:   uuid.MustParse("bdb598a3-1c86-4e95-93d2-65cda21b4b33"),
		EventID:  uuid.MustParse("ddaf8eb0-a68e-4316-8dc7-834d183faaf6"),
		TicketID: uuid.MustParse("d39ff313-db85-4a3f-9f33-fa8fb0c68019"),
		QRCode:   "QR-7e3d5fe9-6915-46d2-a117-7eb5d645e3f6",
	})

	// order 2 : user tickets
	db.Create(&models.UserTicket{
		ID:       uuid.MustParse("04740e8d-03dd-41d6-ba3f-e1031e59cc88"),
		UserID:   uuid.MustParse("bdb598a3-1c86-4e95-93d2-65cda21b4b33"),
		EventID:  uuid.MustParse("ddaf8eb0-a68e-4316-8dc7-834d183faaf6"),
		TicketID: uuid.MustParse("802a2b63-d941-4557-8dc7-0d0580d0580f"),
		QRCode:   "QR-802a2b63-d941-4557-8dc7-0d0580d0580f",
	})

	// --- Event 3 ---
	upcoming3 := time.Now().AddDate(0, 0, 6)
	event3 := models.Event{
		ID:          uuid.MustParse("ddaf8eb0-a68e-4316-92c7-834d183faaf6"),
		Title:       eventImage3,
		Image:       "https://placehold.co/400x400?text=Acoustic+Night",
		Description: "Join us for an unforgettable evening filled with live music performances by top artists. Experience the energy and excitement of a concert like never before! Don't miss out on this musical extravaganza! Get your tickets now!. All your favorite artists in one place! Mark your calendars and get ready for an amazing night! Tickets are selling fast, grab yours now!",
		Location:    "Jakarta",
		Date:        upcoming3,
		StartTime:   18,
		EndTime:     22,
		Status:      "active",
		CreatedAt:   time.Now().AddDate(0, 0, -1),
		UpdatedAt:   time.Now().AddDate(0, 0, -1),
	}
	db.Create(&newEvent)

	// Event 3 : tickets
	ticketBEvent3 := models.Ticket{
		ID:            uuid.MustParse("d39ff313-db85-4a3f-9f33-fa8fb0c68019"),
		EventID:       event3.ID,
		Name:          "VIP GOLD PACK",
		Price:         550000,
		Quota:         50,
		Limit:         2,
		Sold:          40,
		Refundable:    false,
		RefundPercent: 0,
	}
	ticketAEvent3 := models.Ticket{
		ID:            uuid.MustParse("102a2b63-d941-4557-8dc7-0d0580d0580f"),
		EventID:       event3.ID,
		Name:          "REGULAR SILVER PACK",
		Price:         250000,
		Quota:         50,
		Limit:         2,
		Sold:          15,
		Refundable:    false,
		RefundPercent: 0,
	}
	db.Create(&ticketAEvent3)
	db.Create(&ticketBEvent3)

	// --- Event 4 ---
	event4 := models.Event{
		ID:          uuid.MustParse("aabbccdd-eeff-4400-8899-aabbccddeeff"),
		Title:       "Surabaya Music Carnival 2025",
		Image:       eventImage4,
		Description: "A spectacular music carnival with various genres, food stalls, and art installations. Perfect for music lovers and culture seekers!.Join us for an unforgettable evening filled with live music performances by top artists. Experience the energy and excitement of a concert like never before! Don't miss out on this musical extravaganza! Get your tickets now!. All your favorite artists in one place! Mark your calendars and get ready for an amazing night! Tickets are selling fast, grab yours now!",
		Location:    "Surabaya",
		Date:        time.Now().AddDate(0, 0, 7),
		StartTime:   15,
		EndTime:     23,
		Status:      "active",
		CreatedAt:   time.Now().AddDate(0, 0, -1),
		UpdatedAt:   time.Now().AddDate(0, 0, -1),
	}
	db.Create(&event4)

	// Event 4 : tickets
	ticket4VIP := models.Ticket{
		ID:            uuid.MustParse("99aa88bb-ccdd-4eee-aaaa-112233445566"),
		EventID:       event4.ID,
		Name:          "VIP TICKET",
		Price:         300000,
		Quota:         100,
		Limit:         2,
		Sold:          10,
		Refundable:    true,
		RefundPercent: 50,
	}
	ticket4Reg := models.Ticket{
		ID:            uuid.MustParse("77cc66dd-bbaa-4fff-8899-998877665544"),
		EventID:       event4.ID,
		Name:          "REGULAR TICKET",
		Price:         180000,
		Quota:         100,
		Limit:         2,
		Sold:          20,
		Refundable:    true,
		RefundPercent: 50,
	}
	db.Create(&ticket4VIP)
	db.Create(&ticket4Reg)

	// --- Order for Event 4 by customer4 ---
	event4Order := models.Order{
		ID:         uuid.MustParse("de44ccaa-11ee-47cc-bb44-556677889900"),
		UserID:     customer4.ID,
		EventID:    event4.ID,
		Fullname:   "Maulana Customer",
		Email:      "maulana@event.com",
		Phone:      "08123456789",
		TotalPrice: 480000,
		Status:     "paid",
		IsRefunded: false,
		CreatedAt:  time.Now().AddDate(0, 0, -1),
		UpdatedAt:  time.Now().AddDate(0, 0, -1),
	}
	db.Create(&event4Order)

	// Order details
	db.Create(&models.OrderDetail{
		ID:         uuid.MustParse("a1b2c3d4-e5f6-47aa-8888-112233445577"),
		OrderID:    event4Order.ID,
		TicketID:   ticket4VIP.ID,
		TicketName: "VIP TICKET",
		Quantity:   1,
		Price:      300000,
		CreatedAt:  time.Now().AddDate(0, 0, -1),
	})
	db.Create(&models.OrderDetail{
		ID:         uuid.MustParse("b2c3d4e5-f6a7-4888-bbbb-223344556688"),
		OrderID:    event4Order.ID,
		TicketID:   ticket4Reg.ID,
		TicketName: "REGULAR TICKET",
		Quantity:   1,
		Price:      180000,
		CreatedAt:  time.Now().AddDate(0, 0, -1),
	})

	// Payment
	event4PaidAt := time.Now().AddDate(0, 0, -1)
	db.Create(&models.Payment{
		ID:       uuid.MustParse("c3d4e5f6-a788-4bbb-aaaa-334455667799"),
		UserID:   uuid.MustParse("bdb598a3-1c86-5e95-93d2-65cda21b4b33"),
		OrderID:  event4Order.ID,
		Fullname: "Maulana Customer",
		Email:    "maulana@event.com",
		Method:   "stripe",
		Amount:   480000,
		Status:   "paid",
		PaidAt:   &event4PaidAt,
	})

	// User Tickets
	db.Create(&models.UserTicket{
		ID:       uuid.MustParse("d4e5f6a7-8899-4ccc-aaaa-445566778800"),
		UserID:   uuid.MustParse("bdb598a3-1c86-5e95-93d2-65cda21b4b33"),
		EventID:  event4.ID,
		TicketID: ticket4VIP.ID,
		QRCode:   "QR-d4e5f6a7-8899-4ccc-aaaa-445566778800",
	})
	db.Create(&models.UserTicket{
		ID:       uuid.MustParse("e5f6a788-99aa-4ddd-bbbb-556677889911"),
		UserID:   uuid.MustParse("bdb598a3-1c86-5e95-93d2-65cda21b4b33"),
		EventID:  event4.ID,
		TicketID: ticket4Reg.ID,
		QRCode:   "QR-e5f6a788-99aa-4ddd-bbbb-556677889911",
	})

	// --- Event 5 ---
	event5 := models.Event{
		ID:          uuid.New(),
		Title:       "Jakarta Acoustic Festival Award",
		Image:       eventImage5,
		Description: "A spectacular music carnival with various genres, food stalls, and art installations. Perfect for music lovers and culture seekers!",
		Location:    "Jakarta",
		Date:        time.Now().AddDate(0, 0, 5),
		StartTime:   13,
		EndTime:     18,
		Status:      "active",
		CreatedAt:   time.Now().AddDate(0, 0, -1),
		UpdatedAt:   time.Now().AddDate(0, 0, -1),
	}
	db.Create(&event5)

	// Event 5 : tickets
	ticket5VIP := models.Ticket{
		ID:            uuid.New(),
		EventID:       event5.ID,
		Name:          "VIP TICKET",
		Price:         300000,
		Quota:         100,
		Limit:         2,
		Sold:          10,
		Refundable:    true,
		RefundPercent: 50,
	}
	ticket5Reg := models.Ticket{
		ID:            uuid.New(),
		EventID:       event5.ID,
		Name:          "REGULAR TICKET",
		Price:         180000,
		Quota:         100,
		Limit:         2,
		Sold:          20,
		Refundable:    true,
		RefundPercent: 50,
	}
	db.Create(&ticket5VIP)
	db.Create(&ticket5Reg)

	// --- Event 6 ---
	event6 := models.Event{
		ID:          uuid.New(),
		Title:       "Bandung Reggae Festival ",
		Image:       eventImage6,
		Description: "Wonderful reggae festival with top artists, food stalls, and art installations. Perfect for music lovers and culture seekers! Join us for an unforgettable evening filled with live music performances by top artists. Experience the energy and excitement of a concert like never before! Don't miss out on this musical extravaganza! Get your tickets now!. All your favorite artists in one place! Mark your calendars and get ready for an amazing night! Tickets are selling fast, grab yours now!",
		Location:    "Bandung",
		Date:        time.Now().AddDate(0, 0, 5),
		StartTime:   12,
		EndTime:     15,
		Status:      "active",
		CreatedAt:   time.Now().AddDate(0, 0, -1),
		UpdatedAt:   time.Now().AddDate(0, 0, -1),
	}
	db.Create(&event6)

	// Event 6 : tickets
	ticket6VIP := models.Ticket{
		ID:            uuid.New(),
		EventID:       event6.ID,
		Name:          "VIP TICKET",
		Price:         300000,
		Quota:         100,
		Limit:         2,
		Sold:          10,
		Refundable:    true,
		RefundPercent: 50,
	}
	ticket6Reg := models.Ticket{
		ID:            uuid.New(),
		EventID:       event6.ID,
		Name:          "REGULAR TICKET",
		Price:         180000,
		Quota:         100,
		Limit:         2,
		Sold:          20,
		Refundable:    true,
		RefundPercent: 50,
	}
	db.Create(&ticket6VIP)
	db.Create(&ticket6Reg)

}
