package coimbra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷东杜PedondoBarony struct {
	feud.BaseBarony
}

var BPedondo雷东杜 feud.Barony = &雷东杜PedondoBarony{}

func init() {
	f := BPedondo雷东杜.(*雷东杜PedondoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pedondo",
		TitleName: "雷东杜",
		TitleCode: "b_pedondo",
	}
}
