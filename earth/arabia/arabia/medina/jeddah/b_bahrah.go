package jeddah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拜赫拉BahrahBarony struct {
	feud.BaseBarony
}

var BBahrah拜赫拉 feud.Barony = &拜赫拉BahrahBarony{}

func init() {
    f := BBahrah拜赫拉.(*拜赫拉BahrahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bahrah",
		TitleName: "拜赫拉",
		TitleCode: "b_bahrah",
	}
}
