package reggio

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米莱托MiletoBarony struct {
	feud.BaseBarony
}

var BMileto米莱托 feud.Barony = &米莱托MiletoBarony{}

func init() {
    f := BMileto米莱托.(*米莱托MiletoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mileto",
		TitleName: "米莱托",
		TitleCode: "b_mileto",
	}
}
