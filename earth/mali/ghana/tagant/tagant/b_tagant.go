package tagant

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔甘特TagantBarony struct {
	feud.BaseBarony
}

var BTagant塔甘特 feud.Barony = &塔甘特TagantBarony{}

func init() {
    f := BTagant塔甘特.(*塔甘特TagantBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tagant",
		TitleName: "塔甘特",
		TitleCode: "b_tagant",
	}
}
