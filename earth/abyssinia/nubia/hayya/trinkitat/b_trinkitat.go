package trinkitat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特林基塔特TrinkitatBarony struct {
	feud.BaseBarony
}

var BTrinkitat特林基塔特 feud.Barony = &特林基塔特TrinkitatBarony{}

func init() {
    f := BTrinkitat特林基塔特.(*特林基塔特TrinkitatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "trinkitat",
		TitleName: "特林基塔特",
		TitleCode: "b_trinkitat",
	}
}
