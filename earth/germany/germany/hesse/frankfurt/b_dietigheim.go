package frankfurt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪蒂希海姆DietigheimBarony struct {
	feud.BaseBarony
}

var BDietigheim迪蒂希海姆 feud.Barony = &迪蒂希海姆DietigheimBarony{}

func init() {
	f := BDietigheim迪蒂希海姆.(*迪蒂希海姆DietigheimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dietigheim",
		TitleName: "迪蒂希海姆",
		TitleCode: "b_dietigheim",
	}
}
