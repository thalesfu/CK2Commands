package bjarmia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 因齐IntsyBarony struct {
	feud.BaseBarony
}

var BIntsy因齐 feud.Barony = &因齐IntsyBarony{}

func init() {
    f := BIntsy因齐.(*因齐IntsyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "intsy",
		TitleName: "因齐",
		TitleCode: "b_intsy",
	}
}
