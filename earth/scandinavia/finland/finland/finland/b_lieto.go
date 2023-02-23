package finland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 列托LietoBarony struct {
	feud.BaseBarony
}

var BLieto列托 feud.Barony = &列托LietoBarony{}

func init() {
	f := BLieto列托.(*列托LietoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lieto",
		TitleName: "列托",
		TitleCode: "b_lieto",
	}
}
