package main

import "errors"

type Request struct {
	Text   string `json:"text"`
	QrText string `json:"qr_text"`
}

func (r Request) Validate() error {
	if r.Text == "" || r.QrText == "" {
		return errors.New("both `text` and `qr_text` are required")
	}

	return nil
}

type Response struct {
	Status string `json:"status"`
}
