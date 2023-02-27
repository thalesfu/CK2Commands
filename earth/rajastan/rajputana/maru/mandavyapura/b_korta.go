package mandavyapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拘罗吒KortaBarony struct {
	feud.BaseBarony
}

var BKorta拘罗吒 feud.Barony = &拘罗吒KortaBarony{}

func init() {
    f := BKorta拘罗吒.(*拘罗吒KortaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "korta",
		TitleName: "拘罗吒",
		TitleCode: "b_korta",
	}
}
