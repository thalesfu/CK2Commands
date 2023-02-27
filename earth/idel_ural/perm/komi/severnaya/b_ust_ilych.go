package severnaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌斯季伊雷奇Ust_ilychBarony struct {
	feud.BaseBarony
}

var BUst_ilych乌斯季伊雷奇 feud.Barony = &乌斯季伊雷奇Ust_ilychBarony{}

func init() {
    f := BUst_ilych乌斯季伊雷奇.(*乌斯季伊雷奇Ust_ilychBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ust_ilych",
		TitleName: "乌斯季伊雷奇",
		TitleCode: "b_ust_ilych",
	}
}
