package skane

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫雷瓦德HerrevadBarony struct {
	feud.BaseBarony
}

var BHerrevad赫雷瓦德 feud.Barony = &赫雷瓦德HerrevadBarony{}

func init() {
	f := BHerrevad赫雷瓦德.(*赫雷瓦德HerrevadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "herrevad",
		TitleName: "赫雷瓦德",
		TitleCode: "b_herrevad",
	}
}
