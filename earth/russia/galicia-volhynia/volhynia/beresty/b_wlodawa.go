package beresty

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗沃达瓦WlodawaBarony struct {
	feud.BaseBarony
}

var BWlodawa弗沃达瓦 feud.Barony = &弗沃达瓦WlodawaBarony{}

func init() {
    f := BWlodawa弗沃达瓦.(*弗沃达瓦WlodawaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wlodawa",
		TitleName: "弗沃达瓦",
		TitleCode: "b_wlodawa",
	}
}
