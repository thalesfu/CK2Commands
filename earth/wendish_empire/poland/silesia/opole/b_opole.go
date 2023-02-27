package opole

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥波莱OpoleBarony struct {
	feud.BaseBarony
}

var BOpole奥波莱 feud.Barony = &奥波莱OpoleBarony{}

func init() {
    f := BOpole奥波莱.(*奥波莱OpoleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "opole",
		TitleName: "奥波莱",
		TitleCode: "b_opole",
	}
}
