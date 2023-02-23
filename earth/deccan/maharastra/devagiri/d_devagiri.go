package devagiri

import (
	"github.com/thalesfu/CK2Commands/earth/deccan/maharastra/devagiri/devagiri"
	"github.com/thalesfu/CK2Commands/earth/deccan/maharastra/devagiri/nanded"
	"github.com/thalesfu/CK2Commands/earth/deccan/maharastra/devagiri/parnakheta"
	"github.com/thalesfu/CK2Commands/earth/deccan/maharastra/devagiri/pratishthana"
	"github.com/thalesfu/CK2Commands/earth/deccan/maharastra/devagiri/vatsagulma"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DevagiriDuke interface {
	feud.Duke
	CDevagiri提婆耆厘() devagiri.DevagiriCounty
	CNanded难提陀() nanded.NandedCounty
	CParnakheta半挐契吒() parnakheta.ParnakhetaCounty
	CPratishthana波罗提瑟吒那() pratishthana.PratishthanaCounty
	CVatsagulma婆蹉崛摩() vatsagulma.VatsagulmaCounty
}

type 提婆耆厘DevagiriDuke struct {
	feud.BaseDuke
	提婆耆厘Devagiri       devagiri.DevagiriCounty
	难提陀Nanded          nanded.NandedCounty
	半挐契吒Parnakheta     parnakheta.ParnakhetaCounty
	波罗提瑟吒那Pratishthana pratishthana.PratishthanaCounty
	婆蹉崛摩Vatsagulma     vatsagulma.VatsagulmaCounty
}

func (d *提婆耆厘DevagiriDuke) CDevagiri提婆耆厘() devagiri.DevagiriCounty {
	return d.提婆耆厘Devagiri
}

func (d *提婆耆厘DevagiriDuke) CNanded难提陀() nanded.NandedCounty {
	return d.难提陀Nanded
}

func (d *提婆耆厘DevagiriDuke) CParnakheta半挐契吒() parnakheta.ParnakhetaCounty {
	return d.半挐契吒Parnakheta
}

func (d *提婆耆厘DevagiriDuke) CPratishthana波罗提瑟吒那() pratishthana.PratishthanaCounty {
	return d.波罗提瑟吒那Pratishthana
}

func (d *提婆耆厘DevagiriDuke) CVatsagulma婆蹉崛摩() vatsagulma.VatsagulmaCounty {
	return d.婆蹉崛摩Vatsagulma
}

var DDevagiri提婆耆厘 DevagiriDuke = &提婆耆厘DevagiriDuke{}

func init() {
	f := DDevagiri提婆耆厘.(*提婆耆厘DevagiriDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "devagiri",
		TitleName: "提婆耆厘",
		TitleCode: "d_devagiri",
		Counties:  map[string]feud.County{},
	}

	f.提婆耆厘Devagiri = devagiri.CDevagiri提婆耆厘
	f.提婆耆厘Devagiri.SetParent(f)

	f.难提陀Nanded = nanded.CNanded难提陀
	f.难提陀Nanded.SetParent(f)

	f.半挐契吒Parnakheta = parnakheta.CParnakheta半挐契吒
	f.半挐契吒Parnakheta.SetParent(f)

	f.波罗提瑟吒那Pratishthana = pratishthana.CPratishthana波罗提瑟吒那
	f.波罗提瑟吒那Pratishthana.SetParent(f)

	f.婆蹉崛摩Vatsagulma = vatsagulma.CVatsagulma婆蹉崛摩
	f.婆蹉崛摩Vatsagulma.SetParent(f)

}
