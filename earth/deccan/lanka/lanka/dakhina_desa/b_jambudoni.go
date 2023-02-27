package dakhina_desa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 旬波多尼JambudoniBarony struct {
	feud.BaseBarony
}

var BJambudoni旬波多尼 feud.Barony = &旬波多尼JambudoniBarony{}

func init() {
    f := BJambudoni旬波多尼.(*旬波多尼JambudoniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jambudoni",
		TitleName: "旬波多尼",
		TitleCode: "b_jambudoni",
	}
}
