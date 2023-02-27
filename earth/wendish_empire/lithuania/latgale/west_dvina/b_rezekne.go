package west_dvina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷泽克内RezekneBarony struct {
	feud.BaseBarony
}

var BRezekne雷泽克内 feud.Barony = &雷泽克内RezekneBarony{}

func init() {
    f := BRezekne雷泽克内.(*雷泽克内RezekneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rezekne",
		TitleName: "雷泽克内",
		TitleCode: "b_rezekne",
	}
}
