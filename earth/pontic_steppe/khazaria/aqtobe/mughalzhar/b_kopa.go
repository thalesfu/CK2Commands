package mughalzhar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科帕KopaBarony struct {
	feud.BaseBarony
}

var BKopa科帕 feud.Barony = &科帕KopaBarony{}

func init() {
    f := BKopa科帕.(*科帕KopaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kopa",
		TitleName: "科帕",
		TitleCode: "b_kopa",
	}
}
