package kromy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 涅恰耶瓦NechayevaBarony struct {
	feud.BaseBarony
}

var BNechayeva涅恰耶瓦 feud.Barony = &涅恰耶瓦NechayevaBarony{}

func init() {
    f := BNechayeva涅恰耶瓦.(*涅恰耶瓦NechayevaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nechayeva",
		TitleName: "涅恰耶瓦",
		TitleCode: "b_nechayeva",
	}
}
