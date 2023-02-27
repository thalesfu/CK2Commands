package jarva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维尔扬迪ViljandiBarony struct {
	feud.BaseBarony
}

var BViljandi维尔扬迪 feud.Barony = &维尔扬迪ViljandiBarony{}

func init() {
    f := BViljandi维尔扬迪.(*维尔扬迪ViljandiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "viljandi",
		TitleName: "维尔扬迪",
		TitleCode: "b_viljandi",
	}
}
