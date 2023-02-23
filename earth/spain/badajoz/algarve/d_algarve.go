package algarve

import (
	"github.com/thalesfu/CK2Commands/earth/spain/badajoz/algarve/faro"
	"github.com/thalesfu/CK2Commands/earth/spain/badajoz/algarve/silves"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AlgarveDuke interface {
	feud.Duke
	CFaro法鲁() faro.FaroCounty
	CSilves锡尔维什() silves.SilvesCounty
}

type 阿尔加维AlgarveDuke struct {
	feud.BaseDuke
	法鲁Faro     faro.FaroCounty
	锡尔维什Silves silves.SilvesCounty
}

func (d *阿尔加维AlgarveDuke) CFaro法鲁() faro.FaroCounty {
	return d.法鲁Faro
}

func (d *阿尔加维AlgarveDuke) CSilves锡尔维什() silves.SilvesCounty {
	return d.锡尔维什Silves
}

var DAlgarve阿尔加维 AlgarveDuke = &阿尔加维AlgarveDuke{}

func init() {
	f := DAlgarve阿尔加维.(*阿尔加维AlgarveDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "algarve",
		TitleName: "阿尔加维",
		TitleCode: "d_algarve",
		Counties:  map[string]feud.County{},
	}

	f.法鲁Faro = faro.CFaro法鲁
	f.法鲁Faro.SetParent(f)

	f.锡尔维什Silves = silves.CSilves锡尔维什
	f.锡尔维什Silves.SetParent(f)

}
