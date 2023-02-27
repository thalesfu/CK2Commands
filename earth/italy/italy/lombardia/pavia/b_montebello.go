package pavia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙特贝洛MontebelloBarony struct {
	feud.BaseBarony
}

var BMontebello蒙特贝洛 feud.Barony = &蒙特贝洛MontebelloBarony{}

func init() {
    f := BMontebello蒙特贝洛.(*蒙特贝洛MontebelloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montebello",
		TitleName: "蒙特贝洛",
		TitleCode: "b_montebello",
	}
}
