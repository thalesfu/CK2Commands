package alcala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡古恩萨SiguenzaBarony struct {
	feud.BaseBarony
}

var BSiguenza锡古恩萨 feud.Barony = &锡古恩萨SiguenzaBarony{}

func init() {
	f := BSiguenza锡古恩萨.(*锡古恩萨SiguenzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "siguenza",
		TitleName: "锡古恩萨",
		TitleCode: "b_siguenza",
	}
}
