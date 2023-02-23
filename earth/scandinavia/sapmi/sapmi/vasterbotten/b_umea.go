package vasterbotten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 于默奥UmeaBarony struct {
	feud.BaseBarony
}

var BUmea于默奥 feud.Barony = &于默奥UmeaBarony{}

func init() {
	f := BUmea于默奥.(*于默奥UmeaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "umea",
		TitleName: "于默奥",
		TitleCode: "b_umea",
	}
}
