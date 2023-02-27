package murcia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆尔西亚MurciaBarony struct {
	feud.BaseBarony
}

var BMurcia穆尔西亚 feud.Barony = &穆尔西亚MurciaBarony{}

func init() {
    f := BMurcia穆尔西亚.(*穆尔西亚MurciaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "murcia",
		TitleName: "穆尔西亚",
		TitleCode: "b_murcia",
	}
}
