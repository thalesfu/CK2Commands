package koloneia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃金AknBarony struct {
	feud.BaseBarony
}

var BAkn埃金 feud.Barony = &埃金AknBarony{}

func init() {
	f := BAkn埃金.(*埃金AknBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akn",
		TitleName: "埃金",
		TitleCode: "b_akn",
	}
}
