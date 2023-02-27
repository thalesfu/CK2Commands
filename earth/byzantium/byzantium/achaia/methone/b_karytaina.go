package methone

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡里泰纳KarytainaBarony struct {
	feud.BaseBarony
}

var BKarytaina卡里泰纳 feud.Barony = &卡里泰纳KarytainaBarony{}

func init() {
    f := BKarytaina卡里泰纳.(*卡里泰纳KarytainaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karytaina",
		TitleName: "卡里泰纳",
		TitleCode: "b_karytaina",
	}
}
