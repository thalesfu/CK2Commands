package carniola

import (
	"github.com/thalesfu/CK2Commands/earth/germany/carinthia/carniola/istria"
	"github.com/thalesfu/CK2Commands/earth/germany/carinthia/carniola/krain"
	"github.com/thalesfu/CK2Commands/earth/germany/carinthia/carniola/pettau"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CarniolaDuke interface {
	feud.Duke
	CIstria伊斯特里亚() istria.IstriaCounty
	CKrain克雷纳() krain.KrainCounty
	CPettau佩陶() pettau.PettauCounty
}

type 卡尔尼奥拉CarniolaDuke struct {
	feud.BaseDuke
	伊斯特里亚Istria istria.IstriaCounty
	克雷纳Krain    krain.KrainCounty
	佩陶Pettau    pettau.PettauCounty
}

func (d *卡尔尼奥拉CarniolaDuke) CIstria伊斯特里亚() istria.IstriaCounty {
	return d.伊斯特里亚Istria
}

func (d *卡尔尼奥拉CarniolaDuke) CKrain克雷纳() krain.KrainCounty {
	return d.克雷纳Krain
}

func (d *卡尔尼奥拉CarniolaDuke) CPettau佩陶() pettau.PettauCounty {
	return d.佩陶Pettau
}

var DCarniola卡尔尼奥拉 CarniolaDuke = &卡尔尼奥拉CarniolaDuke{}

func init() {
	f := DCarniola卡尔尼奥拉.(*卡尔尼奥拉CarniolaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "carniola",
		TitleName: "卡尔尼奥拉",
		TitleCode: "d_carniola",
		Counties:  map[string]feud.County{},
	}

	f.伊斯特里亚Istria = istria.CIstria伊斯特里亚
	f.伊斯特里亚Istria.SetParent(f)

	f.克雷纳Krain = krain.CKrain克雷纳
	f.克雷纳Krain.SetParent(f)

	f.佩陶Pettau = pettau.CPettau佩陶
	f.佩陶Pettau.SetParent(f)

}
