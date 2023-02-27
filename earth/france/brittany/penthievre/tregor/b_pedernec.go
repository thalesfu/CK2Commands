package tregor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩代尔内克PedernecBarony struct {
	feud.BaseBarony
}

var BPedernec佩代尔内克 feud.Barony = &佩代尔内克PedernecBarony{}

func init() {
    f := BPedernec佩代尔内克.(*佩代尔内克PedernecBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pedernec",
		TitleName: "佩代尔内克",
		TitleCode: "b_pedernec",
	}
}
