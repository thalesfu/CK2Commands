package keriya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 媲摩PimoBarony struct {
	feud.BaseBarony
}

var BPimo媲摩 feud.Barony = &媲摩PimoBarony{}

func init() {
    f := BPimo媲摩.(*媲摩PimoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pimo",
		TitleName: "媲摩",
		TitleCode: "b_pimo",
	}
}
