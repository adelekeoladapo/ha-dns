package dto

type DnsRequest struct {
	X   string `validate:"required"`
	Y   string `validate:"required"`
	Z   string `validate:"required"`
	Vel string `validate:"required"`
}
