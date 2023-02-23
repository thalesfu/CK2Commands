package sicily

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/sicily/sicily/agrigento"
	"github.com/thalesfu/CK2Commands/earth/byzantium/sicily/sicily/malta"
	"github.com/thalesfu/CK2Commands/earth/byzantium/sicily/sicily/messina"
	"github.com/thalesfu/CK2Commands/earth/byzantium/sicily/sicily/palermo"
	"github.com/thalesfu/CK2Commands/earth/byzantium/sicily/sicily/siracusa"
	"github.com/thalesfu/CK2Commands/earth/byzantium/sicily/sicily/trapani"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SicilyDuke interface {
	feud.Duke
	CAgrigento吉尔真蒂() agrigento.AgrigentoCounty
	CMalta马耳他() malta.MaltaCounty
	CMessina墨西拿() messina.MessinaCounty
	CPalermo巴勒莫() palermo.PalermoCounty
	CSiracusa锡拉库萨() siracusa.SiracusaCounty
	CTrapani特拉帕尼() trapani.TrapaniCounty
}

type 西西里SicilyDuke struct {
	feud.BaseDuke
	吉尔真蒂Agrigento agrigento.AgrigentoCounty
	马耳他Malta      malta.MaltaCounty
	墨西拿Messina    messina.MessinaCounty
	巴勒莫Palermo    palermo.PalermoCounty
	锡拉库萨Siracusa  siracusa.SiracusaCounty
	特拉帕尼Trapani   trapani.TrapaniCounty
}

func (d *西西里SicilyDuke) CAgrigento吉尔真蒂() agrigento.AgrigentoCounty {
	return d.吉尔真蒂Agrigento
}

func (d *西西里SicilyDuke) CMalta马耳他() malta.MaltaCounty {
	return d.马耳他Malta
}

func (d *西西里SicilyDuke) CMessina墨西拿() messina.MessinaCounty {
	return d.墨西拿Messina
}

func (d *西西里SicilyDuke) CPalermo巴勒莫() palermo.PalermoCounty {
	return d.巴勒莫Palermo
}

func (d *西西里SicilyDuke) CSiracusa锡拉库萨() siracusa.SiracusaCounty {
	return d.锡拉库萨Siracusa
}

func (d *西西里SicilyDuke) CTrapani特拉帕尼() trapani.TrapaniCounty {
	return d.特拉帕尼Trapani
}

var DSicily西西里 SicilyDuke = &西西里SicilyDuke{}

func init() {
	f := DSicily西西里.(*西西里SicilyDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "sicily",
		TitleName: "西西里",
		TitleCode: "d_sicily",
		Counties:  map[string]feud.County{},
	}

	f.吉尔真蒂Agrigento = agrigento.CAgrigento吉尔真蒂
	f.吉尔真蒂Agrigento.SetParent(f)

	f.马耳他Malta = malta.CMalta马耳他
	f.马耳他Malta.SetParent(f)

	f.墨西拿Messina = messina.CMessina墨西拿
	f.墨西拿Messina.SetParent(f)

	f.巴勒莫Palermo = palermo.CPalermo巴勒莫
	f.巴勒莫Palermo.SetParent(f)

	f.锡拉库萨Siracusa = siracusa.CSiracusa锡拉库萨
	f.锡拉库萨Siracusa.SetParent(f)

	f.特拉帕尼Trapani = trapani.CTrapani特拉帕尼
	f.特拉帕尼Trapani.SetParent(f)

}
