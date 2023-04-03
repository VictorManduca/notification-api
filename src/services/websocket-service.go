package services

import (
	"encoding/json"

	"notification-api/src/models"

	"github.com/olahol/melody"
)

func EmitUserCreatedEvent(user models.User, m *melody.Melody) error {
	u, err := json.Marshal(user)
	if err != nil {
		return err
	}

	m.Broadcast(u)

	return nil
}
