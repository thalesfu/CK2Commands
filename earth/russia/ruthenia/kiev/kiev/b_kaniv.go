package kiev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尼夫KanivBarony struct {
	feud.BaseBarony
}

var BKaniv卡尼夫 feud.Barony = &卡尼夫KanivBarony{}

func init() {
	f := BKaniv卡尼夫.(*卡尼夫KanivBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kaniv",
		TitleName: "卡尼夫",
		TitleCode: "b_kaniv",
	}
}
