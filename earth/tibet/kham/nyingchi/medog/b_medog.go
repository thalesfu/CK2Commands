package medog

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 墨脱MedogBarony struct {
	feud.BaseBarony
}

var BMedog墨脱 feud.Barony = &墨脱MedogBarony{}

func init() {
	f := BMedog墨脱.(*墨脱MedogBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "medog",
		TitleName: "墨脱",
		TitleCode: "b_medog",
	}
}
