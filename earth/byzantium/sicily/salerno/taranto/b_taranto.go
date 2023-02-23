package taranto

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔兰托TarantoBarony struct {
	feud.BaseBarony
}

var BTaranto塔兰托 feud.Barony = &塔兰托TarantoBarony{}

func init() {
	f := BTaranto塔兰托.(*塔兰托TarantoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taranto",
		TitleName: "塔兰托",
		TitleCode: "b_taranto",
	}
}
