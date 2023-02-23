package amdo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎仁ZaringBarony struct {
	feud.BaseBarony
}

var BZaring扎仁 feud.Barony = &扎仁ZaringBarony{}

func init() {
	f := BZaring扎仁.(*扎仁ZaringBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zaring",
		TitleName: "扎仁",
		TitleCode: "b_zaring",
	}
}
