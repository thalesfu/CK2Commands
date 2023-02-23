package air

import (
	"github.com/thalesfu/CK2Commands/earth/kanem/hausaland/air/air"
	"github.com/thalesfu/CK2Commands/earth/kanem/hausaland/air/fachi"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AirDuke interface {
	feud.Duke
	CAir艾尔() air.AirCounty
	CFachi法希() fachi.FachiCounty
}

type 艾尔AirDuke struct {
	feud.BaseDuke
	艾尔Air   air.AirCounty
	法希Fachi fachi.FachiCounty
}

func (d *艾尔AirDuke) CAir艾尔() air.AirCounty {
	return d.艾尔Air
}

func (d *艾尔AirDuke) CFachi法希() fachi.FachiCounty {
	return d.法希Fachi
}

var DAir艾尔 AirDuke = &艾尔AirDuke{}

func init() {
	f := DAir艾尔.(*艾尔AirDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "air",
		TitleName: "艾尔",
		TitleCode: "d_air",
		Counties:  map[string]feud.County{},
	}

	f.艾尔Air = air.CAir艾尔
	f.艾尔Air.SetParent(f)

	f.法希Fachi = fachi.CFachi法希
	f.法希Fachi.SetParent(f)

}
