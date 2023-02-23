package tourraine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希农ChinonBarony struct {
	feud.BaseBarony
}

var BChinon希农 feud.Barony = &希农ChinonBarony{}

func init() {
	f := BChinon希农.(*希农ChinonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chinon",
		TitleName: "希农",
		TitleCode: "b_chinon",
	}
}
