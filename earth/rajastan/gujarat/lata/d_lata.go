package lata

import (
	"github.com/thalesfu/CK2Commands/earth/rajastan/gujarat/lata/daman"
	"github.com/thalesfu/CK2Commands/earth/rajastan/gujarat/lata/navasarika"
	"github.com/thalesfu/CK2Commands/earth/rajastan/gujarat/lata/vadodara"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LataDuke interface {
	feud.Duke
	CDaman陀曼那() daman.DamanCounty
	CNavasarika那婆萨利迦() navasarika.NavasarikaCounty
	CVadodara婆度陀罗() vadodara.VadodaraCounty
}

type 罗吒LataDuke struct {
	feud.BaseDuke
	陀曼那Daman        daman.DamanCounty
	那婆萨利迦Navasarika navasarika.NavasarikaCounty
	婆度陀罗Vadodara    vadodara.VadodaraCounty
}

func (d *罗吒LataDuke) CDaman陀曼那() daman.DamanCounty {
	return d.陀曼那Daman
}

func (d *罗吒LataDuke) CNavasarika那婆萨利迦() navasarika.NavasarikaCounty {
	return d.那婆萨利迦Navasarika
}

func (d *罗吒LataDuke) CVadodara婆度陀罗() vadodara.VadodaraCounty {
	return d.婆度陀罗Vadodara
}

var DLata罗吒 LataDuke = &罗吒LataDuke{}

func init() {
	f := DLata罗吒.(*罗吒LataDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "lata",
		TitleName: "罗吒",
		TitleCode: "d_lata",
		Counties:  map[string]feud.County{},
	}

	f.陀曼那Daman = daman.CDaman陀曼那
	f.陀曼那Daman.SetParent(f)

	f.那婆萨利迦Navasarika = navasarika.CNavasarika那婆萨利迦
	f.那婆萨利迦Navasarika.SetParent(f)

	f.婆度陀罗Vadodara = vadodara.CVadodara婆度陀罗
	f.婆度陀罗Vadodara.SetParent(f)

}
