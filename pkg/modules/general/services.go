package general

import (
	db "Authentication/pkg/config/database"
	"Authentication/pkg/entities"
	"time"
)

func (l LoginRequest) findWithEmail() (*entities.Users, error) {
	user := &entities.Users{}
	if err := db.PgDB.Where("email = ?", l.Email).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (l LoginRequest) saveRefreshToken(refreshToken string, expiredAt time.Time) error {
	tokenToSave := &entities.RefreshTokens{
		UserId:       l.ID,
		RefreshToken: refreshToken,
		ExpiredAt:    expiredAt,
	}
	if err := db.PgDB.Create(tokenToSave).Error; err != nil {
		return err
	}
	return nil
}
