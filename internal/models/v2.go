package models

type IgnoreUser struct {
	User_id    string `json:"user_id"`
	Is_ignored bool   `json:"is_ignored"`
}

type IgnoreTopic struct {
	Topic_id   string `json:"topic_id"`
	Is_ignored bool   `json:"is_ignored"`
}
