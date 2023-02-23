package lisboa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣塔伦SantaremBarony struct {
	feud.BaseBarony
}

var BSantarem圣塔伦 feud.Barony = &圣塔伦SantaremBarony{}

func init() {
	f := BSantarem圣塔伦.(*圣塔伦SantaremBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "santarem",
		TitleName: "圣塔伦",
		TitleCode: "b_santarem",
	}
}
