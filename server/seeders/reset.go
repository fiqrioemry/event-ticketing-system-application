package seeders

import (
	"log"
	"server/models"

	"gorm.io/gorm"
)

func ResetDatabase(db *gorm.DB) {
	log.Println("Dropping all tables...")

	err := db.Migrator().DropTable(
		&models.User{},
		&models.Event{},
		&models.Order{},
		&models.Ticket{},
		&models.OrderDetail{},
		&models.Payment{},
		&models.UserTicket{},
		&models.WithdrawalRequest{},
	)
	if err != nil {
		log.Fatalf("Failed to drop tables: %v", err)
	}

	log.Println("all tables dropped successfully.")

	log.Println("migrating tables...")

	err = db.AutoMigrate(
		&models.User{},
		&models.Event{},
		&models.Order{},
		&models.Ticket{},
		&models.OrderDetail{},
		&models.Payment{},
		&models.UserTicket{},
		&models.WithdrawalRequest{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate tables: %v", err)
	}

	log.Println("migration completed successfully.")

	log.Println("seeding dummy data...")

	SeedAll(db)
	SeedFirstData(db)
	log.Println("seeding completed successfully.")
}
