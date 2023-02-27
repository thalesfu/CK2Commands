package poznanskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗斯霍瓦WschowaBarony struct {
	feud.BaseBarony
}

var BWschowa弗斯霍瓦 feud.Barony = &弗斯霍瓦WschowaBarony{}

func init() {
    f := BWschowa弗斯霍瓦.(*弗斯霍瓦WschowaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wschowa",
		TitleName: "弗斯霍瓦",
		TitleCode: "b_wschowa",
	}
}
