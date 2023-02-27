package tripuri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗悉补罗RasidpurBarony struct {
	feud.BaseBarony
}

var BRasidpur罗悉补罗 feud.Barony = &罗悉补罗RasidpurBarony{}

func init() {
    f := BRasidpur罗悉补罗.(*罗悉补罗RasidpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rasidpur",
		TitleName: "罗悉补罗",
		TitleCode: "b_rasidpur",
	}
}
