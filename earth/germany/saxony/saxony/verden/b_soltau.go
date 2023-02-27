package verden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索尔陶SoltauBarony struct {
	feud.BaseBarony
}

var BSoltau索尔陶 feud.Barony = &索尔陶SoltauBarony{}

func init() {
    f := BSoltau索尔陶.(*索尔陶SoltauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "soltau",
		TitleName: "索尔陶",
		TitleCode: "b_soltau",
	}
}
