package tribandapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 支揭由ChitwyeBarony struct {
	feud.BaseBarony
}

var BChitwye支揭由 feud.Barony = &支揭由ChitwyeBarony{}

func init() {
    f := BChitwye支揭由.(*支揭由ChitwyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chitwye",
		TitleName: "支揭由",
		TitleCode: "b_chitwye",
	}
}
