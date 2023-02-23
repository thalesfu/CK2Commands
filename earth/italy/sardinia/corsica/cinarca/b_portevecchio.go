package cinarca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦基奥港PortevecchioBarony struct {
	feud.BaseBarony
}

var BPortevecchio韦基奥港 feud.Barony = &韦基奥港PortevecchioBarony{}

func init() {
	f := BPortevecchio韦基奥港.(*韦基奥港PortevecchioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "portevecchio",
		TitleName: "韦基奥港",
		TitleCode: "b_portevecchio",
	}
}
