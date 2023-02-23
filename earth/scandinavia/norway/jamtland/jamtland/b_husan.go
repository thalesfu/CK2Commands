package jamtland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 许松HusanBarony struct {
	feud.BaseBarony
}

var BHusan许松 feud.Barony = &许松HusanBarony{}

func init() {
	f := BHusan许松.(*许松HusanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "husan",
		TitleName: "许松",
		TitleCode: "b_husan",
	}
}
