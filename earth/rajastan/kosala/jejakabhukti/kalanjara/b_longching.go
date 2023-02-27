package kalanjara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 长青LongchingBarony struct {
	feud.BaseBarony
}

var BLongching长青 feud.Barony = &长青LongchingBarony{}

func init() {
    f := BLongching长青.(*长青LongchingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "longching",
		TitleName: "长青",
		TitleCode: "b_longching",
	}
}
