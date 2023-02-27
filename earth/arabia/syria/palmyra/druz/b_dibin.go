package druz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪比因DibinBarony struct {
	feud.BaseBarony
}

var BDibin迪比因 feud.Barony = &迪比因DibinBarony{}

func init() {
    f := BDibin迪比因.(*迪比因DibinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dibin",
		TitleName: "迪比因",
		TitleCode: "b_dibin",
	}
}
