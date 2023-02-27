package norfolk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃尔门ElmhamBarony struct {
	feud.BaseBarony
}

var BElmham埃尔门 feud.Barony = &埃尔门ElmhamBarony{}

func init() {
    f := BElmham埃尔门.(*埃尔门ElmhamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elmham",
		TitleName: "埃尔门",
		TitleCode: "b_elmham",
	}
}
