package kozhva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊兹亚尤Iz_yayuBarony struct {
	feud.BaseBarony
}

var BIz_yayu伊兹亚尤 feud.Barony = &伊兹亚尤Iz_yayuBarony{}

func init() {
    f := BIz_yayu伊兹亚尤.(*伊兹亚尤Iz_yayuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "iz_yayu",
		TitleName: "伊兹亚尤",
		TitleCode: "b_iz_yayu",
	}
}
