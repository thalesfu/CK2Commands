package kosti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科斯提KostiBarony struct {
	feud.BaseBarony
}

var BKosti科斯提 feud.Barony = &科斯提KostiBarony{}

func init() {
    f := BKosti科斯提.(*科斯提KostiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kosti",
		TitleName: "科斯提",
		TitleCode: "b_kosti",
	}
}
