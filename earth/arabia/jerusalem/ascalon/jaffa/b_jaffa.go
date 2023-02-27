package jaffa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雅法JaffaBarony struct {
	feud.BaseBarony
}

var BJaffa雅法 feud.Barony = &雅法JaffaBarony{}

func init() {
    f := BJaffa雅法.(*雅法JaffaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jaffa",
		TitleName: "雅法",
		TitleCode: "b_jaffa",
	}
}
