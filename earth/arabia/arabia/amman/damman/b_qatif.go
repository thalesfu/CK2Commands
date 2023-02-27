package damman

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 盖提夫QatifBarony struct {
	feud.BaseBarony
}

var BQatif盖提夫 feud.Barony = &盖提夫QatifBarony{}

func init() {
    f := BQatif盖提夫.(*盖提夫QatifBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qatif",
		TitleName: "盖提夫",
		TitleCode: "b_qatif",
	}
}
