package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"ha-dns/dto"
	"ha-dns/service"
	"ha-dns/service/impl"
	"ha-dns/utils"
	"ha-dns/utils/message"
	"ha-dns/utils/responseStatus"
	"net/http"
	"strconv"
)

var dnsService service.DnsService = impl.GetDnsServiceImpl()

func GetLocation(c echo.Context) error {
	request := new(dto.DnsRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ServiceResponse{Status: responseStatus.ERROR, Message: err.Error()})
	}
	if err := utils.Validate(request); len(err) != 0 {
		return c.JSON(http.StatusBadRequest, dto.ServiceResponse{Status: responseStatus.ERROR, Message: fmt.Sprint(err)})
	}
	var x, y, z, vel float64
	var err error
	if x, err = strconv.ParseFloat(request.X, 64); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ServiceResponse{Status: responseStatus.ERROR, Message: fmt.Sprint(err)})
	}
	if y, err = strconv.ParseFloat(request.Y, 64); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ServiceResponse{Status: responseStatus.ERROR, Message: fmt.Sprint(err)})
	}
	if z, err = strconv.ParseFloat(request.Z, 64); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ServiceResponse{Status: responseStatus.ERROR, Message: fmt.Sprint(err)})
	}
	if vel, err = strconv.ParseFloat(request.Vel, 64); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ServiceResponse{Status: responseStatus.ERROR, Message: fmt.Sprint(err)})
	}
	response, err := dnsService.CalculateLocation(x, y, z, vel)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ServiceResponse{Status: responseStatus.ERROR, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.ServiceResponse{Status: responseStatus.SUCCESS, Message: message.GENERAL_SUCCESS_MESSAGE, Data: response})
}
