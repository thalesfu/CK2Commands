package al_ahqaf

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾哈嘎弗Al_ahqafBarony struct {
	feud.BaseBarony
}

var BAl_ahqaf艾哈嘎弗 feud.Barony = &艾哈嘎弗Al_ahqafBarony{}

func init() {
    f := BAl_ahqaf艾哈嘎弗.(*艾哈嘎弗Al_ahqafBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "al_ahqaf",
		TitleName: "艾哈嘎弗",
		TitleCode: "b_al_ahqaf",
	}
}
