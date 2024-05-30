package openid

import (
	db "Authentication/internal/config/database"
	"Authentication/internal/entities"
)

func (r PublicProfile) findWithID() (PublicProfile, error) {
	user := &entities.Users{}
	if err := db.PgDB.Where("id = ?", r.ID).First(user).Error; err != nil {
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

func (r PublicProfile) findWithEmail() (PublicProfile, error) {
	user := &entities.Users{}
	if err := db.PgDB.Where("email = ?", r.Email).First(user).Error; err != nil {
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

func (r PublicProfile) saveData() (PublicProfile, error) {

	user := &entities.Users{
		Email:              r.Email,
		FullName:           r.FullName,
		DisplayName:        r.DisplayName,
		Active:             true,
		Verified:           r.Verified,
		RegistrationSource: r.RegistrationSource,
		Picture:            r.Picture,
	}

	if err := db.PgDB.Create(user).Error; err != nil {
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
