package yezhuga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尤拉YulaBarony struct {
	feud.BaseBarony
}

var BYula尤拉 feud.Barony = &尤拉YulaBarony{}

func init() {
    f := BYula尤拉.(*尤拉YulaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yula",
		TitleName: "尤拉",
		TitleCode: "b_yula",
	}
}
