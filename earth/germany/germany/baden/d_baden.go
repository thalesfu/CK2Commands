package baden

import (
	"github.com/thalesfu/CK2Commands/earth/germany/germany/baden/baden"
	"github.com/thalesfu/CK2Commands/earth/germany/germany/baden/breisgau"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BadenDuke interface {
	feud.Duke
	CBaden巴登() baden.BadenCounty
	CBreisgau布赖施高() breisgau.BreisgauCounty
}

type 巴登BadenDuke struct {
	feud.BaseDuke
	巴登Baden      baden.BadenCounty
	布赖施高Breisgau breisgau.BreisgauCounty
}

func (d *巴登BadenDuke) CBaden巴登() baden.BadenCounty {
	return d.巴登Baden
}

func (d *巴登BadenDuke) CBreisgau布赖施高() breisgau.BreisgauCounty {
	return d.布赖施高Breisgau
}

var DBaden巴登 BadenDuke = &巴登BadenDuke{}

func init() {
	f := DBaden巴登.(*巴登BadenDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "baden",
		TitleName: "巴登",
		TitleCode: "d_baden",
		Counties:  map[string]feud.County{},
	}

	f.巴登Baden = baden.CBaden巴登
	f.巴登Baden.SetParent(f)

	f.布赖施高Breisgau = breisgau.CBreisgau布赖施高
	f.布赖施高Breisgau.SetParent(f)

}
