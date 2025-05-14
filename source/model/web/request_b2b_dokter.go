package web

type RequestB2BDokter struct {
    DokterName string `json:"dokter_name" validate:"required"`
}
