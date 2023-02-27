package lausitz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱布萨LebusaBarony struct {
	feud.BaseBarony
}

var BLebusa莱布萨 feud.Barony = &莱布萨LebusaBarony{}

func init() {
    f := BLebusa莱布萨.(*莱布萨LebusaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lebusa",
		TitleName: "莱布萨",
		TitleCode: "b_lebusa",
	}
}
