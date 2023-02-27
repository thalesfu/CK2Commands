package khopyor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌瓦罗沃UvarovoBarony struct {
	feud.BaseBarony
}

var BUvarovo乌瓦罗沃 feud.Barony = &乌瓦罗沃UvarovoBarony{}

func init() {
    f := BUvarovo乌瓦罗沃.(*乌瓦罗沃UvarovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uvarovo",
		TitleName: "乌瓦罗沃",
		TitleCode: "b_uvarovo",
	}
}
