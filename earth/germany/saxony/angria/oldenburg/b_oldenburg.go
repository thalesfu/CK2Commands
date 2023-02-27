package oldenburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥尔登堡OldenburgBarony struct {
	feud.BaseBarony
}

var BOldenburg奥尔登堡 feud.Barony = &奥尔登堡OldenburgBarony{}

func init() {
    f := BOldenburg奥尔登堡.(*奥尔登堡OldenburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oldenburg",
		TitleName: "奥尔登堡",
		TitleCode: "b_oldenburg",
	}
}
