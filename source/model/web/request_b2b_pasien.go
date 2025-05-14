package web

type RequestB2BPasien struct {
	Id                   uint64   `json:"id" example:"1" default:"0"`
	PasienName string `json:"pasien_name" validate:"required"`
}
