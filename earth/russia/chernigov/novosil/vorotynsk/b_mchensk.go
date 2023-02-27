package vorotynsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 姆岑斯克MchenskBarony struct {
	feud.BaseBarony
}

var BMchensk姆岑斯克 feud.Barony = &姆岑斯克MchenskBarony{}

func init() {
    f := BMchensk姆岑斯克.(*姆岑斯克MchenskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mchensk",
		TitleName: "姆岑斯克",
		TitleCode: "b_mchensk",
	}
}
