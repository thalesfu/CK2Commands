package tukums

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡耶瓦KayevaBarony struct {
	feud.BaseBarony
}

var BKayeva卡耶瓦 feud.Barony = &卡耶瓦KayevaBarony{}

func init() {
    f := BKayeva卡耶瓦.(*卡耶瓦KayevaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kayeva",
		TitleName: "卡耶瓦",
		TitleCode: "b_kayeva",
	}
}
