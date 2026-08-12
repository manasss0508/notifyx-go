package rbmq

import "github.com/google/uuid"

type NotifMsg struct {
	NotficationId uuid.UUID `json:"notfication_id"`
}
