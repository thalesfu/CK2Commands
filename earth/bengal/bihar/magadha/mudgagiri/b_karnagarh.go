package mudgagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 羯罗拏姞利呬KarnagarhBarony struct {
	feud.BaseBarony
}

var BKarnagarh羯罗拏姞利呬 feud.Barony = &羯罗拏姞利呬KarnagarhBarony{}

func init() {
	f := BKarnagarh羯罗拏姞利呬.(*羯罗拏姞利呬KarnagarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karnagarh",
		TitleName: "羯罗拏姞利呬",
		TitleCode: "b_karnagarh",
	}
}
