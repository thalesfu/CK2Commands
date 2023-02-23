package asturias

import (
	"github.com/thalesfu/CK2Commands/earth/spain/asturias/asturias/astorga"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AsturiasDuke interface {
	feud.Duke
	CAstorga阿斯托加() astorga.AstorgaCounty
}

type 阿斯图里亚斯AsturiasDuke struct {
	feud.BaseDuke
	阿斯托加Astorga astorga.AstorgaCounty
}

func (d *阿斯图里亚斯AsturiasDuke) CAstorga阿斯托加() astorga.AstorgaCounty {
	return d.阿斯托加Astorga
}

var DAsturias阿斯图里亚斯 AsturiasDuke = &阿斯图里亚斯AsturiasDuke{}

func init() {
	f := DAsturias阿斯图里亚斯.(*阿斯图里亚斯AsturiasDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "asturias",
		TitleName: "阿斯图里亚斯",
		TitleCode: "d_asturias",
		Counties:  map[string]feud.County{},
	}

	f.阿斯托加Astorga = astorga.CAstorga阿斯托加
	f.阿斯托加Astorga.SetParent(f)

}
