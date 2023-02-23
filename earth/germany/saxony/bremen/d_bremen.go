package bremen

import (
	"github.com/thalesfu/CK2Commands/earth/germany/saxony/bremen/bremen"
	"github.com/thalesfu/CK2Commands/earth/germany/saxony/bremen/celle"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BremenDuke interface {
	feud.Duke
	CBremen哈德尔恩() bremen.BremenCounty
	CCelle不来梅() celle.CelleCounty
}

type 不来梅BremenDuke struct {
	feud.BaseDuke
	哈德尔恩Bremen bremen.BremenCounty
	不来梅Celle   celle.CelleCounty
}

func (d *不来梅BremenDuke) CBremen哈德尔恩() bremen.BremenCounty {
	return d.哈德尔恩Bremen
}

func (d *不来梅BremenDuke) CCelle不来梅() celle.CelleCounty {
	return d.不来梅Celle
}

var DBremen不来梅 BremenDuke = &不来梅BremenDuke{}

func init() {
	f := DBremen不来梅.(*不来梅BremenDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "bremen",
		TitleName: "不来梅",
		TitleCode: "d_bremen",
		Counties:  map[string]feud.County{},
	}

	f.哈德尔恩Bremen = bremen.CBremen哈德尔恩
	f.哈德尔恩Bremen.SetParent(f)

	f.不来梅Celle = celle.CCelle不来梅
	f.不来梅Celle.SetParent(f)

}
