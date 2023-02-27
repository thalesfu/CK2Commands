package sripuri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 努瓦帕达NuapadaBarony struct {
	feud.BaseBarony
}

var BNuapada努瓦帕达 feud.Barony = &努瓦帕达NuapadaBarony{}

func init() {
    f := BNuapada努瓦帕达.(*努瓦帕达NuapadaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nuapada",
		TitleName: "努瓦帕达",
		TitleCode: "b_nuapada",
	}
}
