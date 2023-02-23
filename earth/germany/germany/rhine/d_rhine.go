package rhine

import (
	"github.com/thalesfu/CK2Commands/earth/germany/germany/rhine/mainz"
	"github.com/thalesfu/CK2Commands/earth/germany/germany/rhine/pfalz"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RhineDuke interface {
	feud.Duke
	CMainz美因茨() mainz.MainzCounty
	CPfalz沃尔姆斯() pfalz.PfalzCounty
}

type 莱茵兰RhineDuke struct {
	feud.BaseDuke
	美因茨Mainz  mainz.MainzCounty
	沃尔姆斯Pfalz pfalz.PfalzCounty
}

func (d *莱茵兰RhineDuke) CMainz美因茨() mainz.MainzCounty {
	return d.美因茨Mainz
}

func (d *莱茵兰RhineDuke) CPfalz沃尔姆斯() pfalz.PfalzCounty {
	return d.沃尔姆斯Pfalz
}

var DRhine莱茵兰 RhineDuke = &莱茵兰RhineDuke{}

func init() {
	f := DRhine莱茵兰.(*莱茵兰RhineDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "rhine",
		TitleName: "莱茵兰",
		TitleCode: "d_rhine",
		Counties:  map[string]feud.County{},
	}

	f.美因茨Mainz = mainz.CMainz美因茨
	f.美因茨Mainz.SetParent(f)

	f.沃尔姆斯Pfalz = pfalz.CPfalz沃尔姆斯
	f.沃尔姆斯Pfalz.SetParent(f)

}
