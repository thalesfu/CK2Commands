package qazwin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布津扎拉BuinzahraBarony struct {
	feud.BaseBarony
}

var BBuinzahra布津扎拉 feud.Barony = &布津扎拉BuinzahraBarony{}

func init() {
    f := BBuinzahra布津扎拉.(*布津扎拉BuinzahraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "buinzahra",
		TitleName: "布津扎拉",
		TitleCode: "b_buinzahra",
	}
}
