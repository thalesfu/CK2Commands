package raetia

import (
	"github.com/thalesfu/CK2Commands/earth/germany/germany/raetia/chur"
	"github.com/thalesfu/CK2Commands/earth/germany/germany/raetia/schwyz"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RaetiaDuke interface {
	feud.Duke
	CChur库尔() chur.ChurCounty
	CSchwyz施维茨() schwyz.SchwyzCounty
}

type 雷蒂亚RaetiaDuke struct {
	feud.BaseDuke
	库尔Chur    chur.ChurCounty
	施维茨Schwyz schwyz.SchwyzCounty
}

func (d *雷蒂亚RaetiaDuke) CChur库尔() chur.ChurCounty {
	return d.库尔Chur
}

func (d *雷蒂亚RaetiaDuke) CSchwyz施维茨() schwyz.SchwyzCounty {
	return d.施维茨Schwyz
}

var DRaetia雷蒂亚 RaetiaDuke = &雷蒂亚RaetiaDuke{}

func init() {
	f := DRaetia雷蒂亚.(*雷蒂亚RaetiaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "raetia",
		TitleName: "雷蒂亚",
		TitleCode: "d_raetia",
		Counties:  map[string]feud.County{},
	}

	f.库尔Chur = chur.CChur库尔
	f.库尔Chur.SetParent(f)

	f.施维茨Schwyz = schwyz.CSchwyz施维茨
	f.施维茨Schwyz.SetParent(f)

}
