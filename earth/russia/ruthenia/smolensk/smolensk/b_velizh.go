package smolensk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦利日VelizhBarony struct {
	feud.BaseBarony
}

var BVelizh韦利日 feud.Barony = &韦利日VelizhBarony{}

func init() {
	f := BVelizh韦利日.(*韦利日VelizhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "velizh",
		TitleName: "韦利日",
		TitleCode: "b_velizh",
	}
}
