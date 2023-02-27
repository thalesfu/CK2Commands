package semender

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡哈提巴库KhattibakuBarony struct {
	feud.BaseBarony
}

var BKhattibaku卡哈提巴库 feud.Barony = &卡哈提巴库KhattibakuBarony{}

func init() {
    f := BKhattibaku卡哈提巴库.(*卡哈提巴库KhattibakuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khattibaku",
		TitleName: "卡哈提巴库",
		TitleCode: "b_khattibaku",
	}
}
