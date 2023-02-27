package talava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅特塞内MetseneBarony struct {
	feud.BaseBarony
}

var BMetsene梅特塞内 feud.Barony = &梅特塞内MetseneBarony{}

func init() {
    f := BMetsene梅特塞内.(*梅特塞内MetseneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "metsene",
		TitleName: "梅特塞内",
		TitleCode: "b_metsene",
	}
}
