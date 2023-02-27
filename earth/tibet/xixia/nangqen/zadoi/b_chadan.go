package zadoi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 查旦ChadanBarony struct {
	feud.BaseBarony
}

var BChadan查旦 feud.Barony = &查旦ChadanBarony{}

func init() {
    f := BChadan查旦.(*查旦ChadanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chadan",
		TitleName: "查旦",
		TitleCode: "b_chadan",
	}
}
