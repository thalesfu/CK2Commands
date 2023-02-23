package tyrnovo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡赞勒克KazanlakBarony struct {
	feud.BaseBarony
}

var BKazanlak卡赞勒克 feud.Barony = &卡赞勒克KazanlakBarony{}

func init() {
	f := BKazanlak卡赞勒克.(*卡赞勒克KazanlakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kazanlak",
		TitleName: "卡赞勒克",
		TitleCode: "b_kazanlak",
	}
}
