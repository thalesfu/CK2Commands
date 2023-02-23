package mallorca

import (
	"github.com/thalesfu/CK2Commands/earth/spain/aragon/mallorca/mallorca"
	"github.com/thalesfu/CK2Commands/earth/spain/aragon/mallorca/menorca"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MallorcaDuke interface {
	feud.Duke
	CMallorca马略卡() mallorca.MallorcaCounty
	CMenorca梅诺卡() menorca.MenorcaCounty
}

type 马略卡MallorcaDuke struct {
	feud.BaseDuke
	马略卡Mallorca mallorca.MallorcaCounty
	梅诺卡Menorca  menorca.MenorcaCounty
}

func (d *马略卡MallorcaDuke) CMallorca马略卡() mallorca.MallorcaCounty {
	return d.马略卡Mallorca
}

func (d *马略卡MallorcaDuke) CMenorca梅诺卡() menorca.MenorcaCounty {
	return d.梅诺卡Menorca
}

var DMallorca马略卡 MallorcaDuke = &马略卡MallorcaDuke{}

func init() {
	f := DMallorca马略卡.(*马略卡MallorcaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "mallorca",
		TitleName: "马略卡",
		TitleCode: "d_mallorca",
		Counties:  map[string]feud.County{},
	}

	f.马略卡Mallorca = mallorca.CMallorca马略卡
	f.马略卡Mallorca.SetParent(f)

	f.梅诺卡Menorca = menorca.CMenorca梅诺卡
	f.梅诺卡Menorca.SetParent(f)

}
