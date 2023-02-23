package verona

import (
	"github.com/thalesfu/CK2Commands/earth/italy/italy/verona/mantua"
	"github.com/thalesfu/CK2Commands/earth/italy/italy/verona/padova"
	"github.com/thalesfu/CK2Commands/earth/italy/italy/verona/verona"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VeronaDuke interface {
	feud.Duke
	CMantua曼图阿() mantua.MantuaCounty
	CPadova帕多维亚() padova.PadovaCounty
	CVerona维罗纳() verona.VeronaCounty
}

type 维罗纳VeronaDuke struct {
	feud.BaseDuke
	曼图阿Mantua  mantua.MantuaCounty
	帕多维亚Padova padova.PadovaCounty
	维罗纳Verona  verona.VeronaCounty
}

func (d *维罗纳VeronaDuke) CMantua曼图阿() mantua.MantuaCounty {
	return d.曼图阿Mantua
}

func (d *维罗纳VeronaDuke) CPadova帕多维亚() padova.PadovaCounty {
	return d.帕多维亚Padova
}

func (d *维罗纳VeronaDuke) CVerona维罗纳() verona.VeronaCounty {
	return d.维罗纳Verona
}

var DVerona维罗纳 VeronaDuke = &维罗纳VeronaDuke{}

func init() {
	f := DVerona维罗纳.(*维罗纳VeronaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "verona",
		TitleName: "维罗纳",
		TitleCode: "d_verona",
		Counties:  map[string]feud.County{},
	}

	f.曼图阿Mantua = mantua.CMantua曼图阿
	f.曼图阿Mantua.SetParent(f)

	f.帕多维亚Padova = padova.CPadova帕多维亚
	f.帕多维亚Padova.SetParent(f)

	f.维罗纳Verona = verona.CVerona维罗纳
	f.维罗纳Verona.SetParent(f)

}
