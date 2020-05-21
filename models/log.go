package models

type Log struct {
	ID         string     `json:"id"`
	UID        int64      `json:"uid"`
	ReceiverID int64      `json:"receiver_id"`
	Act        int8       `json:"act"`
	EventID    int        `json:"event_id"`
	Exchanges  []Exchange `json:"exchanges"`
	TS         int64      `json:"ts"`
}

type Exchange struct {
	ExchangeType string `json:"exchange_type"`
	RelatedID    int    `json:"related_id"`
	Count        int    `json:"count"`
}
