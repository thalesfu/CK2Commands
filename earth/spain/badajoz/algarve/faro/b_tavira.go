package faro

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔维拉TaviraBarony struct {
	feud.BaseBarony
}

var BTavira塔维拉 feud.Barony = &塔维拉TaviraBarony{}

func init() {
    f := BTavira塔维拉.(*塔维拉TaviraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tavira",
		TitleName: "塔维拉",
		TitleCode: "b_tavira",
	}
}
