package melitene

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨莫萨塔SamosataBarony struct {
	feud.BaseBarony
}

var BSamosata萨莫萨塔 feud.Barony = &萨莫萨塔SamosataBarony{}

func init() {
	f := BSamosata萨莫萨塔.(*萨莫萨塔SamosataBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "samosata",
		TitleName: "萨莫萨塔",
		TitleCode: "b_samosata",
	}
}
