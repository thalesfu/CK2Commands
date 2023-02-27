package kirghiz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曳咥IrtyshBarony struct {
	feud.BaseBarony
}

var BIrtysh曳咥 feud.Barony = &曳咥IrtyshBarony{}

func init() {
    f := BIrtysh曳咥.(*曳咥IrtyshBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "irtysh",
		TitleName: "曳咥",
		TitleCode: "b_irtysh",
	}
}
