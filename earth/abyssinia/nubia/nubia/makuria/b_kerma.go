package makuria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科尔玛KermaBarony struct {
	feud.BaseBarony
}

var BKerma科尔玛 feud.Barony = &科尔玛KermaBarony{}

func init() {
    f := BKerma科尔玛.(*科尔玛KermaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kerma",
		TitleName: "科尔玛",
		TitleCode: "b_kerma",
	}
}
