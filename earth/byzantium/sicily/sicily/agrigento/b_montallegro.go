package agrigento

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙塔莱格罗MontallegroBarony struct {
	feud.BaseBarony
}

var BMontallegro蒙塔莱格罗 feud.Barony = &蒙塔莱格罗MontallegroBarony{}

func init() {
    f := BMontallegro蒙塔莱格罗.(*蒙塔莱格罗MontallegroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montallegro",
		TitleName: "蒙塔莱格罗",
		TitleCode: "b_montallegro",
	}
}
