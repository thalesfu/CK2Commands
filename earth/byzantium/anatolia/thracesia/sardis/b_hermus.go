package sardis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫耳墨斯HermusBarony struct {
	feud.BaseBarony
}

var BHermus赫耳墨斯 feud.Barony = &赫耳墨斯HermusBarony{}

func init() {
	f := BHermus赫耳墨斯.(*赫耳墨斯HermusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hermus",
		TitleName: "赫耳墨斯",
		TitleCode: "b_hermus",
	}
}
