package models

type QRResult struct {
	Ticket        string `json:"ticket"`
	ExpireSeconds int    `json:"expire_seconds "`
	URL           string `json:"url "`
}
