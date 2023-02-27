package rutog

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 热帮日土Rawang_rutogBarony struct {
	feud.BaseBarony
}

var BRawang_rutog热帮日土 feud.Barony = &热帮日土Rawang_rutogBarony{}

func init() {
    f := BRawang_rutog热帮日土.(*热帮日土Rawang_rutogBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rawang_rutog",
		TitleName: "热帮日土",
		TitleCode: "b_rawang_rutog",
	}
}
