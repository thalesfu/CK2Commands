package schwaben

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 旺根WangenBarony struct {
	feud.BaseBarony
}

var BWangen旺根 feud.Barony = &旺根WangenBarony{}

func init() {
    f := BWangen旺根.(*旺根WangenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wangen",
		TitleName: "旺根",
		TitleCode: "b_wangen",
	}
}
