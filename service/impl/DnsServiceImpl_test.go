package impl

import (
	"ha-dns/dto"
	"reflect"
	"testing"
)

func TestDnsServiceImpl_CalculateLocation(t *testing.T) {
	type args struct {
		x   float64
		y   float64
		z   float64
		vel float64
	}
	tests := []struct {
		name         string
		args         args
		wantResponse dto.DnsResponse
		wantErr      bool
	}{
		{
			"test01",
			args{10, 7, 2, 5},
			dto.DnsResponse{Loc: 24},
			false,
		},
		{
			"test02",
			args{15.6, 1.2, 6.098, 51},
			dto.DnsResponse{Loc: 73.898},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := DnsServiceImpl{}
			gotResponse, err := o.CalculateLocation(tt.args.x, tt.args.y, tt.args.z, tt.args.vel)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("CalculateLocation() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}
