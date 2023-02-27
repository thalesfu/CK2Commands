package satakunta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌尔维拉UlvilaBarony struct {
	feud.BaseBarony
}

var BUlvila乌尔维拉 feud.Barony = &乌尔维拉UlvilaBarony{}

func init() {
    f := BUlvila乌尔维拉.(*乌尔维拉UlvilaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ulvila",
		TitleName: "乌尔维拉",
		TitleCode: "b_ulvila",
	}
}
