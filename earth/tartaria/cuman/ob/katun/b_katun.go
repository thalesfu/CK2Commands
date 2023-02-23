package katun

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡通KatunBarony struct {
	feud.BaseBarony
}

var BKatun卡通 feud.Barony = &卡通KatunBarony{}

func init() {
	f := BKatun卡通.(*卡通KatunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "katun",
		TitleName: "卡通",
		TitleCode: "b_katun",
	}
}
