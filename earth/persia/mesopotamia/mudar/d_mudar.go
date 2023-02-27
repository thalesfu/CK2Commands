package mudar

import (
	"github.com/thalesfu/CK2Commands/earth/persia/mesopotamia/mudar/al_bichri"
	"github.com/thalesfu/CK2Commands/earth/persia/mesopotamia/mudar/al_jazira"
	"github.com/thalesfu/CK2Commands/earth/persia/mesopotamia/mudar/rahbah"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MudarDuke interface {
    feud.Duke
    CAl_bichri比什里() 	al_bichri.Al_bichriCounty
    CAl_jazira哈塞克() 	al_jazira.Al_jaziraCounty
    CRahbah拉赫巴() 	rahbah.RahbahCounty
}

type 穆达尔MudarDuke struct {
	feud.BaseDuke
	比什里Al_bichri 	al_bichri.Al_bichriCounty
	哈塞克Al_jazira 	al_jazira.Al_jaziraCounty
	拉赫巴Rahbah 	rahbah.RahbahCounty
}

func (d *穆达尔MudarDuke) CAl_bichri比什里() al_bichri.Al_bichriCounty {
	return d.比什里Al_bichri
}
    
func (d *穆达尔MudarDuke) CAl_jazira哈塞克() al_jazira.Al_jaziraCounty {
	return d.哈塞克Al_jazira
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

	f.比什里Al_bichri = al_bichri.CAl_bichri比什里
	f.比什里Al_bichri.SetParent(f)
	
	f.哈塞克Al_jazira = al_jazira.CAl_jazira哈塞克
	f.哈塞克Al_jazira.SetParent(f)
	
	f.拉赫巴Rahbah = rahbah.CRahbah拉赫巴
	f.拉赫巴Rahbah.SetParent(f)
	
}
