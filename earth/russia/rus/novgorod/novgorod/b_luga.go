package novgorod

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢加LugaBarony struct {
	feud.BaseBarony
}

var BLuga卢加 feud.Barony = &卢加LugaBarony{}

func init() {
	f := BLuga卢加.(*卢加LugaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "luga",
		TitleName: "卢加",
		TitleCode: "b_luga",
	}
}
