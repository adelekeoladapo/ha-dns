package impl

import (
	"ha-dns/dto"
	"ha-dns/service"
)

const SECTOR_ID = 001

type DnsServiceImpl struct {
}

func (o DnsServiceImpl) CalculateLocation(x, y, z, vel float64) (response dto.DnsResponse, e error) {
	response.Loc = (x * SECTOR_ID) + (y * SECTOR_ID) + (z * SECTOR_ID) + vel
	return
}

func GetDnsServiceImpl() service.DnsService {
	return &DnsServiceImpl{}
}
