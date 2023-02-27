package narva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷佩尔RepelBarony struct {
	feud.BaseBarony
}

var BRepel雷佩尔 feud.Barony = &雷佩尔RepelBarony{}

func init() {
    f := BRepel雷佩尔.(*雷佩尔RepelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "repel",
		TitleName: "雷佩尔",
		TitleCode: "b_repel",
	}
}
