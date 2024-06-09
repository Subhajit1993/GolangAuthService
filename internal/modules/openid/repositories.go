package openid

import (
	db "Authentication/internal/config/database"
	"Authentication/internal/entities"
	"time"
)

func (p PublicProfile) findWithID() (PublicProfile, error) {
	user := &entities.Users{}
	if err := db.PgDB.Where("id = ?", p.ID).First(user).Error; err != nil {
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

func (p PublicProfile) findWithEmail() (PublicProfile, error) {
	user := &entities.Users{}
	if err := db.PgDB.Where("email = ?", p.Email).First(user).Error; err != nil {
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

func (p PublicProfile) saveData() (PublicProfile, error) {

	user := &entities.Users{
		Email:              p.Email,
		FullName:           p.FullName,
		DisplayName:        p.DisplayName,
		Active:             true,
		Verified:           p.Verified,
		RegistrationSource: p.RegistrationSource,
		Picture:            p.Picture,
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

func (p PublicProfile) saveRefreshToken(refreshToken string, expiredAt time.Time) error {
	tokenToSave := &entities.RefreshTokens{
		UserId:       p.ID,
		RefreshToken: refreshToken,
		ExpiredAt:    expiredAt,
	}
	if err := db.PgDB.Create(tokenToSave).Error; err != nil {
		return err
	}
	return nil
}
