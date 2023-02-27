package jarva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫胡MohuBarony struct {
	feud.BaseBarony
}

var BMohu莫胡 feud.Barony = &莫胡MohuBarony{}

func init() {
    f := BMohu莫胡.(*莫胡MohuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mohu",
		TitleName: "莫胡",
		TitleCode: "b_mohu",
	}
}
