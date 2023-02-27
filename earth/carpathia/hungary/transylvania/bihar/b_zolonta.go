package bihar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佐龙陶ZolontaBarony struct {
	feud.BaseBarony
}

var BZolonta佐龙陶 feud.Barony = &佐龙陶ZolontaBarony{}

func init() {
    f := BZolonta佐龙陶.(*佐龙陶ZolontaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zolonta",
		TitleName: "佐龙陶",
		TitleCode: "b_zolonta",
	}
}
