package kollipake

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拘梨波稽KollipakeBarony struct {
	feud.BaseBarony
}

var BKollipake拘梨波稽 feud.Barony = &拘梨波稽KollipakeBarony{}

func init() {
	f := BKollipake拘梨波稽.(*拘梨波稽KollipakeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kollipake",
		TitleName: "拘梨波稽",
		TitleCode: "b_kollipake",
	}
}
