package syrmia

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/croatia/syrmia/syrmia"
	"github.com/thalesfu/CK2Commands/earth/byzantium/croatia/syrmia/vukovar"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SyrmiaDuke interface {
	feud.Duke
	CSyrmia叙尔米亚() syrmia.SyrmiaCounty
	CVukovar武科瓦尔() vukovar.VukovarCounty
}

type 叙尔米亚SyrmiaDuke struct {
	feud.BaseDuke
	叙尔米亚Syrmia  syrmia.SyrmiaCounty
	武科瓦尔Vukovar vukovar.VukovarCounty
}

func (d *叙尔米亚SyrmiaDuke) CSyrmia叙尔米亚() syrmia.SyrmiaCounty {
	return d.叙尔米亚Syrmia
}

func (d *叙尔米亚SyrmiaDuke) CVukovar武科瓦尔() vukovar.VukovarCounty {
	return d.武科瓦尔Vukovar
}

var DSyrmia叙尔米亚 SyrmiaDuke = &叙尔米亚SyrmiaDuke{}

func init() {
	f := DSyrmia叙尔米亚.(*叙尔米亚SyrmiaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "syrmia",
		TitleName: "叙尔米亚",
		TitleCode: "d_syrmia",
		Counties:  map[string]feud.County{},
	}

	f.叙尔米亚Syrmia = syrmia.CSyrmia叙尔米亚
	f.叙尔米亚Syrmia.SetParent(f)

	f.武科瓦尔Vukovar = vukovar.CVukovar武科瓦尔
	f.武科瓦尔Vukovar.SetParent(f)

}
