package quzdar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博登BodeinBarony struct {
	feud.BaseBarony
}

var BBodein博登 feud.Barony = &博登BodeinBarony{}

func init() {
    f := BBodein博登.(*博登BodeinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bodein",
		TitleName: "博登",
		TitleCode: "b_bodein",
	}
}
