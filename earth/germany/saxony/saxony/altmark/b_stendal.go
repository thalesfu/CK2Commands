package altmark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 施滕达尔StendalBarony struct {
	feud.BaseBarony
}

var BStendal施滕达尔 feud.Barony = &施滕达尔StendalBarony{}

func init() {
	f := BStendal施滕达尔.(*施滕达尔StendalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stendal",
		TitleName: "施滕达尔",
		TitleCode: "b_stendal",
	}
}
