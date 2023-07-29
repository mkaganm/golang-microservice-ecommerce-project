package scripts

import (
	"cronitor/internal/data"
	"log"
	"time"
)

// DeleteExpiredVerifications is a function that deletes expired verifications
func DeleteExpiredVerifications() {
	db := data.InitPostgresDB()
	defer data.ClosePostgresDB(db)

	// Before two hours
	expirationTime := time.Now().Add(-2 * time.Hour)

	result := db.Exec("DELETE FROM verifications WHERE created_at < ?", expirationTime)
	if result.Error != nil {
		log.Default().Println("Error deleting expired verifications:", result.Error)
		return
	}

	log.Default().Println("Expired verifications deleted.")
}
