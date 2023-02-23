package kolzum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索赫纳SokhnaBarony struct {
	feud.BaseBarony
}

var BSokhna索赫纳 feud.Barony = &索赫纳SokhnaBarony{}

func init() {
	f := BSokhna索赫纳.(*索赫纳SokhnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sokhna",
		TitleName: "索赫纳",
		TitleCode: "b_sokhna",
	}
}
