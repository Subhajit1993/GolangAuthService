package passwordless

import (
	db "Authentication/internal/config/database"
	"Authentication/internal/entities"
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
