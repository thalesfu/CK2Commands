package utva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌特瓦UtvaBarony struct {
	feud.BaseBarony
}

var BUtva乌特瓦 feud.Barony = &乌特瓦UtvaBarony{}

func init() {
    f := BUtva乌特瓦.(*乌特瓦UtvaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "utva",
		TitleName: "乌特瓦",
		TitleCode: "b_utva",
	}
}
