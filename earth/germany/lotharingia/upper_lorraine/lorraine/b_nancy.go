package lorraine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 南锡NancyBarony struct {
	feud.BaseBarony
}

var BNancy南锡 feud.Barony = &南锡NancyBarony{}

func init() {
    f := BNancy南锡.(*南锡NancyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nancy",
		TitleName: "南锡",
		TitleCode: "b_nancy",
	}
}
