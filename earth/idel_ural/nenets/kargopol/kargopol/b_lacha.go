package kargopol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉恰LachaBarony struct {
	feud.BaseBarony
}

var BLacha拉恰 feud.Barony = &拉恰LachaBarony{}

func init() {
    f := BLacha拉恰.(*拉恰LachaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lacha",
		TitleName: "拉恰",
		TitleCode: "b_lacha",
	}
}
