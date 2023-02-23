package lemdiyya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅德阿MedeaBarony struct {
	feud.BaseBarony
}

var BMedea梅德阿 feud.Barony = &梅德阿MedeaBarony{}

func init() {
	f := BMedea梅德阿.(*梅德阿MedeaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "medea",
		TitleName: "梅德阿",
		TitleCode: "b_medea",
	}
}
