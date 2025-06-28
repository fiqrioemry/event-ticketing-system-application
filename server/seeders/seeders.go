package seeders

import (
	"log"
	"strconv"
	"time"

	"server/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedAll(db *gorm.DB) {
	log.Println("Seeding started...")
	password, _ := bcrypt.GenerateFromPassword([]byte("123456"), 10)

	// --- Users ---
	adminID := uuid.New()
	customer1ID := uuid.New()
	customer2ID := uuid.New()

	admin := models.User{
		ID: adminID, Fullname: "Event Admin", Email: "admin@event.com",
		Password: string(password), Role: "admin",
		AvatarURL: "https://api.dicebear.com/6.x/initials/svg?seed=admin",
	}
	customer1 := models.User{
		ID: customer1ID, Fullname: "Alice Customer", Email: "alice@event.com",
		Password: string(password), Role: "user",
		AvatarURL: "https://api.dicebear.com/6.x/initials/svg?seed=alice",
	}
	customer2 := models.User{
		ID: customer2ID, Fullname: "Bob Customer", Email: "bob@event.com",
		Password: string(password), Role: "user",
		AvatarURL: "https://api.dicebear.com/6.x/initials/svg?seed=bob",
	}

	db.Create(&admin)
	db.Create(&customer1)
	db.Create(&customer2)

	// --- Events & Tickets ---
	var ticketIDs []uuid.UUID
	var eventIDs []uuid.UUID
	for i := 1; i <= 2; i++ {
		eventID := uuid.New()
		eventIDs = append(eventIDs, eventID)
		event := models.Event{
			ID:          eventID,
			Title:       "Event " + strconv.Itoa(i),
			Image:       "https://placehold.co/400x400?text=Konser+Musik+Senja",
			Description: "Description for event",
			Location:    "Jakarta",
			Date:        time.Now(),
			StartTime:   9,
			EndTime:     17,
			Status:      "active",
		}
		db.Create(&event)

		for _, name := range []string{"VVIP", "VIP", "Regular"} {
			ticketID := uuid.New()
			ticketIDs = append(ticketIDs, ticketID)
			ticket := models.Ticket{
				ID:            ticketID,
				EventID:       eventID,
				Name:          name,
				Price:         100000,
				Quota:         100,
				Limit:         2,
				Sold:          10,
				Refundable:    true,
				RefundPercent: 50,
			}
			db.Create(&ticket)
		}
	}

	// --- Orders, OrderDetails, Payments, UserTickets ---
	paidTime := time.Now().AddDate(0, 0, -2)
	refundedTime := time.Now().AddDate(0, 0, -1)

	orderUsers := []uuid.UUID{customer1ID, customer2ID}
	for i, userID := range orderUsers {
		orderID := uuid.New()
		eventID := eventIDs[i]
		ticket1 := ticketIDs[i*3]
		ticket2 := ticketIDs[i*3+1]

		detail1 := models.OrderDetail{
			ID: uuid.New(), OrderID: orderID,
			TicketID: ticket1, TicketName: "VIP", Quantity: 1, Price: 150000,
		}
		detail2 := models.OrderDetail{
			ID: uuid.New(), OrderID: orderID,
			TicketID: ticket2, TicketName: "Regular", Quantity: 1, Price: 100000,
		}

		total := detail1.Price + detail2.Price
		order := models.Order{
			ID: orderID, UserID: userID, EventID: eventID,
			Fullname: "Customer", Email: "customer@example.com", Phone: "08123456789",
			TotalPrice: total, Status: "paid",
		}

		if i == 0 {
			order.IsRefunded = true
			order.RefundedAt = &refundedTime
			order.RefundAmount = total * 0.5
			order.RefundReason = "Permintaan user"
		}

		db.Create(&order)
		db.Create(&detail1)
		db.Create(&detail2)

		payment := models.Payment{
			ID: uuid.New(), UserID: userID, OrderID: orderID,
			Fullname: "Customer", Email: "customer@example.com",
			Method: "stripe", Amount: total, Status: "paid", PaidAt: &paidTime,
		}
		db.Create(&payment)

		db.Create(&models.UserTicket{
			ID: uuid.New(), UserID: userID, EventID: eventID,
			TicketID: ticket1, QRCode: "QR-1",
		})
		db.Create(&models.UserTicket{
			ID: uuid.New(), UserID: userID, EventID: eventID,
			TicketID: ticket2, QRCode: "QR-2",
		})
	}

	// --- Withdrawal ---
	db.Create(&models.WithdrawalRequest{
		ID: uuid.New(), UserID: customer1ID,
		Amount: 125000, Status: "approved",
		Reason: "Refund event", CreatedAt: refundedTime,
		ApprovedAt: &paidTime, ReviewedBy: &adminID,
	})

	log.Println("Seeding completed.")
}

func SeedFirstData(db *gorm.DB) {
	log.Println("Seeding started...")
	paidTime := time.Now().AddDate(0, 0, -2)
	password, _ := bcrypt.GenerateFromPassword([]byte("123456"), 10)

	// --- Additional Customer and Purchase ---
	newCustomer := models.User{
		ID:        uuid.MustParse("bdb598a3-1c86-4e95-93d2-65cda21b4b33"),
		Fullname:  "Charlie Customer",
		Email:     "charlie@event.com",
		Password:  string(password),
		Role:      "user",
		AvatarURL: "https://api.dicebear.com/6.x/initials/svg?seed=charlie",
	}
	db.Create(&newCustomer)

	newEvent := models.Event{
		ID:          uuid.MustParse("ddaf8eb0-a68e-4316-8dc7-834d183faaf6"),
		Title:       "Event 3",
		Image:       "https://placehold.co/400x400?text=Acoustic+Night",
		Description: "Description for event 3",
		Location:    "Bandung",
		Date:        time.Now().AddDate(0, 1, 0),
		StartTime:   18,
		EndTime:     22,
		Status:      "active",
	}
	db.Create(&newEvent)

	ticketA := models.Ticket{
		ID:            uuid.MustParse("d39ff313-db85-4a3f-9f33-fa8fb0c68019"),
		EventID:       uuid.MustParse("ddaf8eb0-a68e-4316-8dc7-834d183faaf6"),
		Name:          "Gold",
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
		Name:          "Silver",
		Price:         90000,
		Quota:         50,
		Limit:         2,
		Sold:          3,
		Refundable:    true,
		RefundPercent: 70,
	}
	db.Create(&ticketA)
	db.Create(&ticketB)

	newOrder := models.Order{
		ID:         uuid.MustParse("24525c25-1944-4d57-8bbd-c4169cc46f04"),
		UserID:     uuid.MustParse("bdb598a3-1c86-4e95-93d2-65cda21b4b33"),
		EventID:    uuid.MustParse("ddaf8eb0-a68e-4316-8dc7-834d183faaf6"),
		Fullname:   "Charlie Customer",
		Email:      "charlie@event.com",
		Phone:      "08123456788",
		TotalPrice: 210000,
		Status:     "paid",
	}
	db.Create(&newOrder)

	db.Create(&models.OrderDetail{
		ID:         uuid.MustParse("798186a9-2a2e-42cd-bcd4-8ea98f47dbf5"),
		OrderID:    uuid.MustParse("24525c25-1944-4d57-8bbd-c4169cc46f04"),
		TicketID:   uuid.MustParse("d39ff313-db85-4a3f-9f33-fa8fb0c68019"),
		TicketName: "Gold",
		Quantity:   1,
		Price:      120000,
	})
	db.Create(&models.OrderDetail{
		ID:         uuid.MustParse("5c6a27fe-82c2-4135-be8b-cfc40d1b57d4"),
		OrderID:    uuid.MustParse("24525c25-1944-4d57-8bbd-c4169cc46f04"),
		TicketID:   uuid.MustParse("802a2b63-d941-4557-8dc7-0d0580d0580f"),
		TicketName: "Silver",
		Quantity:   1,
		Price:      90000,
	})

	db.Create(&models.Payment{
		ID:       uuid.MustParse("2ed5aec9-af02-4573-8fe0-c0f376af9bc1"),
		UserID:   uuid.MustParse("bdb598a3-1c86-4e95-93d2-65cda21b4b33"),
		OrderID:  uuid.MustParse("24525c25-1944-4d57-8bbd-c4169cc46f04"),
		Fullname: "Charlie Customer",
		Email:    "charlie@event.com",
		Method:   "stripe",
		Amount:   210000,
		Status:   "paid",
		PaidAt:   &paidTime,
	})

	db.Create(&models.UserTicket{
		ID:       uuid.MustParse("7e3d5fe9-6915-46d2-a117-7eb5d645e3f6"),
		UserID:   uuid.MustParse("bdb598a3-1c86-4e95-93d2-65cda21b4b33"),
		EventID:  uuid.MustParse("ddaf8eb0-a68e-4316-8dc7-834d183faaf6"),
		TicketID: uuid.MustParse("d39ff313-db85-4a3f-9f33-fa8fb0c68019"),
		QRCode:   "QR-3",
	})
	db.Create(&models.UserTicket{
		ID:       uuid.MustParse("04740e8d-03dd-41d6-ba3f-e1031e59cc88"),
		UserID:   uuid.MustParse("bdb598a3-1c86-4e95-93d2-65cda21b4b33"),
		EventID:  uuid.MustParse("ddaf8eb0-a68e-4316-8dc7-834d183faaf6"),
		TicketID: uuid.MustParse("802a2b63-d941-4557-8dc7-0d0580d0580f"),
		QRCode:   "QR-4",
	})
}
