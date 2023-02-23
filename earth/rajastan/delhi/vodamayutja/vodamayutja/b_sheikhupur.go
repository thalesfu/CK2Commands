package vodamayutja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舍侯补罗SheikhupurBarony struct {
	feud.BaseBarony
}

var BSheikhupur舍侯补罗 feud.Barony = &舍侯补罗SheikhupurBarony{}

func init() {
	f := BSheikhupur舍侯补罗.(*舍侯补罗SheikhupurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sheikhupur",
		TitleName: "舍侯补罗",
		TitleCode: "b_sheikhupur",
	}
}
