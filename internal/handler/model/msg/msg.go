package msg

import "time"

type MSG struct {
	Type string `json:"type"`

	Data []byte `json:"data"`

	TTL int `json:"ttl"`
}

type Notify struct {
	Data string `json:"data"`
	TTL  int    `json:"ttl"`
}

type JJJ struct {
	Id        string        `json:"id"`
	UserId    int           `json:"user_id"`
	Status    string        `json:"status"`
	Content   string        `json:"content"`
	TTL       time.Duration `json:"ttl"`
	CreatedAt time.Time     `json:"created_at"`
	ExpiredAt time.Time     `json:"expired_at"`
}

type STRUCT struct {
	Type string `json:"type"`
	Data []JJJ  `json:"data"`
}
