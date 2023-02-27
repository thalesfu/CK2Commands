package litomerice

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 日塔瓦ZitavaBarony struct {
	feud.BaseBarony
}

var BZitava日塔瓦 feud.Barony = &日塔瓦ZitavaBarony{}

func init() {
    f := BZitava日塔瓦.(*日塔瓦ZitavaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zitava",
		TitleName: "日塔瓦",
		TitleCode: "b_zitava",
	}
}
