package bost

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坎大哈KandaharBarony struct {
	feud.BaseBarony
}

var BKandahar坎大哈 feud.Barony = &坎大哈KandaharBarony{}

func init() {
    f := BKandahar坎大哈.(*坎大哈KandaharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kandahar",
		TitleName: "坎大哈",
		TitleCode: "b_kandahar",
	}
}
