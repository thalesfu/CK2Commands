package durdzukia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米祖尔MizurBarony struct {
	feud.BaseBarony
}

var BMizur米祖尔 feud.Barony = &米祖尔MizurBarony{}

func init() {
    f := BMizur米祖尔.(*米祖尔MizurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mizur",
		TitleName: "米祖尔",
		TitleCode: "b_mizur",
	}
}
