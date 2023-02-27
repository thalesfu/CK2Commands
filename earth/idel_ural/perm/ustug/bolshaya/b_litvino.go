package bolshaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利特维诺LitvinoBarony struct {
	feud.BaseBarony
}

var BLitvino利特维诺 feud.Barony = &利特维诺LitvinoBarony{}

func init() {
    f := BLitvino利特维诺.(*利特维诺LitvinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "litvino",
		TitleName: "利特维诺",
		TitleCode: "b_litvino",
	}
}
