package talava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨特克莱SatekleBarony struct {
	feud.BaseBarony
}

var BSatekle萨特克莱 feud.Barony = &萨特克莱SatekleBarony{}

func init() {
    f := BSatekle萨特克莱.(*萨特克莱SatekleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "satekle",
		TitleName: "萨特克莱",
		TitleCode: "b_satekle",
	}
}
