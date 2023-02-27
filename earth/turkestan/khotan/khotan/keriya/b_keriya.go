package keriya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克里雅KeriyaBarony struct {
	feud.BaseBarony
}

var BKeriya克里雅 feud.Barony = &克里雅KeriyaBarony{}

func init() {
    f := BKeriya克里雅.(*克里雅KeriyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "keriya",
		TitleName: "克里雅",
		TitleCode: "b_keriya",
	}
}
