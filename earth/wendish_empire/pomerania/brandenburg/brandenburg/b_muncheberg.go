package brandenburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 明谢贝格MunchebergBarony struct {
	feud.BaseBarony
}

var BMuncheberg明谢贝格 feud.Barony = &明谢贝格MunchebergBarony{}

func init() {
    f := BMuncheberg明谢贝格.(*明谢贝格MunchebergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "muncheberg",
		TitleName: "明谢贝格",
		TitleCode: "b_muncheberg",
	}
}
