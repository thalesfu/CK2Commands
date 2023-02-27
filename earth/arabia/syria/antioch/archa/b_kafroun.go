package archa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡夫伦KafrounBarony struct {
	feud.BaseBarony
}

var BKafroun卡夫伦 feud.Barony = &卡夫伦KafrounBarony{}

func init() {
    f := BKafroun卡夫伦.(*卡夫伦KafrounBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kafroun",
		TitleName: "卡夫伦",
		TitleCode: "b_kafroun",
	}
}
