package server

// import (
// 	"context"
// 	"crypto/rand"
// 	"encoding/hex"
// 	"time"

// 	"github.com/emaforlin/bussiness-service/entities"
// )

// const tokenBytesLength = 16

// // SendJoinInvitation implements Service.
// func (s *service) SendJoinInvitation(ctx context.Context, invitation entities.InvitationDto) (string, error) {
// 	token, _ := generateToken(tokenBytesLength)

// 	err := s.db.Create(&entities.Invitation{
// 		InviterID:      invitation.InviterID,
// 		Token:          token,
// 		RecipientEmail: invitation.RecipientEmail,
// 		AssignedRole:   invitation.AssignedRole,
// 		Expiration:     time.Now().Add(3 * 24 * time.Hour).Unix(),
// 	}).Error

// 	if err != nil {
// 		return "", err
// 	}

// 	s.logger.Debug("Invitation sent")
// 	return token, nil
// }

// func generateToken(length int) (string, error) {
// 	bytes := make([]byte, length)
// 	if _, err := rand.Read(bytes); err != nil {
// 		return "", err
// 	}
// 	return hex.EncodeToString(bytes), nil
// }
