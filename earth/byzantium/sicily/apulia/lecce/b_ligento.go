package lecce

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利真托LigentoBarony struct {
	feud.BaseBarony
}

var BLigento利真托 feud.Barony = &利真托LigentoBarony{}

func init() {
	f := BLigento利真托.(*利真托LigentoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ligento",
		TitleName: "利真托",
		TitleCode: "b_ligento",
	}
}
