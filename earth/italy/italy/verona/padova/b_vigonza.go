package padova

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维贡扎VigonzaBarony struct {
	feud.BaseBarony
}

var BVigonza维贡扎 feud.Barony = &维贡扎VigonzaBarony{}

func init() {
	f := BVigonza维贡扎.(*维贡扎VigonzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vigonza",
		TitleName: "维贡扎",
		TitleCode: "b_vigonza",
	}
}
