package khojand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索赫SokhBarony struct {
	feud.BaseBarony
}

var BSokh索赫 feud.Barony = &索赫SokhBarony{}

func init() {
    f := BSokh索赫.(*索赫SokhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sokh",
		TitleName: "索赫",
		TitleCode: "b_sokh",
	}
}
