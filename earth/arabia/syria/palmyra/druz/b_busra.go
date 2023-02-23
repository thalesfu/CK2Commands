package druz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布斯拉BusraBarony struct {
	feud.BaseBarony
}

var BBusra布斯拉 feud.Barony = &布斯拉BusraBarony{}

func init() {
	f := BBusra布斯拉.(*布斯拉BusraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "busra",
		TitleName: "布斯拉",
		TitleCode: "b_busra",
	}
}
