package piombino

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波普洛尼亚PopuloniaBarony struct {
	feud.BaseBarony
}

var BPopulonia波普洛尼亚 feud.Barony = &波普洛尼亚PopuloniaBarony{}

func init() {
	f := BPopulonia波普洛尼亚.(*波普洛尼亚PopuloniaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "populonia",
		TitleName: "波普洛尼亚",
		TitleCode: "b_populonia",
	}
}
