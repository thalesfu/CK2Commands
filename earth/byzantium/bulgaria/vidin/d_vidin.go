package vidin

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/bulgaria/vidin/naissus"
	"github.com/thalesfu/CK2Commands/earth/byzantium/bulgaria/vidin/vidin"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VidinDuke interface {
	feud.Duke
	CNaissus奈索斯() naissus.NaissusCounty
	CVidin维丁() vidin.VidinCounty
}

type 维丁VidinDuke struct {
	feud.BaseDuke
	奈索斯Naissus naissus.NaissusCounty
	维丁Vidin    vidin.VidinCounty
}

func (d *维丁VidinDuke) CNaissus奈索斯() naissus.NaissusCounty {
	return d.奈索斯Naissus
}

func (d *维丁VidinDuke) CVidin维丁() vidin.VidinCounty {
	return d.维丁Vidin
}

var DVidin维丁 VidinDuke = &维丁VidinDuke{}

func init() {
	f := DVidin维丁.(*维丁VidinDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "vidin",
		TitleName: "维丁",
		TitleCode: "d_vidin",
		Counties:  map[string]feud.County{},
	}

	f.奈索斯Naissus = naissus.CNaissus奈索斯
	f.奈索斯Naissus.SetParent(f)

	f.维丁Vidin = vidin.CVidin维丁
	f.维丁Vidin.SetParent(f)

}
