package service

import (
	"encoding/json"

	"github.com/manasss0508/notifyx-go/internals/datamodels/rbmq"
)

func DeserializeMessage(msgBytes []byte) (rbmq.NotifMsg, error) {
	var msg rbmq.NotifMsg
	if err := json.Unmarshal(msgBytes, &msg); err != nil {
		return rbmq.NotifMsg{}, err
	}

	return msg, nil
}
