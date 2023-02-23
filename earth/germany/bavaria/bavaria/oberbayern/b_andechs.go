package oberbayern

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安德希斯AndechsBarony struct {
	feud.BaseBarony
}

var BAndechs安德希斯 feud.Barony = &安德希斯AndechsBarony{}

func init() {
	f := BAndechs安德希斯.(*安德希斯AndechsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "andechs",
		TitleName: "安德希斯",
		TitleCode: "b_andechs",
	}
}
