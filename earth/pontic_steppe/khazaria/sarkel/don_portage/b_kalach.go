package don_portage

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉奇KalachBarony struct {
	feud.BaseBarony
}

var BKalach卡拉奇 feud.Barony = &卡拉奇KalachBarony{}

func init() {
    f := BKalach卡拉奇.(*卡拉奇KalachBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalach",
		TitleName: "卡拉奇",
		TitleCode: "b_kalach",
	}
}
