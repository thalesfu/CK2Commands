package mosul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 什汗ShekhanBarony struct {
	feud.BaseBarony
}

var BShekhan什汗 feud.Barony = &什汗ShekhanBarony{}

func init() {
    f := BShekhan什汗.(*什汗ShekhanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shekhan",
		TitleName: "什汗",
		TitleCode: "b_shekhan",
	}
}
