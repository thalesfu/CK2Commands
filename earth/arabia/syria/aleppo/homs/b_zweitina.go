package homs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 祖韦提奈ZweitinaBarony struct {
	feud.BaseBarony
}

var BZweitina祖韦提奈 feud.Barony = &祖韦提奈ZweitinaBarony{}

func init() {
	f := BZweitina祖韦提奈.(*祖韦提奈ZweitinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zweitina",
		TitleName: "祖韦提奈",
		TitleCode: "b_zweitina",
	}
}
