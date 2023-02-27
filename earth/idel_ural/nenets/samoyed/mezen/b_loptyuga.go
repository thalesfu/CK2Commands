package mezen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛普秋加LoptyugaBarony struct {
	feud.BaseBarony
}

var BLoptyuga洛普秋加 feud.Barony = &洛普秋加LoptyugaBarony{}

func init() {
    f := BLoptyuga洛普秋加.(*洛普秋加LoptyugaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "loptyuga",
		TitleName: "洛普秋加",
		TitleCode: "b_loptyuga",
	}
}
