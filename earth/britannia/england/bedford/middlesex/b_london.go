package middlesex

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伦敦LondonBarony struct {
	feud.BaseBarony
}

var BLondon伦敦 feud.Barony = &伦敦LondonBarony{}

func init() {
	f := BLondon伦敦.(*伦敦LondonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "london",
		TitleName: "伦敦",
		TitleCode: "b_london",
	}
}
