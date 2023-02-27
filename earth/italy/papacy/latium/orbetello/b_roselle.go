package orbetello

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗塞莱RoselleBarony struct {
	feud.BaseBarony
}

var BRoselle罗塞莱 feud.Barony = &罗塞莱RoselleBarony{}

func init() {
    f := BRoselle罗塞莱.(*罗塞莱RoselleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "roselle",
		TitleName: "罗塞莱",
		TitleCode: "b_roselle",
	}
}
