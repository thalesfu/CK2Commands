package nice

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坎波格罗索CampogrossoBarony struct {
	feud.BaseBarony
}

var BCampogrosso坎波格罗索 feud.Barony = &坎波格罗索CampogrossoBarony{}

func init() {
    f := BCampogrosso坎波格罗索.(*坎波格罗索CampogrossoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "campogrosso",
		TitleName: "坎波格罗索",
		TitleCode: "b_campogrosso",
	}
}
