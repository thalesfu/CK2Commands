package murzuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞卜奈SebnaBarony struct {
	feud.BaseBarony
}

var BSebna塞卜奈 feud.Barony = &塞卜奈SebnaBarony{}

func init() {
    f := BSebna塞卜奈.(*塞卜奈SebnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sebna",
		TitleName: "塞卜奈",
		TitleCode: "b_sebna",
	}
}
