package rugen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特里布钦TribuzinBarony struct {
	feud.BaseBarony
}

var BTribuzin特里布钦 feud.Barony = &特里布钦TribuzinBarony{}

func init() {
    f := BTribuzin特里布钦.(*特里布钦TribuzinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tribuzin",
		TitleName: "特里布钦",
		TitleCode: "b_tribuzin",
	}
}
