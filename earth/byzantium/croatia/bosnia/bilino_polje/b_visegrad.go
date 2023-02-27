package bilino_polje

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维舍格勒VisegradBarony struct {
	feud.BaseBarony
}

var BVisegrad维舍格勒 feud.Barony = &维舍格勒VisegradBarony{}

func init() {
    f := BVisegrad维舍格勒.(*维舍格勒VisegradBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "visegrad",
		TitleName: "维舍格勒",
		TitleCode: "b_visegrad",
	}
}
