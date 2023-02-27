package vetluga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡伊尔夫KajirvBarony struct {
	feud.BaseBarony
}

var BKajirv卡伊尔夫 feud.Barony = &卡伊尔夫KajirvBarony{}

func init() {
    f := BKajirv卡伊尔夫.(*卡伊尔夫KajirvBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kajirv",
		TitleName: "卡伊尔夫",
		TitleCode: "b_kajirv",
	}
}
