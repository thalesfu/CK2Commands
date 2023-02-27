package pinega

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维兰VilanBarony struct {
	feud.BaseBarony
}

var BVilan维兰 feud.Barony = &维兰VilanBarony{}

func init() {
    f := BVilan维兰.(*维兰VilanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vilan",
		TitleName: "维兰",
		TitleCode: "b_vilan",
	}
}
