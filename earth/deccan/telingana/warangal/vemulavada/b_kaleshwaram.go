package vemulavada

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡勒什瓦拉姆KaleshwaramBarony struct {
	feud.BaseBarony
}

var BKaleshwaram卡勒什瓦拉姆 feud.Barony = &卡勒什瓦拉姆KaleshwaramBarony{}

func init() {
	f := BKaleshwaram卡勒什瓦拉姆.(*卡勒什瓦拉姆KaleshwaramBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kaleshwaram",
		TitleName: "卡勒什瓦拉姆",
		TitleCode: "b_kaleshwaram",
	}
}
