package hisar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦罗那郁罗KalanaurBarony struct {
	feud.BaseBarony
}

var BKalanaur迦罗那郁罗 feud.Barony = &迦罗那郁罗KalanaurBarony{}

func init() {
	f := BKalanaur迦罗那郁罗.(*迦罗那郁罗KalanaurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalanaur",
		TitleName: "迦罗那郁罗",
		TitleCode: "b_kalanaur",
	}
}
