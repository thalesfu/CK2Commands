package buqtirma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡赞泰KazantayBarony struct {
	feud.BaseBarony
}

var BKazantay卡赞泰 feud.Barony = &卡赞泰KazantayBarony{}

func init() {
    f := BKazantay卡赞泰.(*卡赞泰KazantayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kazantay",
		TitleName: "卡赞泰",
		TitleCode: "b_kazantay",
	}
}
