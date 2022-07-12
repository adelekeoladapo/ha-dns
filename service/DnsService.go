package service

import "ha-dns/dto"

type DnsService interface {
	CalculateLocation(x, y, z, vel float64) (dto.DnsResponse, error)
}
