package kataka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拘拿罗迦KonarakBarony struct {
	feud.BaseBarony
}

var BKonarak拘拿罗迦 feud.Barony = &拘拿罗迦KonarakBarony{}

func init() {
	f := BKonarak拘拿罗迦.(*拘拿罗迦KonarakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "konarak",
		TitleName: "拘拿罗迦",
		TitleCode: "b_konarak",
	}
}
