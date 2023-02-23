package venice

import (
	"github.com/thalesfu/CK2Commands/earth/italy/venice/venice/venezia"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VeniceDuke interface {
	feud.Duke
	CVenezia威尼斯() venezia.VeneziaCounty
}

type 威尼斯VeniceDuke struct {
	feud.BaseDuke
	威尼斯Venezia venezia.VeneziaCounty
}

func (d *威尼斯VeniceDuke) CVenezia威尼斯() venezia.VeneziaCounty {
	return d.威尼斯Venezia
}

var DVenice威尼斯 VeniceDuke = &威尼斯VeniceDuke{}

func init() {
	f := DVenice威尼斯.(*威尼斯VeniceDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "venice",
		TitleName: "威尼斯",
		TitleCode: "d_venice",
		Counties:  map[string]feud.County{},
	}

	f.威尼斯Venezia = venezia.CVenezia威尼斯
	f.威尼斯Venezia.SetParent(f)

}
