package kalpi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌伊罗伊OiraiBarony struct {
	feud.BaseBarony
}

var BOirai乌伊罗伊 feud.Barony = &乌伊罗伊OiraiBarony{}

func init() {
    f := BOirai乌伊罗伊.(*乌伊罗伊OiraiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oirai",
		TitleName: "乌伊罗伊",
		TitleCode: "b_oirai",
	}
}
