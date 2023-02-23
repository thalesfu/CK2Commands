package kathiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 劳代尔LodarBarony struct {
	feud.BaseBarony
}

var BLodar劳代尔 feud.Barony = &劳代尔LodarBarony{}

func init() {
	f := BLodar劳代尔.(*劳代尔LodarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lodar",
		TitleName: "劳代尔",
		TitleCode: "b_lodar",
	}
}
