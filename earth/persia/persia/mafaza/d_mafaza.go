package mafaza

import (
	"github.com/thalesfu/CK2Commands/earth/persia/persia/mafaza/lut"
	"github.com/thalesfu/CK2Commands/earth/persia/persia/mafaza/yazd"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MafazaDuke interface {
	feud.Duke
	CLut玛法扎() lut.LutCounty
	CYazd亚兹德() yazd.YazdCounty
}

type 玛法扎MafazaDuke struct {
	feud.BaseDuke
	玛法扎Lut  lut.LutCounty
	亚兹德Yazd yazd.YazdCounty
}

func (d *玛法扎MafazaDuke) CLut玛法扎() lut.LutCounty {
	return d.玛法扎Lut
}

func (d *玛法扎MafazaDuke) CYazd亚兹德() yazd.YazdCounty {
	return d.亚兹德Yazd
}

var DMafaza玛法扎 MafazaDuke = &玛法扎MafazaDuke{}

func init() {
	f := DMafaza玛法扎.(*玛法扎MafazaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "mafaza",
		TitleName: "玛法扎",
		TitleCode: "d_mafaza",
		Counties:  map[string]feud.County{},
	}

	f.玛法扎Lut = lut.CLut玛法扎
	f.玛法扎Lut.SetParent(f)

	f.亚兹德Yazd = yazd.CYazd亚兹德
	f.亚兹德Yazd.SetParent(f)

}
