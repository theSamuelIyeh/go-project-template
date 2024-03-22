package services

// var DB *gorm.DB
// var err error

// func InitDB() (*gorm.DB, error) {
// 	dsn := os.Getenv("SUPABASE_DB_URL")
// 	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	fmt.Println("db connected successfully")
// 	return DB, nil
// }

// func AutoMigrate(db *gorm.DB) error {
// 	err := db.AutoMigrate(&models.User{}, &models.Exam{}, &models.Question{}, &models.ExamAttempt{}, &models.ExamResponse{})
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func DeleteTables(db *gorm.DB) error {
// 	err := db.Migrator().DropTable(&models.User{}, &models.Exam{}, &models.Question{}, &models.ExamAttempt{}, &models.ExamResponse{})
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
