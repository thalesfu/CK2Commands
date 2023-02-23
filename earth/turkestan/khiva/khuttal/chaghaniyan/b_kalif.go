package chaghaniyan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡利甫KalifBarony struct {
	feud.BaseBarony
}

var BKalif卡利甫 feud.Barony = &卡利甫KalifBarony{}

func init() {
	f := BKalif卡利甫.(*卡利甫KalifBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalif",
		TitleName: "卡利甫",
		TitleCode: "b_kalif",
	}
}
