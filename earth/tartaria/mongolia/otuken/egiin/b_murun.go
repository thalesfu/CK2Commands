package egiin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 木伦MurunBarony struct {
	feud.BaseBarony
}

var BMurun木伦 feud.Barony = &木伦MurunBarony{}

func init() {
    f := BMurun木伦.(*木伦MurunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "murun",
		TitleName: "木伦",
		TitleCode: "b_murun",
	}
}
