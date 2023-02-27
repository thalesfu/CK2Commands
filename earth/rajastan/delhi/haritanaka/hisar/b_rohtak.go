package hisar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢醯呾迦RohtakBarony struct {
	feud.BaseBarony
}

var BRohtak卢醯呾迦 feud.Barony = &卢醯呾迦RohtakBarony{}

func init() {
    f := BRohtak卢醯呾迦.(*卢醯呾迦RohtakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rohtak",
		TitleName: "卢醯呾迦",
		TitleCode: "b_rohtak",
	}
}
