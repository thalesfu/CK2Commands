package buzachi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉基丘KarakishuBarony struct {
	feud.BaseBarony
}

var BKarakishu卡拉基丘 feud.Barony = &卡拉基丘KarakishuBarony{}

func init() {
	f := BKarakishu卡拉基丘.(*卡拉基丘KarakishuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karakishu",
		TitleName: "卡拉基丘",
		TitleCode: "b_karakishu",
	}
}
