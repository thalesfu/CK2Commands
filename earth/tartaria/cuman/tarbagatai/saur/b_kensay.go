package saur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坑萨依KensayBarony struct {
	feud.BaseBarony
}

var BKensay坑萨依 feud.Barony = &坑萨依KensayBarony{}

func init() {
    f := BKensay坑萨依.(*坑萨依KensayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kensay",
		TitleName: "坑萨依",
		TitleCode: "b_kensay",
	}
}
