package abosyssel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利斯比约LisbjergBarony struct {
	feud.BaseBarony
}

var BLisbjerg利斯比约 feud.Barony = &利斯比约LisbjergBarony{}

func init() {
    f := BLisbjerg利斯比约.(*利斯比约LisbjergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lisbjerg",
		TitleName: "利斯比约",
		TitleCode: "b_lisbjerg",
	}
}
