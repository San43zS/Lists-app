package notification

type Notification struct {
	Id         int    `json:"-"`
	Info       string `json:"info"`
	TTL        int    `json:"ttl"`
	TimeCreate int    `json:"timeCreate"`
}
