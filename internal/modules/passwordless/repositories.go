package passwordless

import (
	db "Authentication/internal/config/database"
	"Authentication/internal/entities"
	"os"
)

func (request PublicProfile) findWithID() (PublicProfile, error) {
	user := &entities.Users{}
	if err := db.PgDB.Where("id = ?", request.ID).First(user).Error; err != nil {
		return PublicProfile{}, err
	}
	return PublicProfile{
		ID:          user.ID,
		Email:       user.Email,
		FullName:    user.FullName,
		DisplayName: user.DisplayName,
		Picture:     user.Picture,
	}, nil
}

func (request passwordlessRegistration) getPasswordlessRegistrationData() (passwordlessRegistration, error) {
	sqlFile, err := os.ReadFile("internal/sql/queries/passwordless/get_registration_data.sql")
	if err != nil {
		return passwordlessRegistration{}, err
	}
	sql := string(sqlFile)
	passwordless := &passwordlessRegistration{}
	if err := db.PgDB.Raw(sql, request.UserId).Scan(passwordless).Error; err != nil {
		return passwordlessRegistration{}, err
	}
	return *passwordless, nil
}
