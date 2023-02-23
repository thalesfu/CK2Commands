package kabul

import (
	"github.com/thalesfu/CK2Commands/earth/persia/afghanistan/kabul/bamiyan"
	"github.com/thalesfu/CK2Commands/earth/persia/afghanistan/kabul/kabul"
	"github.com/thalesfu/CK2Commands/earth/persia/afghanistan/kabul/kunduz"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KabulDuke interface {
	feud.Duke
	CBamiyan梵衍那() bamiyan.BamiyanCounty
	CKabul迦布罗() kabul.KabulCounty
	CKunduz昆都士() kunduz.KunduzCounty
}

type 迦布罗KabulDuke struct {
	feud.BaseDuke
	梵衍那Bamiyan bamiyan.BamiyanCounty
	迦布罗Kabul   kabul.KabulCounty
	昆都士Kunduz  kunduz.KunduzCounty
}

func (d *迦布罗KabulDuke) CBamiyan梵衍那() bamiyan.BamiyanCounty {
	return d.梵衍那Bamiyan
}

func (d *迦布罗KabulDuke) CKabul迦布罗() kabul.KabulCounty {
	return d.迦布罗Kabul
}

func (d *迦布罗KabulDuke) CKunduz昆都士() kunduz.KunduzCounty {
	return d.昆都士Kunduz
}

var DKabul迦布罗 KabulDuke = &迦布罗KabulDuke{}

func init() {
	f := DKabul迦布罗.(*迦布罗KabulDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "kabul",
		TitleName: "迦布罗",
		TitleCode: "d_kabul",
		Counties:  map[string]feud.County{},
	}

	f.梵衍那Bamiyan = bamiyan.CBamiyan梵衍那
	f.梵衍那Bamiyan.SetParent(f)

	f.迦布罗Kabul = kabul.CKabul迦布罗
	f.迦布罗Kabul.SetParent(f)

	f.昆都士Kunduz = kunduz.CKunduz昆都士
	f.昆都士Kunduz.SetParent(f)

}
