package gurjaratra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加利安瓦拉KarianwalaBarony struct {
	feud.BaseBarony
}

var BKarianwala加利安瓦拉 feud.Barony = &加利安瓦拉KarianwalaBarony{}

func init() {
	f := BKarianwala加利安瓦拉.(*加利安瓦拉KarianwalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karianwala",
		TitleName: "加利安瓦拉",
		TitleCode: "b_karianwala",
	}
}
