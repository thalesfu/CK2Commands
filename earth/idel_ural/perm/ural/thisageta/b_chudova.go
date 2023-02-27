package thisageta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 丘多瓦ChudovaBarony struct {
	feud.BaseBarony
}

var BChudova丘多瓦 feud.Barony = &丘多瓦ChudovaBarony{}

func init() {
    f := BChudova丘多瓦.(*丘多瓦ChudovaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chudova",
		TitleName: "丘多瓦",
		TitleCode: "b_chudova",
	}
}
