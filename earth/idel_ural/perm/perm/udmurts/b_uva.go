package udmurts

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌瓦UvaBarony struct {
	feud.BaseBarony
}

var BUva乌瓦 feud.Barony = &乌瓦UvaBarony{}

func init() {
    f := BUva乌瓦.(*乌瓦UvaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uva",
		TitleName: "乌瓦",
		TitleCode: "b_uva",
	}
}
