package maan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布赛拉BuseiraBarony struct {
	feud.BaseBarony
}

var BBuseira布赛拉 feud.Barony = &布赛拉BuseiraBarony{}

func init() {
    f := BBuseira布赛拉.(*布赛拉BuseiraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "buseira",
		TitleName: "布赛拉",
		TitleCode: "b_buseira",
	}
}
