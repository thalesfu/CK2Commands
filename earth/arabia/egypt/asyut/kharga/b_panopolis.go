package kharga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕诺波利斯PanopolisBarony struct {
	feud.BaseBarony
}

var BPanopolis帕诺波利斯 feud.Barony = &帕诺波利斯PanopolisBarony{}

func init() {
	f := BPanopolis帕诺波利斯.(*帕诺波利斯PanopolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "panopolis",
		TitleName: "帕诺波利斯",
		TitleCode: "b_panopolis",
	}
}
