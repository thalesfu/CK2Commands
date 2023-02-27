package asas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎佐乌ZaazouaaBarony struct {
	feud.BaseBarony
}

var BZaazouaa扎佐乌 feud.Barony = &扎佐乌ZaazouaaBarony{}

func init() {
    f := BZaazouaa扎佐乌.(*扎佐乌ZaazouaaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zaazouaa",
		TitleName: "扎佐乌",
		TitleCode: "b_zaazouaa",
	}
}
