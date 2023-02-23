package valabhi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波利旦那PalatinaBarony struct {
	feud.BaseBarony
}

var BPalatina波利旦那 feud.Barony = &波利旦那PalatinaBarony{}

func init() {
	f := BPalatina波利旦那.(*波利旦那PalatinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "palatina",
		TitleName: "波利旦那",
		TitleCode: "b_palatina",
	}
}
