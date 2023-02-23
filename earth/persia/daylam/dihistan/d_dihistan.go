package dihistan

import (
	"github.com/thalesfu/CK2Commands/earth/persia/daylam/dihistan/dihistan"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DihistanDuke interface {
	feud.Duke
	CDihistan大益斯坦() dihistan.DihistanCounty
}

type 大益斯坦DihistanDuke struct {
	feud.BaseDuke
	大益斯坦Dihistan dihistan.DihistanCounty
}

func (d *大益斯坦DihistanDuke) CDihistan大益斯坦() dihistan.DihistanCounty {
	return d.大益斯坦Dihistan
}

var DDihistan大益斯坦 DihistanDuke = &大益斯坦DihistanDuke{}

func init() {
	f := DDihistan大益斯坦.(*大益斯坦DihistanDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "dihistan",
		TitleName: "大益斯坦",
		TitleCode: "d_dihistan",
		Counties:  map[string]feud.County{},
	}

	f.大益斯坦Dihistan = dihistan.CDihistan大益斯坦
	f.大益斯坦Dihistan.SetParent(f)

}
