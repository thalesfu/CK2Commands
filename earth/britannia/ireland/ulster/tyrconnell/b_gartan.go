package tyrconnell

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加坦GartanBarony struct {
	feud.BaseBarony
}

var BGartan加坦 feud.Barony = &加坦GartanBarony{}

func init() {
    f := BGartan加坦.(*加坦GartanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gartan",
		TitleName: "加坦",
		TitleCode: "b_gartan",
	}
}
