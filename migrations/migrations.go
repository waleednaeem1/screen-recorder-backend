package migrations

import (
	"screen-recorder-backend/config"
	"screen-recorder-backend/models"
	"gorm.io/gorm"
)

func AutoMigrateModels() error {
	// Auto migrate the models
	if err := config.DB.AutoMigrate(
		&models.User{},
		&models.ScreenRecording{},
		&models.NewsletterSubscription{},
		&models.FAQ{},
		&models.PrivacyPolicy{},
		&models.CustomerReview{},
		&models.Pricing{},
		&models.GeneralSettings{},
		&models.TeamMember{},
		&models.ProfileType{}, // Ensure ProfileType is included in migration
	); err != nil {
		return err
	}

	// Populate ProfileType with default values
	return populateProfileTypes(config.DB)
}

func populateProfileTypes(db *gorm.DB) error {
	profileTypes := []models.ProfileType{
		{
			RecordingsLeft:       10,
			VideoLengthAllocated: 5,
			Type:                 "basic",
		},
		{
			RecordingsLeft:       -1, // Use -1 to indicate unlimited
			VideoLengthAllocated: -1, // Use -1 to indicate unlimited
			Type:                 "business",
		},
	}

	for _, profileType := range profileTypes {
		// Check if the profile type already exists
		var existingProfileType models.ProfileType
		if err := db.Where("type = ?", profileType.Type).First(&existingProfileType).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// If it doesn't exist, create it
				if err := db.Create(&profileType).Error; err != nil {
					return err
				}
			} else {
				return err // Return any other error
			}
		}
	}

	return nil
}