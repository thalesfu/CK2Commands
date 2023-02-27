package pelusia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞得SaidBarony struct {
	feud.BaseBarony
}

var BSaid塞得 feud.Barony = &塞得SaidBarony{}

func init() {
    f := BSaid塞得.(*塞得SaidBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "said",
		TitleName: "塞得",
		TitleCode: "b_said",
	}
}
