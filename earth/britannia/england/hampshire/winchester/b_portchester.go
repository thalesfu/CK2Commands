package winchester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 彻斯特港PortchesterBarony struct {
	feud.BaseBarony
}

var BPortchester彻斯特港 feud.Barony = &彻斯特港PortchesterBarony{}

func init() {
	f := BPortchester彻斯特港.(*彻斯特港PortchesterBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "portchester",
		TitleName: "彻斯特港",
		TitleCode: "b_portchester",
	}
}
