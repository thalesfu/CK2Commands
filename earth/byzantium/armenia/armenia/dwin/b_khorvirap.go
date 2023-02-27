package dwin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍尔维拉普KhorvirapBarony struct {
	feud.BaseBarony
}

var BKhorvirap霍尔维拉普 feud.Barony = &霍尔维拉普KhorvirapBarony{}

func init() {
    f := BKhorvirap霍尔维拉普.(*霍尔维拉普KhorvirapBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khorvirap",
		TitleName: "霍尔维拉普",
		TitleCode: "b_khorvirap",
	}
}
