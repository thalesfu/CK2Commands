package kangly

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉沙KarashaBarony struct {
	feud.BaseBarony
}

var BKarasha卡拉沙 feud.Barony = &卡拉沙KarashaBarony{}

func init() {
    f := BKarasha卡拉沙.(*卡拉沙KarashaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karasha",
		TitleName: "卡拉沙",
		TitleCode: "b_karasha",
	}
}
