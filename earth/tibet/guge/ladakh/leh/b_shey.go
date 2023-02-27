package leh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雪依SheyBarony struct {
	feud.BaseBarony
}

var BShey雪依 feud.Barony = &雪依SheyBarony{}

func init() {
    f := BShey雪依.(*雪依SheyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shey",
		TitleName: "雪依",
		TitleCode: "b_shey",
	}
}
