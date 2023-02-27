package tjust

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维奈斯VinasBarony struct {
	feud.BaseBarony
}

var BVinas维奈斯 feud.Barony = &维奈斯VinasBarony{}

func init() {
    f := BVinas维奈斯.(*维奈斯VinasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vinas",
		TitleName: "维奈斯",
		TitleCode: "b_vinas",
	}
}
