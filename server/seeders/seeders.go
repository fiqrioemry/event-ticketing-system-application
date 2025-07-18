package seeders

import (
	"time"

	"github.com/fiqrioemry/event_ticketing_system_app/server/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedAll(db *gorm.DB) {
	// --- Users ---
	password, _ := bcrypt.GenerateFromPassword([]byte("123456"), 10)

	admin := models.User{
		ID:       uuid.MustParse("6b0a80e0-7609-4665-b1c3-b125b58d5b12"),
		Fullname: "Event Admin",
		Email:    "admin@event.com",
		Password: string(password),
		Role:     "admin",
		Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=admin",
	}
	customer1 := models.User{
		ID:       uuid.MustParse("e267b1a5-263f-4701-a89a-91e61baf5ec0"),
		Fullname: "Alice Customer",
		Email:    "alice@event.com",
		Password: string(password),
		Role:     "user",
		Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=alice",
	}
	customer2 := models.User{
		ID:       uuid.MustParse("84dfb0de-e90e-4d9d-95c3-7980d79cf66e"),
		Fullname: "Bob Customer",
		Email:    "bob@event.com",
		Password: string(password),
		Role:     "user",
		Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=bob",
	}
	customer3 := models.User{
		ID:       uuid.MustParse("bdb598a3-1c86-4e95-93d2-65cda21b4b33"),
		Fullname: "Charlie Customer",
		Email:    "charlie@event.com",
		Password: string(password),
		Role:     "user",
		Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=charlie",
	}
	customer4 := models.User{
		ID:       uuid.MustParse("bdb598a3-1c86-5e95-93d2-65cda21b4b33"),
		Fullname: "Maulana Customer",
		Email:    "maulana@event.com",
		Password: string(password),
		Role:     "user",
		Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=maulana",
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

func SeedAdditionalEvents(db *gorm.DB) {
	// Image URLs
	eventImages := []string{
		"https://res.cloudinary.com/dnlrzhdcs/image/upload/v1752843488/ticketing_app/8_asneai.webp",
		"https://res.cloudinary.com/dnlrzhdcs/image/upload/v1752843489/ticketing_app/19_nucf7u.webp",
		"https://res.cloudinary.com/dnlrzhdcs/image/upload/v1752843489/ticketing_app/7_k4hiry.webp",
		"https://res.cloudinary.com/dnlrzhdcs/image/upload/v1752843489/ticketing_app/11_lcxueu.webp",
		"https://res.cloudinary.com/dnlrzhdcs/image/upload/v1752843489/ticketing_app/12_o0tkfz.webp",
		"https://res.cloudinary.com/dnlrzhdcs/image/upload/v1752843488/ticketing_app/13_ozfxnr.webp",
		"https://res.cloudinary.com/dnlrzhdcs/image/upload/v1752843489/ticketing_app/9_lwhqly.webp",
		"https://res.cloudinary.com/dnlrzhdcs/image/upload/v1752843488/ticketing_app/15_rzbgc9.webp",
		"https://res.cloudinary.com/dnlrzhdcs/image/upload/v1752843488/ticketing_app/10_jxggqh.webp",
		"https://res.cloudinary.com/dnlrzhdcs/image/upload/v1752843488/ticketing_app/17_ilyojt.webp",
		"https://res.cloudinary.com/dnlrzhdcs/image/upload/v1752843488/ticketing_app/14_u2jlti.webp",
		"https://res.cloudinary.com/dnlrzhdcs/image/upload/v1752843487/ticketing_app/18_x03whb.webp",
	}

	// Event 7 - Jakarta Jazz Festival 2025
	event7 := models.Event{
		ID:          uuid.New(),
		Title:       "Jakarta Jazz Festival 2025",
		Image:       eventImages[0],
		Description: "The biggest jazz festival in Southeast Asia returns! Experience world-class jazz performances by international and local artists. Three days of non-stop jazz music, food courts, and cultural exhibitions. This year featuring Grammy-winning artists and emerging talents from across the globe. Don't miss this legendary musical celebration!",
		Location:    "Jakarta",
		Date:        time.Now().AddDate(0, 0, 4),
		StartTime:   16,
		EndTime:     23,
		Status:      "active",
		CreatedAt:   time.Now().AddDate(0, 0, -2),
		UpdatedAt:   time.Now().AddDate(0, 0, -2),
	}
	db.Create(&event7)

	// Event 7 Tickets
	event7VIP := models.Ticket{
		ID:            uuid.New(),
		EventID:       event7.ID,
		Name:          "VIP JAZZ PASS",
		Price:         750000,
		Quota:         150,
		Limit:         2,
		Sold:          45,
		Refundable:    true,
		RefundPercent: 80,
	}
	event7Regular := models.Ticket{
		ID:            uuid.New(),
		EventID:       event7.ID,
		Name:          "GENERAL ADMISSION",
		Price:         350000,
		Quota:         300,
		Limit:         4,
		Sold:          120,
		Refundable:    true,
		RefundPercent: 70,
	}
	db.Create(&event7VIP)
	db.Create(&event7Regular)

	// Event 8 - Bali Sunset Electronic Music Festival
	event8 := models.Event{
		ID:          uuid.New(),
		Title:       "Bali Sunset Electronic Music Festival",
		Image:       eventImages[1],
		Description: "Dance under the stars at Bali's most spectacular electronic music festival! World-renowned DJs spinning the latest beats as the sun sets over the Indian Ocean. Multiple stages featuring house, techno, progressive, and ambient music. Beach side location with international food vendors and art installations. An unforgettable tropical electronic music experience!",
		Location:    "Bali",
		Date:        time.Now().AddDate(0, 0, 6),
		StartTime:   17,
		EndTime:     2,
		Status:      "active",
		CreatedAt:   time.Now().AddDate(0, 0, -1),
		UpdatedAt:   time.Now().AddDate(0, 0, -1),
	}
	db.Create(&event8)

	// Event 8 Tickets
	event8VIP := models.Ticket{
		ID:            uuid.New(),
		EventID:       event8.ID,
		Name:          "SUNSET VIP DECK",
		Price:         850000,
		Quota:         100,
		Limit:         2,
		Sold:          35,
		Refundable:    true,
		RefundPercent: 75,
	}
	event8Regular := models.Ticket{
		ID:            uuid.New(),
		EventID:       event8.ID,
		Name:          "BEACH ACCESS",
		Price:         450000,
		Quota:         500,
		Limit:         6,
		Sold:          180,
		Refundable:    true,
		RefundPercent: 65,
	}
	db.Create(&event8VIP)
	db.Create(&event8Regular)

	// Event 9 - Bandung Indie Rock Revolution
	event9 := models.Event{
		ID:          uuid.New(),
		Title:       "Bandung Indie Rock Revolution",
		Image:       eventImages[2],
		Description: "The ultimate indie rock festival featuring the best underground bands from Indonesia and neighboring countries. Discover new sounds, support emerging artists, and rock out to alternative music in the cool mountain air of Bandung. Food trucks, merchandise stalls, and meet-and-greet sessions with your favorite indie artists. A celebration of independent music culture!",
		Location:    "Bandung",
		Date:        time.Now().AddDate(0, 0, 8),
		StartTime:   14,
		EndTime:     22,
		Status:      "active",
		CreatedAt:   time.Now().AddDate(0, 0, -1),
		UpdatedAt:   time.Now().AddDate(0, 0, -1),
	}
	db.Create(&event9)

	// Event 9 Tickets
	event9VIP := models.Ticket{
		ID:            uuid.New(),
		EventID:       event9.ID,
		Name:          "INDIE ROCK VIP",
		Price:         400000,
		Quota:         80,
		Limit:         2,
		Sold:          25,
		Refundable:    true,
		RefundPercent: 70,
	}
	event9Regular := models.Ticket{
		ID:            uuid.New(),
		EventID:       event9.ID,
		Name:          "ROCK FEST PASS",
		Price:         220000,
		Quota:         400,
		Limit:         4,
		Sold:          150,
		Refundable:    true,
		RefundPercent: 60,
	}
	db.Create(&event9VIP)
	db.Create(&event9Regular)

	// Event 10 - Surabaya Hip Hop Championship
	event10 := models.Event{
		ID:          uuid.New(),
		Title:       "Surabaya Hip Hop Championship 2025",
		Image:       eventImages[3],
		Description: "The biggest hip hop battle and showcase in East Java! Featuring rap battles, breakdancing competitions, graffiti exhibitions, and live performances by top Indonesian hip hop artists. Witness the raw talent of street culture and urban arts. Special guest appearances by legendary hip hop pioneers. Represent your city and show your skills!",
		Location:    "Surabaya",
		Date:        time.Now().AddDate(0, 0, 5),
		StartTime:   15,
		EndTime:     23,
		Status:      "active",
		CreatedAt:   time.Now().AddDate(0, 0, -2),
		UpdatedAt:   time.Now().AddDate(0, 0, -2),
	}
	db.Create(&event10)

	// Event 10 Tickets
	event10VIP := models.Ticket{
		ID:            uuid.New(),
		EventID:       event10.ID,
		Name:          "HIP HOP VIP PASS",
		Price:         320000,
		Quota:         60,
		Limit:         2,
		Sold:          18,
		Refundable:    true,
		RefundPercent: 65,
	}
	event10Regular := models.Ticket{
		ID:            uuid.New(),
		EventID:       event10.ID,
		Name:          "STREET CULTURE PASS",
		Price:         180000,
		Quota:         250,
		Limit:         4,
		Sold:          85,
		Refundable:    true,
		RefundPercent: 55,
	}
	db.Create(&event10VIP)
	db.Create(&event10Regular)

	// Event 11 - Yogyakarta Traditional Music Fusion
	event11 := models.Event{
		ID:          uuid.New(),
		Title:       "Yogyakarta Traditional Music Fusion Festival",
		Image:       eventImages[4],
		Description: "A unique fusion of traditional Javanese music with modern genres. Experience gamelan orchestras collaborating with rock bands, traditional singers with electronic music, and cultural dance performances. Celebrating the rich heritage of Yogyakarta while embracing contemporary musical innovation. Cultural workshops and traditional craft exhibitions included!",
		Location:    "Yogyakarta",
		Date:        time.Now().AddDate(0, 0, 7),
		StartTime:   16,
		EndTime:     22,
		Status:      "active",
		CreatedAt:   time.Now().AddDate(0, 0, -1),
		UpdatedAt:   time.Now().AddDate(0, 0, -1),
	}
	db.Create(&event11)

	// Event 11 Tickets
	event11VIP := models.Ticket{
		ID:            uuid.New(),
		EventID:       event11.ID,
		Name:          "CULTURAL VIP EXPERIENCE",
		Price:         380000,
		Quota:         70,
		Limit:         2,
		Sold:          22,
		Refundable:    true,
		RefundPercent: 70,
	}
	event11Regular := models.Ticket{
		ID:            uuid.New(),
		EventID:       event11.ID,
		Name:          "FUSION FESTIVAL PASS",
		Price:         200000,
		Quota:         300,
		Limit:         3,
		Sold:          95,
		Refundable:    true,
		RefundPercent: 60,
	}
	db.Create(&event11VIP)
	db.Create(&event11Regular)

	// Event 12 - Jakarta K-Pop Super Concert
	event12 := models.Event{
		ID:          uuid.New(),
		Title:       "Jakarta K-Pop Super Concert 2025",
		Image:       eventImages[5],
		Description: "The ultimate K-Pop experience in Indonesia! Top Korean idol groups, solo artists, and special collaborations with Indonesian artists. High-energy performances, fan interactions, merchandise booths, and Korean cultural exhibitions. Professional stage production with LED screens, pyrotechnics, and synchronized lighting. A dream come true for K-Pop fans!",
		Location:    "Jakarta",
		Date:        time.Now().AddDate(0, 0, 9),
		StartTime:   18,
		EndTime:     22,
		Status:      "active",
		CreatedAt:   time.Now().AddDate(0, 0, -1),
		UpdatedAt:   time.Now().AddDate(0, 0, -1),
	}
	db.Create(&event12)

	// Event 12 Tickets
	event12VIP := models.Ticket{
		ID:            uuid.New(),
		EventID:       event12.ID,
		Name:          "K-POP VIP PLATINUM",
		Price:         950000,
		Quota:         120,
		Limit:         2,
		Sold:          75,
		Refundable:    true,
		RefundPercent: 80,
	}
	event12Regular := models.Ticket{
		ID:            uuid.New(),
		EventID:       event12.ID,
		Name:          "HALLYU WAVE PASS",
		Price:         480000,
		Quota:         800,
		Limit:         4,
		Sold:          320,
		Refundable:    true,
		RefundPercent: 70,
	}
	db.Create(&event12VIP)
	db.Create(&event12Regular)

	// Event 13 - Bali Reggae Beach Festival
	event13 := models.Event{
		ID:          uuid.New(),
		Title:       "Bali Reggae Beach Festival",
		Image:       eventImages[6],
		Description: "Feel the positive vibes at Bali's premier reggae festival! International and local reggae artists performing on a beautiful beachfront stage. Rastafarian culture celebrations, jamaican food stalls, artisan markets, and spiritual workshops. Watch the sunset while grooving to authentic reggae rhythms. One love, one heart, one festival!",
		Location:    "Bali",
		Date:        time.Now().AddDate(0, 0, 10),
		StartTime:   15,
		EndTime:     23,
		Status:      "active",
		CreatedAt:   time.Now().AddDate(0, 0, -1),
		UpdatedAt:   time.Now().AddDate(0, 0, -1),
	}
	db.Create(&event13)

	// Event 13 Tickets
	event13VIP := models.Ticket{
		ID:            uuid.New(),
		EventID:       event13.ID,
		Name:          "REGGAE VIP ISLAND",
		Price:         520000,
		Quota:         90,
		Limit:         2,
		Sold:          30,
		Refundable:    true,
		RefundPercent: 75,
	}
	event13Regular := models.Ticket{
		ID:            uuid.New(),
		EventID:       event13.ID,
		Name:          "BEACH REGGAE PASS",
		Price:         280000,
		Quota:         350,
		Limit:         4,
		Sold:          115,
		Refundable:    true,
		RefundPercent: 65,
	}
	db.Create(&event13VIP)
	db.Create(&event13Regular)

	// Event 14 - Bandung Food & Music Carnival
	event14 := models.Event{
		ID:          uuid.New(),
		Title:       "Bandung Food & Music Carnival 2025",
		Image:       eventImages[7],
		Description: "A perfect combination of culinary delights and live music! Over 100 food vendors featuring local Bandung specialties, Indonesian street food, and international cuisine. Live acoustic performances, cooking demonstrations, food competitions, and family-friendly entertainment. Celebrate the rich food culture of Bandung with great music!",
		Location:    "Bandung",
		Date:        time.Now().AddDate(0, 0, 3),
		StartTime:   12,
		EndTime:     21,
		Status:      "active",
		CreatedAt:   time.Now().AddDate(0, 0, -1),
		UpdatedAt:   time.Now().AddDate(0, 0, -1),
	}
	db.Create(&event14)

	// Event 14 Tickets
	event14VIP := models.Ticket{
		ID:            uuid.New(),
		EventID:       event14.ID,
		Name:          "FOODIE VIP EXPERIENCE",
		Price:         280000,
		Quota:         100,
		Limit:         2,
		Sold:          40,
		Refundable:    true,
		RefundPercent: 70,
	}
	event14Regular := models.Ticket{
		ID:            uuid.New(),
		EventID:       event14.ID,
		Name:          "CARNIVAL ENTRY PASS",
		Price:         120000,
		Quota:         500,
		Limit:         6,
		Sold:          200,
		Refundable:    true,
		RefundPercent: 60,
	}
	db.Create(&event14VIP)
	db.Create(&event14Regular)

	// Event 15 - Surabaya Metal Underground
	event15 := models.Event{
		ID:          uuid.New(),
		Title:       "Surabaya Metal Underground Fest",
		Image:       eventImages[8],
		Description: "The most brutal metal festival in East Java! Death metal, black metal, thrash metal, and hardcore bands from across Indonesia and Southeast Asia. Mosh pits, headbanging, and raw underground energy. Support the metal scene and witness devastating live performances. Not for the faint-hearted. Horns up!",
		Location:    "Surabaya",
		Date:        time.Now().AddDate(0, 0, 11),
		StartTime:   17,
		EndTime:     24,
		Status:      "active",
		CreatedAt:   time.Now().AddDate(0, 0, -1),
		UpdatedAt:   time.Now().AddDate(0, 0, -1),
	}
	db.Create(&event15)

	// Event 15 Tickets
	event15VIP := models.Ticket{
		ID:            uuid.New(),
		EventID:       event15.ID,
		Name:          "METAL VIP PIT PASS",
		Price:         350000,
		Quota:         50,
		Limit:         2,
		Sold:          15,
		Refundable:    true,
		RefundPercent: 60,
	}
	event15Regular := models.Ticket{
		ID:            uuid.New(),
		EventID:       event15.ID,
		Name:          "UNDERGROUND PASS",
		Price:         180000,
		Quota:         200,
		Limit:         3,
		Sold:          70,
		Refundable:    true,
		RefundPercent: 50,
	}
	db.Create(&event15VIP)
	db.Create(&event15Regular)

	// Event 16 - Yogyakarta Student Music Festival
	event16 := models.Event{
		ID:          uuid.New(),
		Title:       "Yogyakarta Student Music Festival 2025",
		Image:       eventImages[9],
		Description: "By students, for students! The biggest student music festival in Indonesia featuring university bands, solo artists, and student music communities from across the archipelago. Battle of the bands competition, music workshops, campus radio showcases, and networking sessions. Affordable tickets for the student community!",
		Location:    "Yogyakarta",
		Date:        time.Now().AddDate(0, 0, 4),
		StartTime:   14,
		EndTime:     22,
		Status:      "active",
		CreatedAt:   time.Now().AddDate(0, 0, -2),
		UpdatedAt:   time.Now().AddDate(0, 0, -2),
	}
	db.Create(&event16)

	// Event 16 Tickets
	event16VIP := models.Ticket{
		ID:            uuid.New(),
		EventID:       event16.ID,
		Name:          "STUDENT VIP PACKAGE",
		Price:         150000,
		Quota:         80,
		Limit:         2,
		Sold:          35,
		Refundable:    true,
		RefundPercent: 70,
	}
	event16Regular := models.Ticket{
		ID:            uuid.New(),
		EventID:       event16.ID,
		Name:          "STUDENT PASS",
		Price:         75000,
		Quota:         600,
		Limit:         4,
		Sold:          280,
		Refundable:    true,
		RefundPercent: 60,
	}
	db.Create(&event16VIP)
	db.Create(&event16Regular)

	// Event 17 - Jakarta International Folk Festival
	event17 := models.Event{
		ID:          uuid.New(),
		Title:       "Jakarta International Folk Festival",
		Image:       eventImages[10],
		Description: "Celebrate world folk music traditions! Artists from different countries showcasing their cultural heritage through music and dance. Indonesian traditional music, international folk bands, cultural exhibitions, and artisan craft markets. Learn about diverse musical traditions while enjoying authentic performances from around the globe.",
		Location:    "Jakarta",
		Date:        time.Now().AddDate(0, 0, 12),
		StartTime:   15,
		EndTime:     21,
		Status:      "active",
		CreatedAt:   time.Now().AddDate(0, 0, -1),
		UpdatedAt:   time.Now().AddDate(0, 0, -1),
	}
	db.Create(&event17)

	// Event 17 Tickets
	event17VIP := models.Ticket{
		ID:            uuid.New(),
		EventID:       event17.ID,
		Name:          "INTERNATIONAL VIP",
		Price:         420000,
		Quota:         90,
		Limit:         2,
		Sold:          25,
		Refundable:    true,
		RefundPercent: 75,
	}
	event17Regular := models.Ticket{
		ID:            uuid.New(),
		EventID:       event17.ID,
		Name:          "FOLK FESTIVAL PASS",
		Price:         220000,
		Quota:         400,
		Limit:         3,
		Sold:          120,
		Refundable:    true,
		RefundPercent: 65,
	}
	db.Create(&event17VIP)
	db.Create(&event17Regular)

	// Event 18 - Bali Acoustic Sunset Sessions
	event18 := models.Event{
		ID:          uuid.New(),
		Title:       "Bali Acoustic Sunset Sessions",
		Image:       eventImages[11],
		Description: "Intimate acoustic performances as the sun sets over Bali's beautiful coastline. Solo artists, duos, and acoustic bands performing original songs and covers in a relaxed, unplugged atmosphere. Bring your blanket, enjoy local refreshments, and experience music in its purest form. A perfect evening of music and nature!",
		Location:    "Bali",
		Date:        time.Now().AddDate(0, 0, 6),
		StartTime:   17,
		EndTime:     20,
		Status:      "active",
		CreatedAt:   time.Now().AddDate(0, 0, -1),
		UpdatedAt:   time.Now().AddDate(0, 0, -1),
	}
	db.Create(&event18)

	// Event 18 Tickets
	event18VIP := models.Ticket{
		ID:            uuid.New(),
		EventID:       event18.ID,
		Name:          "SUNSET VIP LOUNGE",
		Price:         320000,
		Quota:         60,
		Limit:         2,
		Sold:          20,
		Refundable:    true,
		RefundPercent: 80,
	}
	event18Regular := models.Ticket{
		ID:            uuid.New(),
		EventID:       event18.ID,
		Name:          "ACOUSTIC SESSION PASS",
		Price:         150000,
		Quota:         200,
		Limit:         4,
		Sold:          80,
		Refundable:    true,
		RefundPercent: 70,
	}
	db.Create(&event18VIP)
	db.Create(&event18Regular)

	// Event 19 - Bandung Electronic Dance Music (EDM) Night
	event19 := models.Event{
		ID:          uuid.New(),
		Title:       "Bandung Electronic Dance Music Night",
		Image:       eventImages[1],
		Description: "The hottest EDM party in West Java! International and local DJs spinning the latest electronic dance music. Multiple rooms featuring different EDM sub-genres: progressive house, trance, dubstep, and techno. Professional lighting, sound systems, and visual effects. Dance until dawn in the cool mountain air of Bandung!",
		Location:    "Bandung",
		Date:        time.Now().AddDate(0, 0, 8),
		StartTime:   20,
		EndTime:     4,
		Status:      "active",
		CreatedAt:   time.Now().AddDate(0, 0, -1),
		UpdatedAt:   time.Now().AddDate(0, 0, -1),
	}
	db.Create(&event19)

	// Event 19 Tickets
	event19VIP := models.Ticket{
		ID:            uuid.New(),
		EventID:       event19.ID,
		Name:          "EDM VIP CLUB ACCESS",
		Price:         450000,
		Quota:         100,
		Limit:         2,
		Sold:          35,
		Refundable:    false,
		RefundPercent: 0,
	}
	event19Regular := models.Ticket{
		ID:            uuid.New(),
		EventID:       event19.ID,
		Name:          "DANCE FLOOR PASS",
		Price:         250000,
		Quota:         300,
		Limit:         4,
		Sold:          120,
		Refundable:    false,
		RefundPercent: 0,
	}
	db.Create(&event19VIP)
	db.Create(&event19Regular)
}
