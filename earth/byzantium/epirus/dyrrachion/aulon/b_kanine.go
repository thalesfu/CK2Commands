package aulon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尼内KanineBarony struct {
	feud.BaseBarony
}

var BKanine卡尼内 feud.Barony = &卡尼内KanineBarony{}

func init() {
	f := BKanine卡尼内.(*卡尼内KanineBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kanine",
		TitleName: "卡尼内",
		TitleCode: "b_kanine",
	}
}
