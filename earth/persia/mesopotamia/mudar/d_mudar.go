package mudar

import (
	"github.com/thalesfu/CK2Commands/earth/persia/mesopotamia/mudar/rahbah"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MudarDuke interface {
	feud.Duke
	CRahbah拉赫巴() rahbah.RahbahCounty
}

type 穆达尔MudarDuke struct {
	feud.BaseDuke
	拉赫巴Rahbah rahbah.RahbahCounty
}

func (d *穆达尔MudarDuke) CRahbah拉赫巴() rahbah.RahbahCounty {
	return d.拉赫巴Rahbah
}

var DMudar穆达尔 MudarDuke = &穆达尔MudarDuke{}

func init() {
	f := DMudar穆达尔.(*穆达尔MudarDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "mudar",
		TitleName: "穆达尔",
		TitleCode: "d_mudar",
		Counties:  map[string]feud.County{},
	}

	f.拉赫巴Rahbah = rahbah.CRahbah拉赫巴
	f.拉赫巴Rahbah.SetParent(f)

}
