package nikomedeia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦克墩ChalkedonBarony struct {
	feud.BaseBarony
}

var BChalkedon迦克墩 feud.Barony = &迦克墩ChalkedonBarony{}

func init() {
	f := BChalkedon迦克墩.(*迦克墩ChalkedonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chalkedon",
		TitleName: "迦克墩",
		TitleCode: "b_chalkedon",
	}
}
