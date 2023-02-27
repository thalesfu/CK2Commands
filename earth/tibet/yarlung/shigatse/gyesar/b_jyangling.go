package gyesar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒋陵JyanglingBarony struct {
	feud.BaseBarony
}

var BJyangling蒋陵 feud.Barony = &蒋陵JyanglingBarony{}

func init() {
    f := BJyangling蒋陵.(*蒋陵JyanglingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jyangling",
		TitleName: "蒋陵",
		TitleCode: "b_jyangling",
	}
}
