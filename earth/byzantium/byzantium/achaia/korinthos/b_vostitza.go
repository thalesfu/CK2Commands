package korinthos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃斯蒂察VostitzaBarony struct {
	feud.BaseBarony
}

var BVostitza沃斯蒂察 feud.Barony = &沃斯蒂察VostitzaBarony{}

func init() {
    f := BVostitza沃斯蒂察.(*沃斯蒂察VostitzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vostitza",
		TitleName: "沃斯蒂察",
		TitleCode: "b_vostitza",
	}
}
