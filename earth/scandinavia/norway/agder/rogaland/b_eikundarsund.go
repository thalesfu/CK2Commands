package rogaland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃昆德松EikundarsundBarony struct {
	feud.BaseBarony
}

var BEikundarsund埃昆德松 feud.Barony = &埃昆德松EikundarsundBarony{}

func init() {
    f := BEikundarsund埃昆德松.(*埃昆德松EikundarsundBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eikundarsund",
		TitleName: "埃昆德松",
		TitleCode: "b_eikundarsund",
	}
}
