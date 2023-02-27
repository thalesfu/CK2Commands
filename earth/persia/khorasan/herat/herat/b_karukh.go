package herat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡鲁赫KarukhBarony struct {
	feud.BaseBarony
}

var BKarukh卡鲁赫 feud.Barony = &卡鲁赫KarukhBarony{}

func init() {
    f := BKarukh卡鲁赫.(*卡鲁赫KarukhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karukh",
		TitleName: "卡鲁赫",
		TitleCode: "b_karukh",
	}
}
