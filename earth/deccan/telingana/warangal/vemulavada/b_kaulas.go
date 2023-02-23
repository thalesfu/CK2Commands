package vemulavada

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伽郁罗湿KaulasBarony struct {
	feud.BaseBarony
}

var BKaulas伽郁罗湿 feud.Barony = &伽郁罗湿KaulasBarony{}

func init() {
	f := BKaulas伽郁罗湿.(*伽郁罗湿KaulasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kaulas",
		TitleName: "伽郁罗湿",
		TitleCode: "b_kaulas",
	}
}
