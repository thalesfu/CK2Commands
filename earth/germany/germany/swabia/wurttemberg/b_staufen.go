package wurttemberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 施陶芬StaufenBarony struct {
	feud.BaseBarony
}

var BStaufen施陶芬 feud.Barony = &施陶芬StaufenBarony{}

func init() {
	f := BStaufen施陶芬.(*施陶芬StaufenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "staufen",
		TitleName: "施陶芬",
		TitleCode: "b_staufen",
	}
}
