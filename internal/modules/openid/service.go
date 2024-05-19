package openid

import (
	db "Authentication/internal/config/database"
	"Authentication/internal/entities"
)

func (r Profile) save() (*entities.Users, error) {

	user := &entities.Users{
		Email:              r.Email,
		FullName:           r.FullName,
		DisplayName:        r.DisplayName,
		Active:             true,
		Verified:           r.Verified,
		RegistrationSource: r.RegistrationSource,
	}

	if err := db.PgDB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r Profile) find() (*entities.Users, error) {
	user := &entities.Users{}
	if err := db.PgDB.Where("email = ?", r.Email).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
